//go:generate protoc -I ../proto --go_out=plugins=grpc:../proto ../proto/esm.proto

package main

import (
	"log"
	"net"

	"github.com/lisuiheng/esm/ipmi"
	pb "github.com/lisuiheng/esm/rpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	client ipmi.Client
}

// SayHello implements helloworld.GreeterServer
func (s *server) GetGuid(ctx context.Context, in *pb.GetGuidRequest) (*pb.GetGuidReply, error) {
	c, err := ipmi.NewClient(&ipmi.Connection{
		Hostname:  "172.16.101.10",
		Port:      623,
		Username:  "Administrator",
		Password:  "Administrator",
		Interface: "lanplus",
	})
	if err != nil {
		return nil, err
	}
	s.client = c
	guid, err := s.client.GetGuid()
	if err != nil {
		return nil, err
	}
	return &pb.GetGuidReply{Guid: guid}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEsmServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
