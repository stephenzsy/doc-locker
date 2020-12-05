package azure

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/stephenzsy/doc-locker/server/common/configurations"
)

var sdsToAzureCertName = map[configurations.SdsSecretName]string{
	configurations.SdsSecretNameProxyServer: "proxy",
}

func getServicePrincipalToken(configs *configurations.DeploymentConfigurationFile) (token *adal.ServicePrincipalToken, err error) {

	oauthConfig, err := adal.NewOAuthConfig(configs.Cloud.Azure.AadOauthEndpoint, configs.Cloud.Azure.AadTenantId)
	privateKeyPath := configurations.Configurations().SecretsConfiguration().GetKeyPairPath(configurations.SecretNameDeploy)
	pemBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return
	}
	pemBlock, _ := pem.Decode(pemBytes)
	privateKey, err := x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
	if err != nil {
		return
	}

	privateConfig, _, err := configs.GetPrivateConfig(privateKey)
	if err != nil {
		return
	}

	servicePrincipalPrivateKey, err := x509.ParsePKCS1PrivateKey(privateConfig.Cloud.Azure.ServicePrincipalPrivateKey)
	if err != nil {
		return
	}

	servicePrincipalCertificatePemBytes := configs.Cloud.Azure.ServicePrincipalCertificateChain[0]
	servicePrincipalCertificate, err := x509.ParseCertificate(servicePrincipalCertificatePemBytes)
	if err != nil {
		return
	}

	return adal.NewServicePrincipalTokenFromCertificate(
		*oauthConfig,
		configs.Cloud.Azure.ApplicationId,
		servicePrincipalCertificate,
		servicePrincipalPrivateKey,
		configs.Cloud.Azure.KeyVaultResourceId)
}

func getClient(configs *configurations.DeploymentConfigurationFile) (keyvault.BaseClient, error) {
	kvClient := keyvault.New()
	spt, err := getServicePrincipalToken(configs)
	if err != nil {
		return kvClient, err
	}
	kvClient.Authorizer = autorest.NewBearerAuthorizer(spt)
	return kvClient, err
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
	kvClient, err := getClient(configs)
	p = &provisioner{
		client:       kvClient,
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
