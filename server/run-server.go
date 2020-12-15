package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	secretservice "github.com/envoyproxy/go-control-plane/envoy/service/secret/v3"
	"github.com/stephenzsy/doc-locker/server/common/app_context"
	"github.com/stephenzsy/doc-locker/server/common/auth"
	"github.com/stephenzsy/doc-locker/server/common/configurations"
	hostService "github.com/stephenzsy/doc-locker/server/gen/host"
	"github.com/stephenzsy/doc-locker/server/host"
	"github.com/stephenzsy/doc-locker/server/sds"
)

var (
	sdsFlag = flag.Bool("sds", false, "Run SDS server")
)

func serveSds() (err error) {
	flag.Parse()

	// Create a cache
	// Run the xDS server
	ctx, err := app_context.NewAppServiceContext(context.Background(), auth.ServiceCallerIdSds)
	if err != nil {
		return
	}
	sdsLis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 21000))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	secretsConfig, err := configurations.GetSecretsConfiguration(ctx.Elevate())
	if err != nil {
		return
	}
	creds, err := credentials.NewServerTLSFromFile(
		secretsConfig.GetCertPath(configurations.SecretTypeServer, configurations.SecretNameDeploySds),
		secretsConfig.GetPrivateKeyPath(configurations.SecretTypeServer, configurations.SecretNameDeploySds))
	if err != nil {
		return fmt.Errorf("Failed to generate credentials %v", err)
	}
	sdsOpts := []grpc.ServerOption{grpc.Creds(creds)}

	sdsGrpcServer := grpc.NewServer(sdsOpts...)

	server, err := sds.NewServer(ctx)
	if err != nil {
		return
	}
	secretservice.RegisterSecretDiscoveryServiceServer(sdsGrpcServer, &server)
	sdsGrpcServer.Serve(sdsLis)
	return
}

func main() {
	flag.Parse()
	if *sdsFlag {
		if err := serveSds(); err != nil {
			log.Panic(err)
		}
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
