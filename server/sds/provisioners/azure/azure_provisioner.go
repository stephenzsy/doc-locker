package azure

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/v7.1/keyvault"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/stephenzsy/doc-locker/server/common/configurations"
)

func getAzureKeyVaultCertificateName(secretType configurations.SecretType, secretName configurations.SecretName) (string, error) {
	if secretType == configurations.SecretTypeServer && secretName == configurations.SecretNameProxy {
		return "proxy-server", nil
	}
	return "", errors.New(fmt.Sprintf("Certificate name not valid: %s, %s", secretType, secretName))
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
		configs.Cloud.Azure.AadResourceKeyVault)
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

func (p *provisioner) FetchCertificateWithPrivateKey(ctx context.Context,
	secretType configurations.SecretType,
	secretName configurations.SecretName) (certificates []*pem.Block, privateKey *pem.Block, err error) {
	certificateName, err := getAzureKeyVaultCertificateName(secretType, secretName)
	if err != nil {
		return
	}
	result, err := p.client.GetSecret(ctx, p.vaultBaseUrl, certificateName, "")
	if err != nil {
		return
	}

	var rest []byte
	privateKey, rest = pem.Decode(([]byte)(*result.Value))
	if err != nil {
		return
	}
	certificates = make([]*pem.Block, 0, 3)
	for {
		var certPemBlock *pem.Block
		certPemBlock, rest = pem.Decode(rest)
		if certPemBlock == nil {
			break
		}
		certificates = append(certificates, certPemBlock)
	}

	return
}

func (p *provisioner) ImportCertificate(
	ctx context.Context,
	secretType configurations.SecretType,
	secretName configurations.SecretName) (importedCertBundle keyvault.SecretBundle, err error) {
	configs := configurations.Configurations().SecretsConfiguration()
	certBytes, err := ioutil.ReadFile(configs.GetCertPath(secretType, secretName))
	if err != nil {
		return
	}
	privateKey, err := ioutil.ReadFile(configs.GetPrivateKeyPath(secretType, secretName))
	if err != nil {
		return
	}
	certBundle := string(append(privateKey, certBytes...))
	contentType := "application/x-pem-file"
	azureKeyVaultCertificateName, err := getAzureKeyVaultCertificateName(secretType, secretName)
	if err != nil {
		return
	}
	return p.client.SetSecret(ctx, p.vaultBaseUrl, azureKeyVaultCertificateName,
		keyvault.SecretSetParameters{
			Value:       &certBundle,
			ContentType: &contentType,
		})
}
