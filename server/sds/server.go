package sds

import (
	"context"
	"errors"

	v3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	secretservice "github.com/envoyproxy/go-control-plane/envoy/service/secret/v3"
)

type Server struct {
	secretservice.SecretDiscoveryServiceServer
}

func (*Server) DeltaSecrets(_ secretservice.SecretDiscoveryService_DeltaSecretsServer) error {
	return errors.New("not implemented")
}

func (*Server) StreamSecrets(_ secretservice.SecretDiscoveryService_StreamSecretsServer) error {
	return errors.New("not implemented")
}

func (*Server) FetchSecrets(_ context.Context, _ *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error) {
	return nil, errors.New("not implemented")
}
