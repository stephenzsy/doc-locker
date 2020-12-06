package main

import (
	"context"
	"flag"
	"log"

	"github.com/stephenzsy/doc-locker/server/common/configurations"
	sds_provisioner_azure "github.com/stephenzsy/doc-locker/server/sds/provisioners/azure"
)

var (
	secretTypeStr = flag.String("secret-type", "", "Type of the secret")
	secretNameStr = flag.String("secret-name", "", "Name of the secret")
	importCert    = flag.Bool("import-cert", false, "Import certificate into the cloud")
)

func main() {
	flag.Parse()
	azureProvisioner, err := sds_provisioner_azure.NewAzureCertificatesProvisioner()
	if err != nil {
		log.Fatal(err)
	}
	secretType, err := configurations.SecretTypeFromString(*secretTypeStr)
	if err != nil {
		log.Fatal(err)
	}
	secretName, err := configurations.SecretNameFromString(*secretNameStr)
	if err != nil {
		log.Fatal(err)
	}
	result, err := azureProvisioner.ImportCertificate(context.Background(), secretType, secretName)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(result)
}
