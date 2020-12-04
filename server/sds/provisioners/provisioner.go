package provisioners

import (
	"context"

	"github.com/stephenzsy/doc-locker/server/common/configurations"
)

type CertificatesProvisioner interface {
	FetchCertificateWithPrivateKey(context.Context, configurations.SdsSecretName) error
}
