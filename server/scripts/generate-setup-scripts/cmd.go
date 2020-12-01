package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"

	"github.com/cbroglie/mustache"
	"github.com/stephenzsy/doc-locker/server/common/configurations"
)

func genEnvoy(configDir string, templatesDir string, serverSetupConfig *configurations.ServerSetupConfiguration) {
	ldsYamlTemplateFilename := path.Join(templatesDir, "lds_yaml.mustache")
	proxyListnerConfig := serverSetupConfig.ProxyListener()
	rendered, e := mustache.RenderFile(ldsYamlTemplateFilename, map[string]string{
		"listenerAddress": proxyListnerConfig.Address,
		"listenerPort":    strconv.Itoa(int(proxyListnerConfig.Port)),
	})
	if e != nil {
		log.Fatal(e)
	}
	e = ioutil.WriteFile(path.Join(configDir, "envoy", "lds.yaml"), []byte(rendered), 0644)
	if e != nil {
		log.Fatal(e)
	}
}

func genRootCa(
	configDir string,
	templatesDir string,
	caCertificateConfig configurations.YubikeyStoredCertificateConfiguration) {

	templateFilename := path.Join(templatesDir, "generate-root-ca_sh.mustache")

	rendered, e := mustache.RenderFile(templateFilename, map[string]string{
		"slot":           string(caCertificateConfig.Yubikey.Slot),
		"cnfPath":        path.Join(configDir, "setup", "ca-root.cnf"),
		"privateKeyPath": path.Join(configDir, "setup", "ca-root-private-key.pem"),
		"subjectCn":      caCertificateConfig.Subject.CN,
		"serial":         caCertificateConfig.Serial,
		"certPath":       path.Join(configDir, "setup", "ca-root.pem"),
	})
	if e != nil {
		log.Fatal(e)
	}
	e = ioutil.WriteFile(path.Join(configDir, "scripts", "generate-root-ca.sh"), []byte(rendered), 0755)
	if e != nil {
		log.Fatal(e)
	}
}

func genIntermediateCa(
	key string,
	configDir string,
	templatesDir string,
	certSetupConfig configurations.ServerSetupCertificatesConfiguration,
	caCertificateConfig configurations.YubikeyStoredCertificateConfiguration) {
	templateFilename := path.Join(templatesDir, "generate-int-ca_sh.mustache")

	rendered, e := mustache.RenderFile(templateFilename, map[string](interface{}){
		"slot":           string(caCertificateConfig.Yubikey.Slot),
		"cnfPath":        path.Join(configDir, "setup", fmt.Sprintf("ca-%s.cnf", key)),
		"csrPath":        path.Join(configDir, "setup", fmt.Sprintf("ca-%s.csr", key)),
		"privateKeyPath": path.Join(configDir, "setup", fmt.Sprintf("ca-%s-private-key.pem", key)),
		"subjectCn":      caCertificateConfig.Subject.CN,
		"serial":         caCertificateConfig.Serial,
		"caPath":         path.Join(configDir, "setup", "ca-root.pem"),
		"certPath":       path.Join(configDir, "setup", fmt.Sprintf("ca-%s.pem", key)),
		"pkcs11slotId":   configurations.GetPkcs11SlotIdMapping(certSetupConfig.Ca.Root[0].Yubikey.Slot),
		"libPaths": map[string]string{
			"pkcs11": certSetupConfig.LibPaths.Pkcs11,
			"ykcs11": certSetupConfig.LibPaths.Ykcs11,
		},
	})
	if e != nil {
		log.Fatal(e)
	}
	e = ioutil.WriteFile(path.Join(configDir, "scripts", fmt.Sprintf("generate-%s-ca.sh", key)), []byte(rendered), 0755)
	if e != nil {
		log.Fatal(e)
	}
}

func genKeyCert(
	key string,
	configDir string,
	templatesDir string,
	certificateConfig configurations.CertificateConfig,
	certSetupConfig configurations.ServerSetupCertificatesConfiguration,
	caKey string,
	caYubikeySlot configurations.YubikeySlotId,
) {
	templateFilename := path.Join(templatesDir, "generate-key-cert_sh.mustache")

	rendered, e := mustache.RenderFile(templateFilename, map[string](interface{}){
		"csrCnfPath":     path.Join(configDir, "setup", fmt.Sprintf("key-cert-%s-csr.cnf", key)),
		"crtCnfPath":     path.Join(configDir, "setup", fmt.Sprintf("key-cert-%s-crt.cnf", key)),
		"csrPath":        path.Join(configDir, "setup", fmt.Sprintf("key-cert-%s.csr", key)),
		"privateKeyPath": path.Join(configDir, "setup", fmt.Sprintf("key-cert-%s-private-key.pem", key)),
		"subjectCn":      certificateConfig.Subject.CN,
		"serial":         certificateConfig.Serial,
		"caPath":         path.Join(configDir, "setup", fmt.Sprintf("ca-%s.pem", caKey)),
		"certPath":       path.Join(configDir, "setup", fmt.Sprintf("key-cert-%s.pem", key)),
		"pkcs11slotId":   configurations.GetPkcs11SlotIdMapping(caYubikeySlot),
		"rootCaPath":     path.Join(configDir, "setup", "ca-root.pem"),
		"bundleCertPath": path.Join(configDir, "certsk", fmt.Sprintf("key-cert-%s.pem", key)),
		"libPaths": map[string]string{
			"pkcs11": certSetupConfig.LibPaths.Pkcs11,
			"ykcs11": certSetupConfig.LibPaths.Ykcs11,
		},
	})
	if e != nil {
		log.Fatal(e)
	}
	e = ioutil.WriteFile(path.Join(configDir, "scripts", fmt.Sprintf("generate-%s-key-cert.sh", key)), []byte(rendered), 0755)
	if e != nil {
		log.Fatal(e)
	}
}

func genAzureServicePrincipal(
	key string,
	configDir string,
	templatesDir string,
	certificateConfig configurations.CertificateConfig,
	certSetupConfig configurations.ServerSetupCertificatesConfiguration,
	caKey string,
	caYubikeySlot configurations.YubikeySlotId,
) {
	templateFilename := path.Join(templatesDir, "generate-azure-service-principal_sh.mustache")

	rendered, e := mustache.RenderFile(templateFilename, map[string](interface{}){
		"csrCnfPath":     path.Join(configDir, "setup", fmt.Sprintf("azure-service-principal-%s-csr.cnf", key)),
		"crtCnfPath":     path.Join(configDir, "setup", fmt.Sprintf("azure-service-principal-%s-crt.cnf", key)),
		"csrPath":        path.Join(configDir, "setup", fmt.Sprintf("azure-service-principal-%s.csr", key)),
		"privateKeyPath": path.Join(configDir, "setup", fmt.Sprintf("azure-service-principal-%s-private-key.pem", key)),
		"subjectCn":      certificateConfig.Subject.CN,
		"serial":         certificateConfig.Serial,
		"caPath":         path.Join(configDir, "setup", fmt.Sprintf("ca-%s.pem", caKey)),
		"certPath":       path.Join(configDir, "setup", fmt.Sprintf("azure-service-principal-%s.pem", key)),
		"pkcs11slotId":   configurations.GetPkcs11SlotIdMapping(caYubikeySlot),
		"rootCaPath":     path.Join(configDir, "setup", "ca-root.pem"),
		"bundleCertPath": path.Join(configDir, "certs", fmt.Sprintf("azure-service-principal-%s.pem", key)),
		"libPaths": map[string]string{
			"pkcs11": certSetupConfig.LibPaths.Pkcs11,
			"ykcs11": certSetupConfig.LibPaths.Ykcs11,
		},
	})
	if e != nil {
		log.Fatal(e)
	}
	e = ioutil.WriteFile(path.Join(configDir, "scripts", fmt.Sprintf("generate-%s-azure-service-principal.sh", key)), []byte(rendered), 0755)
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	serverConfigTemplatePath := os.Getenv("DOCLOCKER_SETUP_TEMPLATES_DIR")
	if serverConfigTemplatePath == "" {
		log.Fatal("environment name is null: DOCLOCKER_SETUP_TEMPLATES_DIR")
	}
	serverConfigTemplatePath = path.Join(serverConfigTemplatePath, "server-config")
	serverSetupConfig, e := configurations.Configurations().ServerSetup()
	if e != nil {
		log.Fatal(e)
	}
	configDir := configurations.Configurations().ConfigRootDir()
	genEnvoy(
		configDir,
		serverConfigTemplatePath,
		serverSetupConfig)

	certificatesConfig := serverSetupConfig.Certificates()
	genRootCa(
		configDir,
		serverConfigTemplatePath,
		certificatesConfig.Ca.Root[0])
	genIntermediateCa(
		"deploy",
		configDir,
		serverConfigTemplatePath,
		certificatesConfig,
		certificatesConfig.Ca.Deploy[0])

	genKeyCert(
		"deploy-cipher",
		configDir,
		serverConfigTemplatePath,
		certificatesConfig.Keys.Deploy[0],
		certificatesConfig,
		"deploy",
		certificatesConfig.Ca.Deploy[0].Yubikey.Slot,
	)

	genAzureServicePrincipal(
		"deploy",
		configDir,
		serverConfigTemplatePath,
		certificatesConfig.Client.Deploy.AzureServicePrincipal[0],
		certificatesConfig,
		"deploy",
		certificatesConfig.Ca.Deploy[0].Yubikey.Slot,
	)
}
