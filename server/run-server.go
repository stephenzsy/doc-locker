package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	secretservice "github.com/envoyproxy/go-control-plane/envoy/service/secret/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	cachev3 "github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/stephenzsy/doc-locker/server/common/configurations"
	hostService "github.com/stephenzsy/doc-locker/server/gen/host"
	"github.com/stephenzsy/doc-locker/server/host"
	"github.com/stephenzsy/doc-locker/server/sds"
)

var (
	sdsFlag = flag.Bool("sds", false, "Run SDS server")
)

func serveSds() {
	flag.Parse()

	l := sds.Logger{
		Debug: true,
	}
	// Create a cache
	cache := cachev3.NewSnapshotCache(false, cachev3.IDHash{}, l)

	// Create the snapshot that we'll serve to Envoy
	snapshot := cachev3.NewSnapshot(
		"1",
		[]types.Resource{}, // endpoints
		[]types.Resource{},
		[]types.Resource{},
		[]types.Resource{},
		[]types.Resource{}, // runtimes
		[]types.Resource{}, // secrets
	)
	if err := snapshot.Consistent(); err != nil {
		l.Errorf("snapshot inconsistency: %+v\n%+v", snapshot, err)
		os.Exit(1)
	}
	l.Debugf("will serve snapshot %+v", snapshot)

	// Add the snapshot to the cache
	if err := cache.SetSnapshot("test-id", snapshot); err != nil {
		l.Errorf("snapshot error %q for %+v", err, snapshot)
		os.Exit(1)
	}

	// Run the xDS server
	ctx := context.Background()
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
	server := sds.NewServer(ctx, cache)
	secretservice.RegisterSecretDiscoveryServiceServer(sdsGrpcServer, &server)
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
