package sds

import (
	"context"
	"errors"
	"log"
	"time"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	secretservice "github.com/envoyproxy/go-control-plane/envoy/service/secret/v3"
	"github.com/stephenzsy/doc-locker/server/common/configurations"
	sds_provisioner "github.com/stephenzsy/doc-locker/server/sds/provisioners"
	sds_provisioner_azure "github.com/stephenzsy/doc-locker/server/sds/provisioners/azure"
)

type server struct {
	certProvisioner sds_provisioner.CertificatesProvisioner
}

func NewServer(ctx context.Context) server {
	s := server{
		certProvisioner: sds_provisioner_azure.NewAzureCertificatesProvisioner(),
	}
	return s
}

func (*server) DeltaSecrets(_ secretservice.SecretDiscoveryService_DeltaSecretsServer) error {
	log.Fatal("Delta Secrets")
	return errors.New("not implemented")
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
		select {
		case r := <-reqCh:
			//			isRenewal = false

			// Validations
			if r.ErrorDetail != nil {
				//				srv.logRequest(ctx, r, "NACK", t1, nil)
				continue
			}
			// Do not validate nonce/version if we're restarting the server
			if r != nil {
				switch {
				case nonce != r.ResponseNonce:
					// srv.logRequest(ctx, r, "Invalid responseNonce", t1, fmt.Errorf("invalid responseNonce"))
					continue
				case r.VersionInfo == "": // initial request
					versionInfo = s.versionInfo()
				case r.VersionInfo == versionInfo: // ACK
					//srv.logRequest(ctx, r, "ACK", t1, nil)
					continue
				default: // it should not go here
					versionInfo = s.versionInfo()
				}
			} else {
				versionInfo = s.versionInfo()
			}

			for _, name := range r.ResourceNames {

				secretName, err := configurations.SdsSecretNameFromString(name)
				if err != nil {
					errCh <- err
				}

				err = s.certProvisioner.FetchCertificateWithPrivateKey(ctx, secretName)
			}
		case err := <-errCh:
			return err
		}
	}
}

func (*server) FetchSecrets(_ context.Context, _ *discovery.DiscoveryRequest) (*discovery.DiscoveryResponse, error) {
	log.Fatal("Fetch Secrets")
	return nil, errors.New("not implemented")
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
