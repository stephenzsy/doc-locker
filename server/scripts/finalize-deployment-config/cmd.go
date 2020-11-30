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

func UpdateEncryptionMaterial(c *configurations.DeploymentConfigurationFile, configRootDir string) error {
	certContent, err := ioutil.ReadFile(path.Join(configRootDir, "certsk", "key-cert-deploy-cipher.pem"))
	if err != nil {
		return err
	}
	privateKeyPemBlock, certContent := pem.Decode(certContent)
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyPemBlock.Bytes)
	if err != nil {
		return err
	}

	encryptionKey, err := security.GenerateAes256Key()
	if err != nil {
		return err
	}

	// crypto/rand.Reader is a good source of entropy for randomizing the
	// encryption function.
	rng := rand.Reader

	c.EncryptionMaterial.EncryptedKey, err = rsa.EncryptOAEP(sha512.New384(), rng, &privateKey.PublicKey, encryptionKey, []byte{})
	if err != nil {
		return err
	}

	return nil
}

func UpdateAzureServicePrincipal(c *configurations.DeploymentConfigurationFile, configRootDir string) error {
	certContent, err := ioutil.ReadFile(path.Join(configRootDir, "certs", "azure-service-principal-deploy-azure-service-principal.pem"))
	if err != nil {
		return err
	}
	pemBlock, _ := pem.Decode(certContent)

	certificate, err := x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		return err
	}
	thumbprint := sha1.Sum(certificate.Raw)
	c.Cloud.AzureServicePrincipalThumbprint = thumbprint[:]
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
		err = UpdateEncryptionMaterial(deploymentConfig, configRootDir)
		if err != nil {
			log.Fatal(err)
		}
		err = UpdateAzureServicePrincipal(deploymentConfig, configRootDir)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = deploymentConfig.Save(configRootDir)
	if err != nil {
		log.Fatal(err)
	}
}
