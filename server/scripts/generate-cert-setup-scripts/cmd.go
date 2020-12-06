package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/cbroglie/mustache"
	"github.com/stephenzsy/doc-locker/server/common/configurations"
)

func genRootCa(
	configDir string,
	templatesDir string,
	caCertificateConfig configurations.YubikeyStoredCertificateConfiguration) {

	templateFilename := path.Join(templatesDir, "generate-root-ca.sh.mustache")

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
	templateFilename := path.Join(templatesDir, "generate-int-ca.sh.mustache")

	rendered, e := mustache.RenderFile(templateFilename, map[string](interface{}){
		"slot":           string(caCertificateConfig.Yubikey.Slot),
		"cnfPath":        path.Join(configDir, "setup", fmt.Sprintf("ca-%s.cnf", key)),
		"csrPath":        path.Join(configDir, "setup", fmt.Sprintf("ca-%s.csr", key)),
		"privateKeyPath": path.Join(configDir, "setup", fmt.Sprintf("ca-%s-private-key.pem", key)),
		"subjectCn":      caCertificateConfig.Subject.CN,
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

func genKeyPair(
	key string,
	configDir string,
	templatesDir string,
) {
	templateFilename := path.Join(templatesDir, "generate-key-pair.sh.mustache")

	rendered, e := mustache.RenderFile(templateFilename, map[string](interface{}){
		"privateKeyPath": path.Join(configDir, "keypairs", fmt.Sprintf("key-%s.pem", key)),
	})
	if e != nil {
		log.Fatal(e)
	}
	e = ioutil.WriteFile(path.Join(configDir, "scripts", fmt.Sprintf("generate-%s-key-pair.sh", key)), []byte(rendered), 0755)
	if e != nil {
		log.Fatal(e)
	}
}

type Algorithm string

var (
	rsa   Algorithm = "rsa"
	ecdsa Algorithm = "ecdsa"
)

func genEndCert(
	algorithm Algorithm,
	certType configurations.SecretType,
	key configurations.SecretName,
	configDir string,
	templatesDir string,
	certificateConfig configurations.CertificateConfig,
	certSetupConfig configurations.ServerSetupCertificatesConfiguration,
	caKey string,
	caYubikeySlot configurations.YubikeySlotId,
) {
	templateFilename := path.Join(templatesDir, fmt.Sprintf("generate-%s-cert.sh.mustache", certType))

	getSanList := func(l []string) []map[string](interface{}) {
		result := make([]map[string](interface{}), 0, len(l))
		for i, v := range l {
			result = append(result, map[string](interface{}){
				"index": fmt.Sprintf("%d", i+1),
				"value": v,
			})
		}
		return result
	}

	rendered, e := mustache.RenderFile(templateFilename, map[string](interface{}){
		"useRsa":               algorithm == rsa,
		"useEcdsa":             algorithm == ecdsa,
		"csrCnfPath":           path.Join(configDir, "setup", fmt.Sprintf("%s-cert-%s-csr.cnf", certType, key)),
		"crtCnfPath":           path.Join(configDir, "setup", fmt.Sprintf("%s-cert-%s-crt.cnf", certType, key)),
		"csrPath":              path.Join(configDir, "setup", fmt.Sprintf("%s-cert-%s.csr", certType, key)),
		"privateKeyPath":       path.Join(configDir, "setup", fmt.Sprintf("%s-key-%s.pem", certType, key)),
		"subjectCn":            certificateConfig.Subject.CN,
		"caPath":               path.Join(configDir, "setup", fmt.Sprintf("ca-%s.pem", caKey)),
		"certPath":             path.Join(configDir, "setup", fmt.Sprintf("%s-%s.pem", certType, key)),
		"pkcs11slotId":         configurations.GetPkcs11SlotIdMapping(caYubikeySlot),
		"rootCaPath":           path.Join(configDir, "setup", "ca-root.pem"),
		"bundleCertPath":       configurations.Configurations().SecretsConfiguration().GetCertPath(certType, key),
		"configPrivateKeyPath": configurations.Configurations().SecretsConfiguration().GetPrivateKeyPath(certType, key),
		"sans": map[string](interface{}){
			"ips": getSanList(certificateConfig.SANs.IPs),
		},
		"libPaths": map[string]string{
			"pkcs11": certSetupConfig.LibPaths.Pkcs11,
			"ykcs11": certSetupConfig.LibPaths.Ykcs11,
		},
	})
	if e != nil {
		log.Fatal(e)
	}
	e = ioutil.WriteFile(path.Join(configDir, "scripts", fmt.Sprintf("generate-%s-cert-%s.sh", certType, key)), []byte(rendered), 0755)
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
	certificatesConfig := serverSetupConfig.Certificates
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
	genIntermediateCa(
		"service",
		configDir,
		serverConfigTemplatePath,
		certificatesConfig,
		certificatesConfig.Ca.Service[0])

	// deploy
	genKeyPair(
		"deploy",
		configDir,
		serverConfigTemplatePath,
	)
	genEndCert(
		rsa,
		configurations.SecretTypeClient,
		configurations.SecretNameDeployAzureServicePrincipal,
		configDir,
		serverConfigTemplatePath,
		certificatesConfig.Areas.Deploy.AzureServicePrincipal[0],
		certificatesConfig,
		"deploy",
		certificatesConfig.Ca.Deploy[0].Yubikey.Slot,
	)
	genEndCert(
		rsa,
		configurations.SecretTypeServer,
		configurations.SecretNameDeploySds,
		configDir,
		serverConfigTemplatePath,
		certificatesConfig.Areas.Deploy.SdsServer[0],
		certificatesConfig,
		"deploy",
		certificatesConfig.Ca.Deploy[0].Yubikey.Slot,
	)
	genEndCert(
		rsa,
		configurations.SecretTypeClient,
		configurations.SecretNameDeploySds,
		configDir,
		serverConfigTemplatePath,
		certificatesConfig.Areas.Deploy.SdsClient[0],
		certificatesConfig,
		"deploy",
		certificatesConfig.Ca.Deploy[0].Yubikey.Slot,
	)

	genEndCert(
		rsa,
		configurations.SecretTypeServer,
		configurations.SecretNameProxy,
		configDir,
		serverConfigTemplatePath,
		certificatesConfig.Areas.Proxy.Server[0],
		certificatesConfig,
		"service",
		certificatesConfig.Ca.Service[0].Yubikey.Slot,
	)
}
