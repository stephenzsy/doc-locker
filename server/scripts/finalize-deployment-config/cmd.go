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
	"path"

	"github.com/stephenzsy/doc-locker/server/common/configurations"
	"github.com/stephenzsy/doc-locker/server/common/security"
)

func UpdateEncryptionMaterial(c *configurations.DeploymentConfigurationFile, configRootDir string) (*rsa.PrivateKey, error) {
	certContent, err := ioutil.ReadFile(path.Join(configRootDir, "certsk", "key-cert-deploy-cipher.pem"))
	privateKeyPemBlock, certContent := pem.Decode(certContent)
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyPemBlock.Bytes)
	encryptionKey, err := security.GenerateAes256Key()
	c.EncryptionMaterial.EncryptedKey, err = rsa.EncryptOAEP(sha512.New384(), rand.Reader, &privateKey.PublicKey, encryptionKey, []byte{})
	return privateKey, err
}

func UpdateAzureServicePrincipal(
	c *configurations.DeploymentConfigurationFile,
	configRootDir string,
	privateKey *rsa.PrivateKey) error {
	certContent, err := ioutil.ReadFile(path.Join(configRootDir, "certs", "azure-service-principal-deploy.pem"))
	if err != nil {
		return err
	}
	pemBlock, certContent := pem.Decode(certContent)
	certificate, err := x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		return err
	}
	thumbprint := sha1.Sum(certificate.Raw)
	c.Cloud.AzureServicePrincipalThumbprint = thumbprint[:]
	var certificatesChain []configurations.Base64String
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
	c.Cloud.AzureServicePrincipalCertificateChain = certificatesChain

	// private configuration

	privateConfig, encryptionKey, err := c.GetPrivateConfig(privateKey)
	if err != nil {
		return err
	}

	servicePrincipalPrivateKeyContent, err := ioutil.ReadFile(path.Join(configRootDir, "setup", "azure-service-principal-deploy-private-key.pem"))
	if err != nil {
		return err
	}
	pemBlock, certContent = pem.Decode(servicePrincipalPrivateKeyContent)
	_, err = x509.ParseECPrivateKey(pemBlock.Bytes)
	if err != nil {
		return err
	}
	privateConfig.Cloud.AzureServicePrincipalPrivateKey = pemBlock.Bytes
	c.SetPrivateConfig(&privateKey.PublicKey, encryptionKey, privateConfig)

	return nil
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
		err = UpdateAzureServicePrincipal(deploymentConfig, configRootDir, privateKey)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = deploymentConfig.Save(configRootDir)
	if err != nil {
		log.Fatal(err)
	}
}
