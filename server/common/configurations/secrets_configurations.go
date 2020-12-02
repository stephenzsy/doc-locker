package configurations

import (
	"fmt"
	"path"
)

type SecretName string

type SecretType string

const (
	SecretTypeClient SecretType = "client"
	SecretTypeServer            = "server"
)

const (
	SecretNameDeploy                      SecretName = "deploy"
	SecretNameDeployAzureServicePrincipal            = "deploy-azure-service-principal"
	SecretNameDeploySds                              = "deploy-sds"
)

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
