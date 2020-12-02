package sds

import (
	"context"
	"errors"
	"log"

	v3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	secretservice "github.com/envoyproxy/go-control-plane/envoy/service/secret/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/sotw/v3"
	testv3 "github.com/envoyproxy/go-control-plane/pkg/test/v3"
)

type server struct {
	sotw sotw.Server
}

func NewServer(ctx context.Context, config cache.ConfigWatcher) server {
	s := server{
		sotw: sotw.NewServer(ctx, config, &testv3.Callbacks{}),
	}
	return s
}

func (*server) DeltaSecrets(_ secretservice.SecretDiscoveryService_DeltaSecretsServer) error {
	log.Fatal("Delta Secrets")
	return errors.New("not implemented")
}

func (s *server) StreamSecrets(stream secretservice.SecretDiscoveryService_StreamSecretsServer) error {
	return s.sotw.StreamHandler(stream, resource.SecretType)
}

func (*server) FetchSecrets(_ context.Context, _ *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error) {
	log.Fatal("Fetch Secrets")
	return nil, errors.New("not implemented")
}
