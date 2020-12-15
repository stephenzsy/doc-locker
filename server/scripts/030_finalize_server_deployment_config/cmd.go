package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"

	"github.com/stephenzsy/doc-locker/server/common/app_context"
	"github.com/stephenzsy/doc-locker/server/common/auth"
	"github.com/stephenzsy/doc-locker/server/common/configurations"
	"github.com/stephenzsy/doc-locker/server/common/crypto_utils"
)

func UpdateEncryptionMaterial(ctx app_context.AppContext, config *configurations.DeploymentConfigurationFile) (privateKey *rsa.PrivateKey, err error) {
	secretsConfiguration, err := configurations.GetSecretsConfiguration(ctx)
	if err != nil {
		return
	}
	privateKey, err = secretsConfiguration.GetRsaPrivateKey(configurations.SecretTypeKeyPair, configurations.SecretNameDeploy)
	if err != nil {
		return
	}
	// TODO: verify key
	encryptionKey, err := crypto_utils.GenerateAes256Key()
	config.EncryptionMaterial.EncryptedKey, err = rsa.EncryptOAEP(sha512.New384(), rand.Reader, &privateKey.PublicKey, encryptionKey, []byte{})
	return privateKey, err
}

func updateAzureConfigurations(
	ctx app_context.AppContext,
	c *configurations.DeploymentConfigurationFile,
	privateKey *rsa.PrivateKey) (err error) {

	serverSetupConfig, err := configurations.GetServerSetupConfiguration(ctx)
	if err != nil {
		return
	}
	c.Cloud.Azure.ServerSetupCloudAzureConfiguration = serverSetupConfig.Cloud.Azure

	// service principal
	secretsConfiguration, err := configurations.GetSecretsConfiguration(ctx)
	if err != nil {
		return
	}
	certificate, err := secretsConfiguration.GetCertificate(
		configurations.SecretTypeClient,
		configurations.SecretNameDeploySdsAzureServicePrincipal)
	if err != nil {
		return
	}
	c.Cloud.Azure.ServicePrincipalCertificate = certificate.Raw
	thumbprint := sha1.Sum(certificate.Raw)
	c.Cloud.Azure.ServicePrincipalThumbprint = thumbprint[:]

	// private configuration
	privateConfig, encryptionKey, err := c.GetPrivateConfig(privateKey)
	if err != nil {
		return err
	}

	servicePrincipalPrivateKeyContent, err := ioutil.ReadFile(
		secretsConfiguration.GetPrivateKeyPath(
			configurations.SecretTypeClient,
			configurations.SecretNameDeploySdsAzureServicePrincipal))
	if err != nil {
		return err
	}
	pemBlock, _ := pem.Decode(servicePrincipalPrivateKeyContent)
	_, err = x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
	if err != nil {
		return err
	}
	privateConfig.Cloud.Azure.ServicePrincipalPrivateKey = pemBlock.Bytes
	c.SetPrivateConfig(&privateKey.PublicKey, encryptionKey, privateConfig)

	return
}

func main() {
	serviceContext, err := app_context.NewAppServiceContext(context.Background(), auth.SystemCallerIdBootstrap)
	serviceContext = serviceContext.Elevate()
	if err != nil {
		log.Fatal(err)
	}
	deploymentConfig, err := configurations.GetServerDeploymentConfigurationFile(serviceContext)
	if err != nil {
		log.Fatal(err)
	}

	if deploymentConfig.Version == 0 {
		// no configuration file
		privateKey, err := UpdateEncryptionMaterial(serviceContext, &deploymentConfig)
		if err != nil {
			log.Fatal(err)
		}
		err = updateAzureConfigurations(serviceContext, &deploymentConfig, privateKey)
		if err != nil {
			log.Fatal(err)
		}
		deploymentConfig.Version = 1
	}
	err = deploymentConfig.Save(serviceContext)
	if err != nil {
		log.Fatal(err)
	}
}
