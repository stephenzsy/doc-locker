package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"path"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	secretservice "github.com/envoyproxy/go-control-plane/envoy/service/secret/v3"
	"github.com/stephenzsy/doc-locker/server/common/configurations"
	hostService "github.com/stephenzsy/doc-locker/server/gen/host"
	"github.com/stephenzsy/doc-locker/server/host"
	"github.com/stephenzsy/doc-locker/server/sds"
)

var (
	sdsFlag = flag.Bool("sds", false, "Run SDS server")
)

func serveSds() {
	sdsLis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 21000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	config := configurations.Configurations()
	configRootDir := config.ConfigRootDir()
	certPath := path.Join(configRootDir, "certs", "server-cert-deploy-sds.pem")
	keyPath := path.Join(configRootDir, "certsk", "server-key-deploy-sds.pem")
	creds, err := credentials.NewServerTLSFromFile(certPath, keyPath)
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}
	sdsOpts := []grpc.ServerOption{grpc.Creds(creds)}

	sdsGrpcServer := grpc.NewServer(sdsOpts...)
	secretservice.RegisterSecretDiscoveryServiceServer(sdsGrpcServer, &sds.Server{})
	sdsGrpcServer.Serve(sdsLis)
}

func main() {
	flag.Parse()
	if *sdsFlag {
		serveSds()
		return
	}
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 11000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	hostService.RegisterHostServiceServer(grpcServer, &host.HostServiceServer{})

	grpcServer.Serve(lis)
}
