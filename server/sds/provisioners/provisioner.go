package provisioners

import (
	"context"
	"encoding/pem"

	"github.com/stephenzsy/doc-locker/server/common/configurations"
)

type CertificatesProvisioner interface {
	FetchCertificateWithPrivateKey(context.Context, configurations.SecretType, configurations.SecretName) (certificates []*pem.Block, privateKey *pem.Block, err error)
}
