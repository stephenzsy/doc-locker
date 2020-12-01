package configurations

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
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

type HexString []byte

func (s *HexString) MarshalJSON() ([]byte, error) {
	dst := make([]byte, hex.EncodedLen(len(*s)))
	buffer := bytes.NewBufferString(`"`)
	hex.Encode(dst, *s)
	buffer.Write(bytes.ToUpper(dst))
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

type Base64String []byte

func (s *Base64String) MarshalJSON() ([]byte, error) {
	dst := make([]byte, base64.RawStdEncoding.EncodedLen(len(*s)))
	buffer := bytes.NewBufferString(`"`)
	base64.RawStdEncoding.Encode(dst, *s)
	buffer.Write(dst)
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

type DeploymentCloudConfigurationPublic struct {
	AzureServicePrincipalThumbprint       HexString      `json:"azureServicePrincipalThumbprint"`
	AzureServicePrincipalCertificateChain []Base64String `json:"azureServicePrincipalCertificateChain"`
}

type DeploymentConfigurationFile struct {
	Version            uint                               `json:"version"` // schema version
	LastUpdated        time.Time                          `json:"lastUpdated"`
	Cloud              DeploymentCloudConfigurationPublic `json:"cloud"`
	EncryptionMaterial struct {
		EncryptedKey HexString `json:"encryptedKey"`
	} `json:"encryptionMaterial"`
	Secret              Base64String `json:"secret"`
	CanonicalHash       []byte       `json:"canonicalHash"`
	CanonicalSecretHash []byte       `json:"canonicalSecretHash"`
	SigningThumbprint   []byte       `json:"signingThumbprint"`
	Siganature          []byte       `json:"signature"`
	SecretSiganature    []byte       `json:"secretSignature"`
	lockingMu           sync.Mutex
}

type DeploymentCloudConfigurationPrivate struct {
	AzureServicePrincipalPrivateKey Base64String `json:"azureServicePrincipalPrivateKey"`
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

func (c *DeploymentConfigurationFile) GetPrivateConfig(privateKey *rsa.PrivateKey) (DeploymentConfigurationPrivate, []byte, error) {
	c.lockingMu.Lock()
	defer c.lockingMu.Unlock()

	privateConfig := DeploymentConfigurationPrivate{}

	encryptionKey, err := rsa.DecryptOAEP(sha512.New384(), rand.Reader, privateKey, c.EncryptionMaterial.EncryptedKey, []byte{})
	if c.Secret != nil {
		cipherBlock, err := aes.NewCipher(encryptionKey)
		if err != nil {
			return privateConfig, nil, err
		}
		cipherText, iv := c.Secret[:aes.BlockSize], c.Secret[aes.BlockSize:]
		decrypter := cipher.NewCBCDecrypter(cipherBlock, iv)
		encrypted := make([]byte, len(cipherText))
		decrypter.CryptBlocks(encrypted, cipherText)
		err = json.Unmarshal(encrypted, &privateConfig)
	}

	return privateConfig, encryptionKey, err
}

func (c *DeploymentConfigurationFile) SetPrivateConfig(publicKey *rsa.PublicKey, encryptionKey []byte, privateConfig DeploymentConfigurationPrivate) error {
	c.lockingMu.Lock()
	defer c.lockingMu.Unlock()
	content, err := json.Marshal(privateConfig)
	if err != nil {
		return err
	}

	c.EncryptionMaterial.EncryptedKey, err = rsa.EncryptOAEP(sha512.New384(), rand.Reader, publicKey, encryptionKey, []byte{})
	cipherBlock, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return err
	}
	iv, err := security.GenerateAesIv()
	if err != nil {
		return err
	}

	encrypter := cipher.NewCBCEncrypter(cipherBlock, iv)
	encrypted := make([]byte, security.GetAesCipherTextLength(content))
	encrypter.CryptBlocks(encrypted, content)
	c.Secret = append(iv, encrypted...)
	return nil
}

func (c *DeploymentConfigurationFile) Save(configDir string) error {
	marshalled, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path.Join(configDir, "server", "deployment.json"), marshalled, 0644)
}
