package main

import (
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"path"

	"github.com/stephenzsy/doc-locker/server/common/configurations"
)

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
