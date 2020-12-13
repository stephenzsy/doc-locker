package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/stephenzsy/doc-locker/server/common/app_context"
	"github.com/stephenzsy/doc-locker/server/common/configurations"
	sds_provisioner_azure "github.com/stephenzsy/doc-locker/server/sds/provisioners/azure"
)

var (
	importCert = flag.Bool("import-cert", false, "Import certificate into the cloud")
)

func main() {
	importCertFs := flag.NewFlagSet("import-cert", flag.ExitOnError)
	secretTypeStr := importCertFs.String("secret-type", "", "Type of the secret")
	secretNameStr := importCertFs.String("secret-name", "", "Name of the secret")

	if len(os.Args) < 2 {
		fmt.Println("expected 'import-cert' commands")
		os.Exit(1)
	}

	serviceContext, err := app_context.NewAppServiceContext(context.Background(), app_context.WellKnownCallerdBootstrap)
	if err != nil {
		log.Fatal(err)
	}
	serviceContext = serviceContext.Elevate()
	azureProvisioner, err := sds_provisioner_azure.NewAzureCertificatesProvisioner(serviceContext)
	if err != nil {
		log.Fatal(err)
	}

	switch os.Args[1] {
	case "import-cert":
		importCertFs.Parse(os.Args[2:])
		var secretType configurations.SecretType
		if err = secretType.FromString(*secretTypeStr); err != nil {
			log.Fatal(err)
		}
		var secretName configurations.SecretName
		if err = secretName.FromString(*secretNameStr); err != nil {
			log.Fatal(err)
		}
		result, err := azureProvisioner.ImportCertificate(serviceContext, secretType, secretName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Imported certificate with thumbprint: %s\n", *result.X509Thumbprint)
	default:
		log.Fatal("expected 'import-cert' commands")
		os.Exit(1)
	}
}
