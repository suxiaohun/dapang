package main

import (
	"context"
	pb_searcher "dapang/grpc/searcher"
	pb_worker "dapang/grpc/worker"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"time"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type clusterAccumulator func([]*pb_worker.ClusterShardInfo) (uint64, uint64)

type Reducer struct {
	pb_searcher.UnsafeSearchServiceServer
	workers map[string]*grpc.ClientConn
	config  *Config
}

type workerSystemInfo struct {
	workerID string
	response *pb_worker.ListShardsResponse
}

func (r *Reducer) GetSystemInfo(ctx context.Context, in *pb_searcher.GetSystemInfoRequest) (*pb_searcher.GetSystemInfoResponse, error) {
	resp := &pb_searcher.GetSystemInfoResponse{}
	ch := make(chan *workerSystemInfo, len(r.workers))
	for worker, conn := range r.workers {
		go func(worker string, conn *grpc.ClientConn) {
			req := &pb_worker.ListShardsRequest{}
			info := &workerSystemInfo{workerID: worker}
			c := pb_worker.NewWorkerServiceClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.config.RPCTimeoutSeconds)*time.Second)
			defer cancel()

			res, err := c.ListShards(ctx, req)
			if err != nil || res == nil {
				log.Warnf("failed to list shard of %s, %v", worker, err)
			} else {
				info.response = res
			}
			ch <- info
		}(worker, conn)
	}

	regions := make(map[int32]*pb_searcher.RegionInfo)
	counter := newClusterAccumulator()
	for i := 0; i < len(r.workers); i++ {
		mergeSystemInfo(resp, regions, <-ch, counter)
	}

	resp.Shards = 3
	return resp, nil
}

func NewReducer() *Reducer {
	return &Reducer{
		config: &Config{RPCTimeoutSeconds: 2},
	}
}

func newClusterAccumulator() clusterAccumulator {
	var lastVersion int64
	var totalClusterShards uint64
	var totalClusters uint64
	shardSet := make(map[string]bool)
	return func(shards []*pb_worker.ClusterShardInfo) (uint64, uint64) {
		for _, shard := range shards {
			// if newer version got
			if shard.GetVersion() > lastVersion {
				lastVersion = shard.GetVersion()
				totalClusters = shard.GetShardSize()
				totalClusterShards = 1
			} else if shard.GetVersion() == lastVersion { // or already the accumulated version
				if shardSet[shard.GetIndexId()] {
					continue
				}
				totalClusters += shard.GetShardSize()
				totalClusterShards++
			}
			shardSet[shard.GetIndexId()] = true
		}

		return totalClusterShards, totalClusters
	}
}

func mergeSystemInfo(resp *pb_searcher.GetSystemInfoResponse, regions map[int32]*pb_searcher.RegionInfo,
	workerInfo *workerSystemInfo, accumulateCluster clusterAccumulator) {
	if workerInfo.response == nil || workerInfo.response.GetNodeInfo() == nil {
		resp.OfflineWorkers = append(resp.OfflineWorkers, workerInfo.workerID)
		return
	}
	nodeInfo := workerInfo.response.GetNodeInfo()
	if nodeInfo.GetStatus() == pb_worker.NodeStatus_Dead {
		resp.OfflineWorkers = append(resp.OfflineWorkers, workerInfo.workerID)
	} else {
		resp.OnlineWorkers = append(resp.OnlineWorkers, workerInfo.workerID)
		resp.LicenseCapacity = resp.LicenseCapacity + nodeInfo.GetLicenseCapacity()
		resp.UsedPercentage += nodeInfo.GetLoadFactor()
		resp.TotalGpuMemorySize += nodeInfo.GetTotalGpuMemorySize()
		resp.AvailableGpuMemorySize += nodeInfo.GetAvailableGpuMemorySize()
		resp.TotalMemorySize += nodeInfo.GetTotalMemorySize()
		resp.AvailableMemorySize += nodeInfo.GetAvailableMemorySize()
	}

	wInfo := &pb_searcher.WorkerInfo{
		WorkerId:      workerInfo.workerID,
		Shards:        uint64(len(workerInfo.response.GetShardInfos())),
		ClusterShards: uint64(len(workerInfo.response.GetClusterShardInfos())),
	}

	for _, clusterShard := range workerInfo.response.GetClusterShardInfos() {
		wInfo.Clusters += clusterShard.GetShardSize()
	}

	resp.Shards += uint64(len(workerInfo.response.GetShardInfos()))

	// count cluster shards and clusters of the latest loaded version
	resp.ClusterShards, resp.Clusters = accumulateCluster(workerInfo.response.GetClusterShardInfos())

	//count up shard info
	regionIDSet := make(map[int32]bool)
	for _, shard := range workerInfo.response.GetShardInfos() {
		resp.Features += shard.GetShardSize()
		wInfo.Features += shard.GetShardSize()
		if utils.TimestampLess(shard.GetLastTime(), resp.GetFirstTimeToDelete()) {
			resp.FirstTimeToDelete = shard.GetLastTime()
		}
		regionID := shard.GetRegionId()
		if _, ok := regionIDSet[regionID]; !ok {
			regionIDSet[regionID] = true
		}
		if val, ok := regions[regionID]; ok {
			updateRegionInfo(shard, val)
			continue
		}

		//a new region found
		regionInfo := &pb_searcher.RegionInfo{
			RegionId: regionID,
			Features: shard.GetShardSize(),
			Period: &pb_common.TimeRange{
				Start: shard.GetFirstTime(),
				End:   shard.GetLastTime(),
			},
		}
		regions[regionID] = regionInfo
		if resp.GetFirstTimeToDelete() == nil {
			resp.FirstTimeToDelete = shard.GetLastTime()
		}
	}

	wInfo.RegionIds = int32(len(regionIDSet))
	resp.WorkerInfos = append(resp.WorkerInfos, wInfo)
}

func main() {
	flag.Parse()
	r := NewReducer()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb_searcher.RegisterSearchServiceServer(s, r)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
