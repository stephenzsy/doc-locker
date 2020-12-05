package configurations

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"sync"
	"time"

	"github.com/stephenzsy/doc-locker/server/common/security"
)

type DeploymentConfiguration interface {
	GetVersion() uint
}

type DeploymentCloudConfigurationAzurePublic struct {
	ServicePrincipalThumbprint       HexString `json:"servicePrincipalThumbprint"`
	ServicePrincipalCertificateChain [][]byte  `json:"servicePrincipalCertificateChain"`
	KeyVaultBaseUrl                  string    `json:"keyVaultBaseUrl"`
}

type DeploymentCloudConfigurationPublic struct {
	Azure DeploymentCloudConfigurationAzurePublic `json:"azure"`
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
	lockingMu           sync.Mutex
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

func newDeploymentConfiguration(configDir string) (*DeploymentConfigurationFile, error) {
	data := DeploymentConfigurationFile{}
	e := loadConfigFromFile(path.Join(configDir, "server", "deployment.json"), &data)
	if e == nil || os.IsNotExist(e) {
		return &data, nil
	}
	return nil, e
}

func (c *DeploymentConfigurationFile) GetPrivateConfig(privateKey *rsa.PrivateKey) (privateConfig DeploymentConfigurationPrivate, encryptionKey []byte, err error) {
	c.lockingMu.Lock()
	defer c.lockingMu.Unlock()

	encryptionKey, err = rsa.DecryptOAEP(sha512.New384(), rand.Reader, privateKey, c.EncryptionMaterial.EncryptedKey, []byte{})
	if err != nil {
		return
	}
	if c.Secret != nil {
		var decrypted []byte
		decrypted, err = security.AESDecrypt(encryptionKey, &c.Secret)
		if err != nil {
			return
		}
		err = json.Unmarshal(decrypted, &privateConfig)
	}

	return privateConfig, encryptionKey, err
}

func (c *DeploymentConfigurationFile) SetPrivateConfig(publicKey *rsa.PublicKey, encryptionKey []byte, privateConfig DeploymentConfigurationPrivate) (err error) {
	c.lockingMu.Lock()
	defer c.lockingMu.Unlock()
	content, err := json.Marshal(privateConfig)
	if err != nil {
		return
	}

	c.EncryptionMaterial.EncryptedKey, err = rsa.EncryptOAEP(sha512.New384(), rand.Reader, publicKey, encryptionKey, []byte{})
	if err != nil {
		return
	}

	c.Secret, err = security.AESEncrypt(encryptionKey, &content)
	return
}

func (c *DeploymentConfigurationFile) Save(configDir string) error {
	marshalled, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path.Join(configDir, "server", "deployment.json"), marshalled, 0644)
}
