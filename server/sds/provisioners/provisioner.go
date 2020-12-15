package provisioners

import (
	"encoding/pem"

	"github.com/stephenzsy/doc-locker/server/common/app_context"
	"github.com/stephenzsy/doc-locker/server/common/configurations"
)

type CertificatesProvisioner interface {
	FetchCertificateWithPrivateKey(app_context.AppContext, configurations.SecretType, configurations.SecretName) (certificates []*pem.Block, privateKey *pem.Block, err error)
}
