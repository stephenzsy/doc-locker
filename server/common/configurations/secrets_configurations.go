package configurations

import (
	"errors"
	"fmt"
	"path"
)

type SecretType string

const (
	SecretTypeClient SecretType = "client"
	SecretTypeServer SecretType = "server"
)

func SecretTypeFromString(str string) (secretType SecretType, err error) {
	switch str {
	case string(SecretTypeClient):
		return SecretTypeClient, nil
	case string(SecretTypeServer):
		return SecretTypeServer, nil
	}
	err = errors.New("Invalid SecretType value: " + str)
	return
}

type SecretName string

const (
	SecretNameDeploy                      SecretName = "deploy"
	SecretNameDeployAzureServicePrincipal SecretName = "deploy-azure-service-principal"
	SecretNameDeploySds                   SecretName = "deploy-sds"
	SecretNameProxy                       SecretName = "proxy"
)

func SecretNameFromString(str string) (secretName SecretName, err error) {
	switch str {
	case string(SecretNameProxy):
		return SecretNameProxy, nil
	case string(SecretNameDeploySds):
		return SecretNameDeploySds, nil
	}
	err = errors.New("Invalid SecretType value: " + str)
	return
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
