package sds

import (
	"context"
	"errors"
	"log"

	v3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	secretservice "github.com/envoyproxy/go-control-plane/envoy/service/secret/v3"
)

type Server struct {
	secretservice.UnimplementedSecretDiscoveryServiceServer
}

func (*Server) DeltaSecrets(_ secretservice.SecretDiscoveryService_DeltaSecretsServer) error {
	log.Fatal("Delta Secrets")
	return errors.New("not implemented")
}

func (*Server) StreamSecrets(_ secretservice.SecretDiscoveryService_StreamSecretsServer) error {
	log.Fatal("Stream Secrets")
	return errors.New("not implemented")
}

func (*Server) FetchSecrets(_ context.Context, _ *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error) {
	log.Fatal("Fetch Secrets")
	return nil, errors.New("not implemented")
}
