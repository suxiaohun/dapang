package face_worker

import (
	"context"
	"dapang/pb"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

var (
	port = flag.Int("port", 50050, "The server port")
)

type server struct {
	pb.UnimplementedFaceApiServer
}

func (s *server) FaceVerify(ctx context.Context, in *pb.FaceVerifyRequest) (*pb.FaceVerifyResponse, error) {
	return &pb.FaceVerifyResponse{
		Code:    1000,
		Score:   0.6,
		Message: "ok",
	}, nil
}

func (s *server) FaceExtract(ctx context.Context, in *pb.FaceExtractRequest) (*pb.FaceExtractResponse, error) {
	return &pb.FaceExtractResponse{}, nil
}

func (s *server) FeatVerify(ctx context.Context, in *pb.FeatVerifyRequest) (*pb.FeatVerifyResponse, error) {
	return &pb.FeatVerifyResponse{}, nil
}

func Run() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFaceApiServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
