package configurations

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"time"
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

type deploymentCloudConfiguration struct {
	AzureServicePrincipalThumbprint       HexString `json:"azureServicePrincipalThumbprint"`
	AzureServicePrincipalCertificateChain [][]byte  `json:"azureServicePrincipalCertificateChain"`
}

type deploymentConfiguration struct {
	Version     uint                         `json:"version"` // schema version
	LastUpdated time.Time                    `json:"lastUpdated"`
	Cloud       deploymentCloudConfiguration `json:"cloud"`
}

type DeploymentConfigurationFile struct {
	deploymentConfiguration
	EncryptionMaterial struct {
		Key []byte `json:"key"`
	} `json:"encryptedEncryptionKey"`
	CanonicalHash       []byte `json:"canonicalHash"`
	CanonicalSecretHash []byte `json:"canonicalSecretHash"`
	SigningThumbprint   []byte `json:"signingThumbprint"`
	Siganature          []byte `json:"signature"`
	SecretSiganature    []byte `json:"secretSignature"`
}

type DeploymentCloudConfigurationUnlocked struct {
	deploymentCloudConfiguration
	AzureServicePrincipalPrivateKey []byte
}

type DeploymentConfigurationUnlocked struct {
	deploymentConfiguration
	Cloud DeploymentCloudConfigurationUnlocked
}

func newDeploymentConfiguration(configDir string) (*DeploymentConfigurationFile, error) {
	data := DeploymentConfigurationFile{}
	e := loadConfigFromFile(path.Join(configDir, "server", "deployment.json"), &data)
	if e == nil || os.IsNotExist(e) {
		return &data, nil
	}
	return nil, e
}

func (c *DeploymentConfigurationFile) Save(configDir string) error {
	marshalled, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path.Join(configDir, "server", "deployment.json"), marshalled, 0644)
}
