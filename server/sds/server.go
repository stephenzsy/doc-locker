package sds

import (
	"bytes"
	"context"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"time"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	secrets "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	secretservice "github.com/envoyproxy/go-control-plane/envoy/service/secret/v3"
	resourcev3 "github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	"github.com/google/uuid"
	"github.com/stephenzsy/doc-locker/server/common/configurations"
	sds_provisioner "github.com/stephenzsy/doc-locker/server/sds/provisioners"
	sds_provisioner_azure "github.com/stephenzsy/doc-locker/server/sds/provisioners/azure"
)

func getWhitelistedSdsSecretFromString(str string) (secretType configurations.SecretType, secretName configurations.SecretName, err error) {
	if err != nil {
		return
	}
	if str == fmt.Sprintf("%s-%s", configurations.SecretTypeServer, configurations.SecretNameProxy) {
		return configurations.SecretTypeServer, configurations.SecretNameProxy, nil
	}
	err = errors.New("Invalid secret name for SDS: " + str)
	return
}

type server struct {
	certProvisioner sds_provisioner.CertificatesProvisioner
}

func NewServer(ctx context.Context) (s server, err error) {
	certProvisioner, err := sds_provisioner_azure.NewAzureCertificatesProvisioner()
	if err != nil {
		return
	}
	s = server{
		certProvisioner: certProvisioner,
	}
	return
}

func (*server) DeltaSecrets(_ secretservice.SecretDiscoveryService_DeltaSecretsServer) error {
	log.Fatal("Delta Secrets")
	return errors.New("not implemented")
}

type certsEntry struct {
	certsChain []*pem.Block
	privateKey *pem.Block
}

func (s *server) StreamSecrets(stream secretservice.SecretDiscoveryService_StreamSecretsServer) (err error) {
	ctx := stream.Context()
	errCh := make(chan error)
	reqCh := make(chan *discovery.DiscoveryRequest)

	go func() {
		for {
			r, err := stream.Recv()
			if err != nil {
				errCh <- err
				return
			}
			if err := s.validateRequest(ctx, r); err != nil {
				errCh <- err
				return
			}
			reqCh <- r
		}
	}()

	var nonce, versionInfo string

	for {
		var req *discovery.DiscoveryRequest
		select {
		case req = <-reqCh:
			//			isRenewal = false

			// Validations
			if req.ErrorDetail != nil {
				//				srv.logRequest(ctx, r, "NACK", t1, nil)
				continue
			}
			// Do not validate nonce/version if we're restarting the server
			if req != nil {
				log.Print(req)
				switch {
				case nonce != req.ResponseNonce:
					// srv.logRequest(ctx, r, "Invalid responseNonce", t1, fmt.Errorf("invalid responseNonce"))
					continue
				case req.VersionInfo == "": // initial request
					versionInfo = s.versionInfo()
				case req.VersionInfo == versionInfo: // ACK
					//srv.logRequest(ctx, r, "ACK", t1, nil)
					continue
				default: // it should not go here
					versionInfo = s.versionInfo()
				}
			} else {
				versionInfo = s.versionInfo()
			}

		case err = <-errCh:
			return
		}

		entries := make([]certsEntry, 0, len(req.ResourceNames))

		for _, resourceName := range req.ResourceNames {
			secretType, secretName, ee := getWhitelistedSdsSecretFromString(resourceName)
			if ee != nil {
				return ee
			}

			certsChain, privateKey, ee := s.certProvisioner.FetchCertificateWithPrivateKey(ctx, secretType, secretName)
			if ee != nil {
				return ee
			}
			entries = append(entries, certsEntry{
				certsChain: certsChain,
				privateKey: privateKey,
			})
		}

		// Send certificates
		response, err := getDiscoveryResponse(req, versionInfo, entries)
		if err != nil {
			//	srv.logRequest(ctx, req, "Creation of DiscoveryResponse failed", t1, err)
			return err
		}
		if err := stream.Send(response); err != nil {
			//	srv.logRequest(ctx, req, "Send failed", t1, err)
			return err
		}
		nonce = response.GetNonce()
	}
}

func (*server) FetchSecrets(_ context.Context, _ *discovery.DiscoveryRequest) (*discovery.DiscoveryResponse, error) {
	log.Fatal("Fetch Secrets")
	return nil, errors.New("not implemented")
}

func getDiscoveryResponse(req *discovery.DiscoveryRequest, versionInfo string, entries []certsEntry) (response *discovery.DiscoveryResponse, err error) {
	nonce, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("error generating nonce: %w", err)
	}

	var resources []*any.Any
	for i, name := range req.ResourceNames {
		b, err := getSecret(name, entries[i])
		if err != nil {
			return nil, err
		}
		resources = append(resources, &any.Any{
			TypeUrl: resourcev3.SecretType,
			Value:   b,
		})
	}

	return &discovery.DiscoveryResponse{
		VersionInfo: versionInfo,
		Resources:   resources,
		Canary:      false,
		TypeUrl:     req.TypeUrl,
		Nonce:       nonce.String(),
		ControlPlane: &core.ControlPlane{
			Identifier: "hahaha",
		},
	}, nil
}

func getSecret(name string, entry certsEntry) (value []byte, err error) {
	b := proto.NewBuffer(nil)
	b.SetDeterministic(true)

	privateKeyBytes := pem.EncodeToMemory(entry.privateKey)
	certChainBytesBuffer := bytes.Buffer{}
	for _, certBlock := range entry.certsChain {
		certChainBytesBuffer.Write(pem.EncodeToMemory(certBlock))
	}
	secret := secrets.Secret{
		Name: name,
		Type: &secrets.Secret_TlsCertificate{
			TlsCertificate: &secrets.TlsCertificate{
				CertificateChain: &core.DataSource{
					Specifier: &core.DataSource_InlineBytes{
						InlineBytes: certChainBytesBuffer.Bytes(),
					},
				},
				PrivateKey: &core.DataSource{
					Specifier: &core.DataSource_InlineBytes{
						InlineBytes: privateKeyBytes,
					},
				},
			},
		},
	}

	err = b.Marshal(&secret)
	if err != nil {
		return
	}
	return b.Bytes(), err
}

// TODO
func (*server) validateRequest(context.Context, *discovery.DiscoveryRequest) error {
	return nil
}

func (srv *server) versionInfo() string {
	return time.Now().UTC().Format(time.RFC3339)
}

func (srv *server) logRequest(ctx context.Context, r *discovery.DiscoveryRequest, msg string, start time.Time, err error, extra ...interface{}) {
}
