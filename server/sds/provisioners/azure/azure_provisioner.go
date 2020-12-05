package azure

import (
	"context"
	"crypto/x509"
	"errors"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
	"github.com/stephenzsy/doc-locker/server/common/configurations"
)

var sdsToAzureCertName = map[configurations.SdsSecretName]string{
	configurations.SdsSecretNameProxyServer: "proxy",
}

func getClient() keyvault.BaseClient {
	keyClient := keyvault.New()
	return keyClient
}

type provisioner struct {
	client       keyvault.BaseClient
	vaultBaseUrl string
}

func NewAzureCertificatesProvisioner() (p *provisioner, err error) {
	configs, err := configurations.Configurations().Deployment()
	if err != nil {
		return
	}
	p = &provisioner{
		client:       getClient(),
		vaultBaseUrl: configs.Cloud.Azure.KeyVaultBaseUrl,
	}
	return
}

func (p *provisioner) FetchCertificateWithPrivateKey(ctx context.Context, name configurations.SdsSecretName) (err error) {
	result, err := p.client.GetCertificate(ctx, p.vaultBaseUrl, sdsToAzureCertName[name], "")
	if err != nil {
		return
	}
	if !(*result.Policy.KeyProperties.Exportable) {
		return errors.New("Certificate not exportable")
	}

	derBytes := result.Cer
	_, err = x509.ParseCertificate(*derBytes)
	if err != nil {
		return
	}
	return
}
