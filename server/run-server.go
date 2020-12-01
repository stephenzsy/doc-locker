package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	secretservice "github.com/envoyproxy/go-control-plane/envoy/service/secret/v3"
	hostService "github.com/stephenzsy/doc-locker/server/gen/host"
	"github.com/stephenzsy/doc-locker/server/host"
	"github.com/stephenzsy/doc-locker/server/sds"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 11000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	sdsLis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 12000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	hostService.RegisterHostServiceServer(grpcServer, &host.HostServiceServer{})
	sdsGrpcServer := grpc.NewServer(opts...)
	secretservice.RegisterSecretDiscoveryServiceServer(sdsGrpcServer, &sds.Server{})
	grpcServer.Serve(lis)
	sdsGrpcServer.Serve(sdsLis)
}
