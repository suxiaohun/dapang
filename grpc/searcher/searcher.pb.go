// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: searcher/searcher.proto

package searcher

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	api "gitlab.sz.sensetime.com/viper/commonapis/api"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 区域ID.
	// [EN] Region ID.
	RegionId int32 `protobuf:"varint,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	// 本区域特征数.
	// [EN] Number of features in this region.
	Features uint64 `protobuf:"varint,2,opt,name=features,proto3" json:"features,omitempty"`
	// 本区域数据时间范围.
	// [EN] Time range of data in this region.
	Period *api.TimeRange `protobuf:"bytes,3,opt,name=period,proto3" json:"period,omitempty"`
}

func (x *RegionInfo) Reset() {
	*x = RegionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_searcher_searcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionInfo) ProtoMessage() {}

func (x *RegionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_searcher_searcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionInfo.ProtoReflect.Descriptor instead.
func (*RegionInfo) Descriptor() ([]byte, []int) {
	return file_searcher_searcher_proto_rawDescGZIP(), []int{0}
}

func (x *RegionInfo) GetRegionId() int32 {
	if x != nil {
		return x.RegionId
	}
	return 0
}

func (x *RegionInfo) GetFeatures() uint64 {
	if x != nil {
		return x.Features
	}
	return 0
}

func (x *RegionInfo) GetPeriod() *api.TimeRange {
	if x != nil {
		return x.Period
	}
	return nil
}

type WorkerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// worker名.
	// [EN] Worker ID.
	WorkerId string `protobuf:"bytes,1,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
	// 区域ID数.
	// [EN] Number of regions in this worker.
	RegionIds int32 `protobuf:"varint,2,opt,name=region_ids,json=regionIds,proto3" json:"region_ids,omitempty"`
	// 分片数.
	// [EN] Total number of worker shard.
	Shards uint64 `protobuf:"varint,3,opt,name=shards,proto3" json:"shards,omitempty"`
	// 本区域特征数.
	// [EN] Number of features in this worker.
	Features uint64 `protobuf:"varint,4,opt,name=features,proto3" json:"features,omitempty"`
	// 聚类中心分片数.
	// [EN] The number of shards of cluster centroid.
	ClusterShards uint64 `protobuf:"varint,5,opt,name=cluster_shards,json=clusterShards,proto3" json:"cluster_shards,omitempty"`
	// 聚类中心数.
	// [EN] Number of cluster centroid.
	Clusters uint64 `protobuf:"varint,6,opt,name=clusters,proto3" json:"clusters,omitempty"`
}

func (x *WorkerInfo) Reset() {
	*x = WorkerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_searcher_searcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerInfo) ProtoMessage() {}

func (x *WorkerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_searcher_searcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerInfo.ProtoReflect.Descriptor instead.
func (*WorkerInfo) Descriptor() ([]byte, []int) {
	return file_searcher_searcher_proto_rawDescGZIP(), []int{1}
}

func (x *WorkerInfo) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

func (x *WorkerInfo) GetRegionIds() int32 {
	if x != nil {
		return x.RegionIds
	}
	return 0
}

func (x *WorkerInfo) GetShards() uint64 {
	if x != nil {
		return x.Shards
	}
	return 0
}

func (x *WorkerInfo) GetFeatures() uint64 {
	if x != nil {
		return x.Features
	}
	return 0
}

func (x *WorkerInfo) GetClusterShards() uint64 {
	if x != nil {
		return x.ClusterShards
	}
	return 0
}

func (x *WorkerInfo) GetClusters() uint64 {
	if x != nil {
		return x.Clusters
	}
	return 0
}

type GetSystemInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSystemInfoRequest) Reset() {
	*x = GetSystemInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_searcher_searcher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSystemInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSystemInfoRequest) ProtoMessage() {}

func (x *GetSystemInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_searcher_searcher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSystemInfoRequest.ProtoReflect.Descriptor instead.
func (*GetSystemInfoRequest) Descriptor() ([]byte, []int) {
	return file_searcher_searcher_proto_rawDescGZIP(), []int{2}
}

type GetSystemInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 系统总分片数.
	// [EN] Total number of system shard.
	Shards uint64 `protobuf:"varint,1,opt,name=shards,proto3" json:"shards,omitempty"`
	// 系统总特征数.
	// [EN] Total number of system feature.
	Features uint64 `protobuf:"varint,2,opt,name=features,proto3" json:"features,omitempty"`
	// 正常运行worker列表.
	// [EN] Normally running worker list.
	OnlineWorkers []string `protobuf:"bytes,3,rep,name=online_workers,json=onlineWorkers,proto3" json:"online_workers,omitempty"`
	// 下线worker列表.
	// [EN] Offline worker list.
	OfflineWorkers []string `protobuf:"bytes,4,rep,name=offline_workers,json=offlineWorkers,proto3" json:"offline_workers,omitempty"`
	// 区域统计信息.
	// [EN] Regional statistics.
	RegionInfos []*RegionInfo `protobuf:"bytes,5,rep,name=region_infos,json=regionInfos,proto3" json:"region_infos,omitempty"`
	// 系统总用量百分比.
	// [EN] Percentage of system total usage.
	UsedPercentage float32 `protobuf:"fixed32,7,opt,name=used_percentage,json=usedPercentage,proto3" json:"used_percentage,omitempty"`
	// 聚类中心分片数 [SINCE v2.0.0].
	// [EN] The number of shards of cluster centroid [SINCE v2.0.0].
	ClusterShards uint64 `protobuf:"varint,8,opt,name=cluster_shards,json=clusterShards,proto3" json:"cluster_shards,omitempty"`
	// 聚类中心数 [SINCE v2.0.0].
	// [EN] Number of cluster centroid [SINCE v2.0.0].
	Clusters uint64 `protobuf:"varint,9,opt,name=clusters,proto3" json:"clusters,omitempty"`
	// 系统内部状态记录, 仅供内部调试之用 [SINCE v2.0.0].
	// [EN] Internal status record of the system, for internal debug only [SINCE v2.0.0].
	SystemMetadata map[string]string `protobuf:"bytes,10,rep,name=system_metadata,json=systemMetadata,proto3" json:"system_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // TODO LicenseInfo
	// License限制下的总库容(可容纳特征数) [SINCE v2.0.0].
	// [EN] Total storage capacity limit by License (number of allowable features) [SINCE v2.0.0].
	LicenseCapacity uint64 `protobuf:"varint,11,opt,name=license_capacity,json=licenseCapacity,proto3" json:"license_capacity,omitempty"`
	// 系统可用总显存大小(不含预留显存), 单位Bytes [SINCE v2.0.0].
	// [EN] Total GPU memory size of system available (excluding reserved GPU memory), unit: Bytes [SINCE v2.0.0].
	AvailableGpuMemorySize uint64 `protobuf:"varint,12,opt,name=available_gpu_memory_size,json=availableGpuMemorySize,proto3" json:"available_gpu_memory_size,omitempty"`
	// 系统在线workers总显存, 单位Bytes [SINCE v2.0.0].
	// [EN] Total GPU memory of system online workers, unit: Bytes [SINCE v2.0.0].
	TotalGpuMemorySize uint64 `protobuf:"varint,13,opt,name=total_gpu_memory_size,json=totalGpuMemorySize,proto3" json:"total_gpu_memory_size,omitempty"`
	// 最早有效删除时间,早于此时间无法删除任何数据, 无法删除时为空 [SINCE v2.0.0].
	// [EN] The earliest valid delete time, before this time can not delete any data. It will be null when cannot be deleted [SINCE v2.0.0].
	FirstTimeToDelete *timestamp.Timestamp `protobuf:"bytes,14,opt,name=first_time_to_delete,json=firstTimeToDelete,proto3" json:"first_time_to_delete,omitempty"`
	// 系统可用总内存大小, 单位Bytes [SINCE v2.2.0].
	// [EN] Total memory size of system available, unit: Bytes [SINCE v2.2.0].
	AvailableMemorySize uint64 `protobuf:"varint,15,opt,name=available_memory_size,json=availableMemorySize,proto3" json:"available_memory_size,omitempty"`
	// 系统在线workers总内存, 单位Bytes [SINCE v2.2.0].
	// [EN] Total memory of system online workers, unit: Bytes [SINCE v2.2.0].
	TotalMemorySize uint64 `protobuf:"varint,16,opt,name=total_memory_size,json=totalMemorySize,proto3" json:"total_memory_size,omitempty"`
	// worker统计信息 [SINCE v4.2.0].
	// [EN] Worker statistics [SINCE v4.2.0].
	WorkerInfos []*WorkerInfo `protobuf:"bytes,17,rep,name=worker_infos,json=workerInfos,proto3" json:"worker_infos,omitempty"`
}

func (x *GetSystemInfoResponse) Reset() {
	*x = GetSystemInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_searcher_searcher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSystemInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSystemInfoResponse) ProtoMessage() {}

func (x *GetSystemInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_searcher_searcher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSystemInfoResponse.ProtoReflect.Descriptor instead.
func (*GetSystemInfoResponse) Descriptor() ([]byte, []int) {
	return file_searcher_searcher_proto_rawDescGZIP(), []int{3}
}

func (x *GetSystemInfoResponse) GetShards() uint64 {
	if x != nil {
		return x.Shards
	}
	return 0
}

func (x *GetSystemInfoResponse) GetFeatures() uint64 {
	if x != nil {
		return x.Features
	}
	return 0
}

func (x *GetSystemInfoResponse) GetOnlineWorkers() []string {
	if x != nil {
		return x.OnlineWorkers
	}
	return nil
}

func (x *GetSystemInfoResponse) GetOfflineWorkers() []string {
	if x != nil {
		return x.OfflineWorkers
	}
	return nil
}

func (x *GetSystemInfoResponse) GetRegionInfos() []*RegionInfo {
	if x != nil {
		return x.RegionInfos
	}
	return nil
}

func (x *GetSystemInfoResponse) GetUsedPercentage() float32 {
	if x != nil {
		return x.UsedPercentage
	}
	return 0
}

func (x *GetSystemInfoResponse) GetClusterShards() uint64 {
	if x != nil {
		return x.ClusterShards
	}
	return 0
}

func (x *GetSystemInfoResponse) GetClusters() uint64 {
	if x != nil {
		return x.Clusters
	}
	return 0
}

func (x *GetSystemInfoResponse) GetSystemMetadata() map[string]string {
	if x != nil {
		return x.SystemMetadata
	}
	return nil
}

func (x *GetSystemInfoResponse) GetLicenseCapacity() uint64 {
	if x != nil {
		return x.LicenseCapacity
	}
	return 0
}

func (x *GetSystemInfoResponse) GetAvailableGpuMemorySize() uint64 {
	if x != nil {
		return x.AvailableGpuMemorySize
	}
	return 0
}

func (x *GetSystemInfoResponse) GetTotalGpuMemorySize() uint64 {
	if x != nil {
		return x.TotalGpuMemorySize
	}
	return 0
}

func (x *GetSystemInfoResponse) GetFirstTimeToDelete() *timestamp.Timestamp {
	if x != nil {
		return x.FirstTimeToDelete
	}
	return nil
}

func (x *GetSystemInfoResponse) GetAvailableMemorySize() uint64 {
	if x != nil {
		return x.AvailableMemorySize
	}
	return 0
}

func (x *GetSystemInfoResponse) GetTotalMemorySize() uint64 {
	if x != nil {
		return x.TotalMemorySize
	}
	return 0
}

func (x *GetSystemInfoResponse) GetWorkerInfos() []*WorkerInfo {
	if x != nil {
		return x.WorkerInfos
	}
	return nil
}

var File_searcher_searcher_proto protoreflect.FileDescriptor

var file_searcher_searcher_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x65, 0x72, 0x1a, 0x3c, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x73, 0x7a, 0x2e, 0x73,
	0x65, 0x6e, 0x73, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x69, 0x70,
	0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x62,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x84, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x65,
	0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x69, 0x70, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22, 0xbf, 0x01, 0x0a, 0x0a, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x68, 0x61, 0x72, 0x64, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0xe0, 0x06, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x68, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x68, 0x61,
	0x72, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12,
	0x25, 0x0a, 0x0e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x57,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e,
	0x65, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0e, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x12,
	0x37, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x64,
	0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x64, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x61,
	0x72, 0x64, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x53, 0x68, 0x61, 0x72, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x5c, 0x0a, 0x0f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x29, 0x0a, 0x10, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x61,
	0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x6c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x39, 0x0a,
	0x19, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x67, 0x70, 0x75, 0x5f, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x16, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x47, 0x70, 0x75, 0x4d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x31, 0x0a, 0x15, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x67, 0x70, 0x75, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x47, 0x70,
	0x75, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x4b, 0x0a, 0x14, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x6f, 0x5f, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x11, 0x66, 0x69, 0x72, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x54, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2a, 0x0a, 0x11,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x73, 0x1a, 0x41, 0x0a, 0x13, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x32, 0x7e, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12,
	0x13, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x42, 0x16, 0x5a, 0x14, 0x64, 0x61, 0x70, 0x61, 0x6e, 0x67, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_searcher_searcher_proto_rawDescOnce sync.Once
	file_searcher_searcher_proto_rawDescData = file_searcher_searcher_proto_rawDesc
)

func file_searcher_searcher_proto_rawDescGZIP() []byte {
	file_searcher_searcher_proto_rawDescOnce.Do(func() {
		file_searcher_searcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_searcher_searcher_proto_rawDescData)
	})
	return file_searcher_searcher_proto_rawDescData
}

var file_searcher_searcher_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_searcher_searcher_proto_goTypes = []interface{}{
	(*RegionInfo)(nil),            // 0: searcher.RegionInfo
	(*WorkerInfo)(nil),            // 1: searcher.WorkerInfo
	(*GetSystemInfoRequest)(nil),  // 2: searcher.GetSystemInfoRequest
	(*GetSystemInfoResponse)(nil), // 3: searcher.GetSystemInfoResponse
	nil,                           // 4: searcher.GetSystemInfoResponse.SystemMetadataEntry
	(*api.TimeRange)(nil),         // 5: sensetime.viper.commonapis.TimeRange
	(*timestamp.Timestamp)(nil),   // 6: google.protobuf.Timestamp
}
var file_searcher_searcher_proto_depIdxs = []int32{
	5, // 0: searcher.RegionInfo.period:type_name -> sensetime.viper.commonapis.TimeRange
	0, // 1: searcher.GetSystemInfoResponse.region_infos:type_name -> searcher.RegionInfo
	4, // 2: searcher.GetSystemInfoResponse.system_metadata:type_name -> searcher.GetSystemInfoResponse.SystemMetadataEntry
	6, // 3: searcher.GetSystemInfoResponse.first_time_to_delete:type_name -> google.protobuf.Timestamp
	1, // 4: searcher.GetSystemInfoResponse.worker_infos:type_name -> searcher.WorkerInfo
	2, // 5: searcher.SearchService.GetSystemInfo:input_type -> searcher.GetSystemInfoRequest
	3, // 6: searcher.SearchService.GetSystemInfo:output_type -> searcher.GetSystemInfoResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_searcher_searcher_proto_init() }
func file_searcher_searcher_proto_init() {
	if File_searcher_searcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_searcher_searcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegionInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_searcher_searcher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_searcher_searcher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSystemInfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_searcher_searcher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSystemInfoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_searcher_searcher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_searcher_searcher_proto_goTypes,
		DependencyIndexes: file_searcher_searcher_proto_depIdxs,
		MessageInfos:      file_searcher_searcher_proto_msgTypes,
	}.Build()
	File_searcher_searcher_proto = out.File
	file_searcher_searcher_proto_rawDesc = nil
	file_searcher_searcher_proto_goTypes = nil
	file_searcher_searcher_proto_depIdxs = nil
}