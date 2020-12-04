package configurations

import (
	"errors"
	"fmt"
	"path"
)

type SecretType string

const (
	SecretTypeClient SecretType = "client"
	SecretTypeServer            = "server"
)

type SecretName string

const (
	SecretNameDeploy                      SecretName = "deploy"
	SecretNameDeployAzureServicePrincipal            = "deploy-azure-service-principal"
	SecretNameDeploySds                              = "deploy-sds"
)

type SdsSecretName string

const (
	sdsSecretNameUnknown     SdsSecretName = "unknown"
	SdsSecretNameProxyServer SdsSecretName = "proxy_server"
)

func SdsSecretNameFromString(str string) (SdsSecretName, error) {
	switch str {
	case string(SdsSecretNameProxyServer):
		return SdsSecretNameProxyServer, nil
	}
	return sdsSecretNameUnknown, errors.New("Invalid SdsSecretName: " + str)
}

type SecretsConfiguration struct {
	configDir string
}

func (c *SecretsConfiguration) GetCertPath(secretType SecretType, secretName SecretName) string {
	return path.Join(c.configDir, "certs", fmt.Sprintf("%s-cert-%s.pem", secretType, secretName))
}

func (c *SecretsConfiguration) GetPrivateKeyPath(secretType SecretType, secretName SecretName) string {
	return path.Join(c.configDir, "certsk", fmt.Sprintf("%s-key-%s.pem", secretType, secretName))
}

func (c *SecretsConfiguration) GetKeyPairPath(secretName SecretName) string {
	return path.Join(c.configDir, "keypairs", fmt.Sprintf("key-%s.pem", secretName))
}
