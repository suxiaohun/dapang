package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"time"

	pb "dapang/grpc/searcher"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSearchServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetSystemInfo(ctx, &pb.GetSystemInfoRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	b, err := json.Marshal(r)
	if err != nil {
		log.Fatalf("could not parse res to json")
	}

	log.Printf("GetSystemInfo: %v", string(b))

}
