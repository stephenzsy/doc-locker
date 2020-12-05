package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"

	"github.com/stephenzsy/doc-locker/server/common/configurations"
	"github.com/stephenzsy/doc-locker/server/common/security"
)

func UpdateEncryptionMaterial(c *configurations.DeploymentConfigurationFile, configRootDir string) (*rsa.PrivateKey, error) {
	certContent, err := ioutil.ReadFile(
		configurations.Configurations().SecretsConfiguration().GetKeyPairPath(configurations.SecretNameDeploy))
	privateKeyPemBlock, certContent := pem.Decode(certContent)
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyPemBlock.Bytes)
	encryptionKey, err := security.GenerateAes256Key()
	c.EncryptionMaterial.EncryptedKey, err = rsa.EncryptOAEP(sha512.New384(), rand.Reader, &privateKey.PublicKey, encryptionKey, []byte{})
	return privateKey, err
}

func updateAzureConfigurations(
	c *configurations.DeploymentConfigurationFile,
	configRootDir string,
	privateKey *rsa.PrivateKey) (err error) {

	serverSetupConfig, err := configurations.Configurations().ServerSetup()
	if err != nil {
		return
	}
	c.Cloud.Azure.ServerSetupCloudAzureConfiguration = serverSetupConfig.Cloud.Azure

	// service principal
	secretsConfiguration := configurations.Configurations().SecretsConfiguration()

	certContent, err := ioutil.ReadFile(
		secretsConfiguration.GetCertPath(
			configurations.SecretTypeClient,
			configurations.SecretNameDeployAzureServicePrincipal))
	if err != nil {
		return
	}
	pemBlock, certContent := pem.Decode(certContent)
	certificate, err := x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		return
	}
	thumbprint := sha1.Sum(certificate.Raw)
	c.Cloud.Azure.ServicePrincipalThumbprint = thumbprint[:]
	var certificatesChain [][]byte
	for certificate != nil {
		certificatesChain = append(certificatesChain, certificate.Raw)
		pemBlock, certContent = pem.Decode(certContent)
		if pemBlock == nil {
			break
		}
		certificate, err = x509.ParseCertificate(pemBlock.Bytes)
		if err != nil {
			return err
		}
	}
	c.Cloud.Azure.ServicePrincipalCertificateChain = certificatesChain

	// private configuration

	privateConfig, encryptionKey, err := c.GetPrivateConfig(privateKey)
	if err != nil {
		return err
	}

	servicePrincipalPrivateKeyContent, err := ioutil.ReadFile(
		secretsConfiguration.GetPrivateKeyPath(
			configurations.SecretTypeClient,
			configurations.SecretNameDeployAzureServicePrincipal))
	if err != nil {
		return err
	}
	pemBlock, certContent = pem.Decode(servicePrincipalPrivateKeyContent)
	_, err = x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
	if err != nil {
		return err
	}
	privateConfig.Cloud.Azure.ServicePrincipalPrivateKey = pemBlock.Bytes
	c.SetPrivateConfig(&privateKey.PublicKey, encryptionKey, privateConfig)

	c.Version = 1
	return
}

func main() {
	configs := configurations.Configurations()
	configRootDir := configs.ConfigRootDir()
	deploymentConfig, err := configs.Deployment()
	if err != nil {
		log.Fatal(err)
	}
	if deploymentConfig.Version == 0 {
		// no configuration file
		privateKey, err := UpdateEncryptionMaterial(deploymentConfig, configRootDir)
		if err != nil {
			log.Fatal(err)
		}
		err = updateAzureConfigurations(deploymentConfig, configRootDir, privateKey)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = deploymentConfig.Save(configRootDir)
	if err != nil {
		log.Fatal(err)
	}
}
