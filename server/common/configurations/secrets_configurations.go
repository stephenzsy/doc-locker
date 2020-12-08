package configurations

import (
	"encoding/json"
	"fmt"
	"path"
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
	SecretNameDeploy                      SecretName = "deploy"
	SecretNameDeployAzureServicePrincipal SecretName = "deploy-azure-service-principal"
	SecretNameDeploySds                   SecretName = "deploy-sds"
	SecretNameProxy                       SecretName = "proxy"
	SecretNameApi                         SecretName = "api"
	SecretNameSite                        SecretName = "site"
)

func (s *SecretName) FromString(str string) error {
	a := (*string)(s)
	*a = str
	// Validate the valid enum values
	switch *s {
	case
		SecretNameDeploy,
		SecretNameDeployAzureServicePrincipal,
		SecretNameDeploySds,
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

type SecretsConfiguration struct {
	configDir string
}

func (c *SecretsConfiguration) GetCaPath(caRole CaRole) string {
	return path.Join(c.configDir, "certs", fmt.Sprintf("%s-%s.pem", SecretTypeCa, caRole))
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
