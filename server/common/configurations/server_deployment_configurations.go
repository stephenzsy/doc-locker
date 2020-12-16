package configurations

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"github.com/stephenzsy/doc-locker/server/common/app_context"
	"github.com/stephenzsy/doc-locker/server/common/crypto_utils"
)

type DeploymentConfiguration interface {
	GetVersion() uint
}

type DeploymentCloudConfigurationAzurePublic struct {
	ServerSetupCloudAzureConfiguration
	ServicePrincipalThumbprint  HexString `json:"servicePrincipalThumbprint"`
	ServicePrincipalCertificate []byte    `json:"servicePrincipalCertificate"`
}

type DeploymentCloudConfigurationAwsPublic struct {
	ServerSetupCloudAwsConfiguration
}

type DeploymentCloudConfigurationPublic struct {
	Azure DeploymentCloudConfigurationAzurePublic `json:"azure"`
	Aws   DeploymentCloudConfigurationAwsPublic   `json:"aws"`
}

type DeploymentConfigurationFile struct {
	Version            uint                               `json:"version"` // schema version
	LastUpdated        time.Time                          `json:"lastUpdated"`
	Cloud              DeploymentCloudConfigurationPublic `json:"cloud"`
	EncryptionMaterial struct {
		EncryptedKey []byte `json:"encryptedKey"`
	} `json:"encryptionMaterial"`
	Secret              []byte `json:"secret"`
	CanonicalHash       []byte `json:"canonicalHash"`
	CanonicalSecretHash []byte `json:"canonicalSecretHash"`
	SigningThumbprint   []byte `json:"signingThumbprint"`
	Siganature          []byte `json:"signature"`
	SecretSiganature    []byte `json:"secretSignature"`
}

type DeploymentCloudConfigurationAzurePrivate struct {
	ServicePrincipalPrivateKey []byte `json:"servicePrincipalPrivateKey"`
}

type DeploymentCloudConfigurationPrivate struct {
	Azure DeploymentCloudConfigurationAzurePrivate `json:"azure"`
}

type DeploymentConfigurationPrivate struct {
	Cloud DeploymentCloudConfigurationPrivate `json:"cloud"`
}

func (c *DeploymentConfigurationFile) GetPrivateConfig(privateKey *rsa.PrivateKey) (privateConfig DeploymentConfigurationPrivate, encryptionKey []byte, err error) {
	encryptionKey, err = rsa.DecryptOAEP(sha512.New384(), rand.Reader, privateKey, c.EncryptionMaterial.EncryptedKey, []byte{})
	if err != nil {
		return
	}
	if c.Secret != nil {
		var decrypted []byte
		decrypted, err = crypto_utils.AESDecrypt(encryptionKey, &c.Secret)
		if err != nil {
			log.Panic(err)
			return
		}
		err = json.Unmarshal(decrypted, &privateConfig)
	}

	return privateConfig, encryptionKey, err
}

func (c *DeploymentConfigurationFile) SetPrivateConfig(publicKey *rsa.PublicKey, encryptionKey []byte, privateConfig DeploymentConfigurationPrivate) (err error) {
	content, err := json.Marshal(privateConfig)
	if err != nil {
		return
	}

	c.EncryptionMaterial.EncryptedKey, err = rsa.EncryptOAEP(sha512.New384(), rand.Reader, publicKey, encryptionKey, []byte{})
	if err != nil {
		return
	}

	c.Secret, err = crypto_utils.AESEncrypt(encryptionKey, &content)
	return
}

func (c *DeploymentConfigurationFile) Save(ctx app_context.AppContext) error {
	marshalled, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	configDir, err := GetConfigurationsDir(ctx)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path.Join(configDir, "server", "deployment.json"), marshalled, 0600)
}

func GetServerDeploymentConfigurationFile(ctx app_context.AppContext) (config DeploymentConfigurationFile, err error) {
	if err = app_context.VerifyElevated(ctx); err != nil {
		return
	}

	configDir, err := GetConfigurationsDir(ctx)
	if err != nil {
		return
	}
	err = loadConfigFromFile(path.Join(configDir, "server", "deployment.json"), &config)
	if os.IsNotExist(err) {
		err = nil
	}
	return
}
