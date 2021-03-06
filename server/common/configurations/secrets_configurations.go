package configurations

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"path"

	"github.com/stephenzsy/doc-locker/server/common/app_context"
	"github.com/stephenzsy/doc-locker/server/common/auth"
	"github.com/stephenzsy/doc-locker/server/common/crypto_utils"
)

type SecretType string

const (
	SecretTypeCa      SecretType = "ca"
	SecretTypeClient  SecretType = "client"
	SecretTypeServer  SecretType = "server"
	SecretTypeKeyPair SecretType = "key-pair"
)

func (s *SecretType) FromString(str string) error {
	a := (*string)(s)
	*a = str
	// Validate the valid enum values
	switch *s {
	case
		SecretTypeCa,
		SecretTypeClient,
		SecretTypeServer,
		SecretTypeKeyPair:
		return nil
	default:
		return fmt.Errorf("invalid value for SecretType: %s", *a)
	}
}

func (s *SecretType) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	return s.FromString(str)
}

type SecretName string

const (
	SecretNameDeploy                         SecretName = "deploy"
	SecretNameDeploySdsAzureServicePrincipal SecretName = "deploy-sds-asp"
	SecretNameDeploySds                      SecretName = "deploy-sds"
	SecretNameDeploySdsEnvoy                 SecretName = "deploy-sds-envoy"
	SecretNameProxy                          SecretName = "proxy"
	SecretNameApi                            SecretName = "api"
	SecretNameSite                           SecretName = "site"
)

func (s *SecretName) FromString(str string) error {
	a := (*string)(s)
	*a = str
	// Validate the valid enum values
	switch *s {
	case
		SecretNameDeploy,
		SecretNameDeploySdsAzureServicePrincipal,
		SecretNameDeploySds,
		SecretNameDeploySdsEnvoy,
		SecretNameProxy,
		SecretNameApi,
		SecretNameSite:
		return nil
	default:
		return fmt.Errorf("invalid value for SecretName: %s", *a)
	}
}

func (s *SecretName) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	return s.FromString(str)
}

type SecretsConfiguration interface {
	GetCaPath(caRole CaRole) string
	GetCertPath(secretType SecretType, secretName SecretName) string
	GetPrivateKeyPath(secretType SecretType, secretName SecretName) string
	GetECPrivateKey(secretsType SecretType, secretName SecretName) (*ecdsa.PrivateKey, error)
	GetRsaPrivateKey(secretsType SecretType, secretName SecretName) (*rsa.PrivateKey, error)
	GetCertificate(secretsType SecretType, secretName SecretName) (*x509.Certificate, error)
	GetCertificateChain(secretsType SecretType, secretName SecretName) ([]*x509.Certificate, error)
}

type secretsConfiguration struct {
	configDir string
}

func (c secretsConfiguration) GetCaPath(caRole CaRole) string {
	return path.Join(c.configDir, "certs", fmt.Sprintf("%s-%s.pem", SecretTypeCa, caRole))
}

func (c secretsConfiguration) GetCertPath(secretType SecretType, secretName SecretName) string {
	return path.Join(c.configDir, "certs", fmt.Sprintf("%s-cert-%s.pem", secretType, secretName))
}

func (c secretsConfiguration) GetPrivateKeyPath(secretType SecretType, secretName SecretName) string {
	return path.Join(c.configDir, "certsk", fmt.Sprintf("%s-key-%s.pem", secretType, secretName))
}

func (c secretsConfiguration) GetECPrivateKey(secretsType SecretType, secretName SecretName) (*ecdsa.PrivateKey, error) {
	filePath := c.GetPrivateKeyPath(secretsType, secretName)
	return crypto_utils.ParseECPrivateKeyFromPemFile(filePath)
}

func (c secretsConfiguration) GetRsaPrivateKey(secretsType SecretType, secretName SecretName) (*rsa.PrivateKey, error) {
	filePath := c.GetPrivateKeyPath(secretsType, secretName)
	return crypto_utils.ParseRsaPrivateKeyFromPemFile(filePath)
}

func (c secretsConfiguration) GetCertificate(secretsType SecretType, secretName SecretName) (*x509.Certificate, error) {
	filePath := c.GetCertPath(secretsType, secretName)
	return crypto_utils.ParseCertificateFromPemFile(filePath)
}

func (c secretsConfiguration) GetCertificateChain(secretsType SecretType, secretName SecretName) ([]*x509.Certificate, error) {
	filePath := c.GetCertPath(secretsType, secretName)
	certificates, err := crypto_utils.ParseCertificateChainFromPemFile(filePath)
	if err != nil {
		return certificates, fmt.Errorf("Failed to parse certificate chain from file %s: %v", filePath, err)
	}
	return certificates, err
}

func GetSecretsConfiguration(ctx app_context.AppContext) (config SecretsConfiguration, err error) {

	if err = app_context.VerifyElevated(ctx); err != nil {
		return
	}
	if err = app_context.VerifyCallerId(ctx, auth.SystemCallerIdBootstrap, auth.ServiceCallerIdSds); err != nil {
		return
	}

	configDir, err := GetConfigurationsDir(ctx)
	if err != nil {
		return
	}
	config = secretsConfiguration{
		configDir: configDir,
	}

	return
}
