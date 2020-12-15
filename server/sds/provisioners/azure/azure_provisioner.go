package azure

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/v7.1/keyvault"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/stephenzsy/doc-locker/server/common/app_context"
	"github.com/stephenzsy/doc-locker/server/common/configurations"
	"github.com/stephenzsy/doc-locker/server/common/crypto_utils"
)

func getAzureKeyVaultCertificateName(secretType configurations.SecretType, secretName configurations.SecretName) (string, error) {
	if secretType == configurations.SecretTypeServer && secretName == configurations.SecretNameProxy {
		return "proxy-server", nil
	}
	return "", errors.New(fmt.Sprintf("Certificate name not valid: %s, %s", secretType, secretName))
}

func getServicePrincipalToken(
	ctx app_context.AppContext,
	configs configurations.DeploymentConfigurationFile) (token *adal.ServicePrincipalToken, err error) {

	oauthConfig, err := adal.NewOAuthConfig(configs.Cloud.Azure.AadOauthEndpoint, configs.Cloud.Azure.AadTenantId)
	secretsConfig, err := configurations.GetSecretsConfiguration(ctx)
	if err != nil {
		return
	}
	privateKey, err := secretsConfig.GetRsaPrivateKey(configurations.SecretTypeKeyPair, configurations.SecretNameDeploy)
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

	servicePrincipalCertificate, err := x509.ParseCertificate(configs.Cloud.Azure.ServicePrincipalCertificate)
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

func getClient(
	ctx app_context.AppContext,
	configs configurations.DeploymentConfigurationFile) (keyvault.BaseClient, error) {
	kvClient := keyvault.New()
	spt, err := getServicePrincipalToken(ctx, configs)
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

func NewAzureCertificatesProvisioner(ctx app_context.AppContext) (p *provisioner, err error) {
	configs, err := configurations.GetServerDeploymentConfigurationFile(ctx)
	if err != nil {
		return
	}
	kvClient, err := getClient(ctx, configs)
	p = &provisioner{
		client:       kvClient,
		vaultBaseUrl: configs.Cloud.Azure.KeyVaultBaseUrl,
	}
	return
}

func (p *provisioner) FetchCertificateWithPrivateKey(ctx app_context.AppContext,
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
	ctx app_context.AppContext,
	secretType configurations.SecretType,
	secretName configurations.SecretName) (importedCertBundle keyvault.CertificateBundle, err error) {

	configs, err := configurations.GetSecretsConfiguration(ctx)
	if err != nil {
		return
	}

	certificates, err := configs.GetCertificateChain(secretType, secretName)
	if err != nil {
		return
	}
	privateKey, err := configs.GetECPrivateKey(secretType, secretName)
	if err != nil {
		return
	}
	privateKeyPemBytes, err := crypto_utils.MarshalPKCS8PrivateKeyPemBlock(privateKey)
	if err != nil {
		return
	}

	certBundle := string(append(
		privateKeyPemBytes,
		crypto_utils.MarshalCertificatesPemBlock(certificates...)...))
	contentType := "application/x-pem-file"
	exportable := true
	azureKeyVaultCertificateName, err := getAzureKeyVaultCertificateName(secretType, secretName)
	if err != nil {
		return
	}
	return p.client.ImportCertificate(ctx, p.vaultBaseUrl, azureKeyVaultCertificateName,
		keyvault.CertificateImportParameters{
			Base64EncodedCertificate: &certBundle,
			CertificatePolicy: &keyvault.CertificatePolicy{
				KeyProperties: &keyvault.KeyProperties{
					Exportable: &exportable,
					KeyType:    keyvault.EC,
					Curve:      keyvault.P256,
				},
				SecretProperties: &keyvault.SecretProperties{
					ContentType: &contentType,
				},
			},
		})
}
