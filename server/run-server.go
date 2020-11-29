package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/stephenzsy/doc-locker/server/gen/host"
	"github.com/stephenzsy/doc-locker/server/host"
)

func newServer() *host.HostServiceServer {
	s := &host.HostServiceServer{}
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 11000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterHostServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
