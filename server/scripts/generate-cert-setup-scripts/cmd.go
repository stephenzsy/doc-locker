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

func getOpensslEcCurveName(crv configurations.EcCurve) string {
	switch crv {
	case configurations.EcCurveP256:
		return "prime256v1"
	case configurations.EcCurveP384:
		return "secp384r1"
	}
	return ""
}

func getKeyPropsContext(keyProps configurations.KeyProperties) map[string](interface{}) {
	m := map[string](interface{}){
		"useRsa":      keyProps.KeyType == configurations.RSA,
		"useEc":       keyProps.KeyType == configurations.EC,
		"rsaKeySize":  keyProps.RsaKeySize,
		"ecCurveName": getOpensslEcCurveName(keyProps.EcCurve),
	}
	return m
}

func genRootCa(
	configDir string,
	templatesDir string,
	caCertificateConfig configurations.YubikeyStoredCertificateConfiguration) {

	templateFilename := path.Join(templatesDir, "generate-root-ca.sh.mustache")

	rendered, e := mustache.RenderFile(templateFilename, map[string](interface{}){
		"keyProps":       getKeyPropsContext(caCertificateConfig.KeyProperties),
		"slot":           string(caCertificateConfig.Yubikey.Slot),
		"cnfPath":        path.Join(configDir, "setup", fmt.Sprintf("ca-%s.cnf", configurations.CaRoleRoot)),
		"privateKeyPath": path.Join(configDir, "setup", fmt.Sprintf("ca-%s-private-key.pem", configurations.CaRoleRoot)),
		"subjectCn":      caCertificateConfig.Subject.CN,
		"serial":         caCertificateConfig.Serial,
		"certPath":       path.Join(configDir, "setup", fmt.Sprintf("ca-%s.pem", configurations.CaRoleRoot)),
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
	caRole configurations.CaRole,
	configDir string,
	templatesDir string,
	certSetupConfig configurations.ServerSetupCertificatesConfiguration,
	caCertificateConfig configurations.YubikeyStoredCertificateConfiguration) {
	templateFilename := path.Join(templatesDir, "generate-int-ca.sh.mustache")

	rendered, e := mustache.RenderFile(templateFilename, map[string](interface{}){
		"keyProps":       getKeyPropsContext(caCertificateConfig.KeyProperties),
		"slot":           string(caCertificateConfig.Yubikey.Slot),
		"cnfPath":        path.Join(configDir, "setup", fmt.Sprintf("ca-%s.cnf", caRole)),
		"csrPath":        path.Join(configDir, "setup", fmt.Sprintf("ca-%s.csr", caRole)),
		"privateKeyPath": path.Join(configDir, "setup", fmt.Sprintf("ca-%s-private-key.pem", caRole)),
		"subjectCn":      caCertificateConfig.Subject.CN,
		"caPath":         path.Join(configDir, "setup", fmt.Sprintf("ca-%s.pem", caCertificateConfig.Issuer)),
		"certPath":       path.Join(configDir, "setup", fmt.Sprintf("ca-%s.pem", caRole)),
		"pkcs11slotId":   configurations.GetPkcs11SlotIdMapping(certSetupConfig.Ca[configurations.CaRoleRoot][0].Yubikey.Slot),
		"libPaths": map[string]string{
			"pkcs11": certSetupConfig.LibPaths.Pkcs11,
			"ykcs11": certSetupConfig.LibPaths.Ykcs11,
		},
	})
	if e != nil {
		log.Fatal(e)
	}
	e = ioutil.WriteFile(path.Join(configDir, "scripts", fmt.Sprintf("generate-%s-ca.sh", caRole)), []byte(rendered), 0755)
	if e != nil {
		log.Fatal(e)
	}
}

func genKeyPair(
	key string,
	configDir string,
	templatesDir string,
	keyProps configurations.KeyProperties,
) {
	templateFilename := path.Join(templatesDir, "generate-key-pair.sh.mustache")

	rendered, e := mustache.RenderFile(templateFilename, map[string](interface{}){
		"keyProps":       getKeyPropsContext(keyProps),
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

func genEndCert(
	certType configurations.SecretType,
	key configurations.SecretName,
	configDir string,
	templatesDir string,
	certConfig configurations.CertificateConfig,
	certSetupConfig configurations.ServerSetupCertificatesConfiguration,
) {
	templateFilename := path.Join(templatesDir, "generate-end-cert.sh.mustache")
	issuerConfig := certSetupConfig.Ca[certConfig.Issuer][0]

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
		"keyProps":             getKeyPropsContext(certConfig.KeyProperties),
		"useTlsServer":         certType == configurations.SecretTypeServer,
		"useTlsClient":         certType == configurations.SecretTypeClient,
		"csrCnfPath":           path.Join(configDir, "setup", fmt.Sprintf("%s-cert-%s-csr.cnf", certType, key)),
		"crtCnfPath":           path.Join(configDir, "setup", fmt.Sprintf("%s-cert-%s-crt.cnf", certType, key)),
		"csrPath":              path.Join(configDir, "setup", fmt.Sprintf("%s-cert-%s.csr", certType, key)),
		"privateKeyPath":       path.Join(configDir, "setup", fmt.Sprintf("%s-key-%s.pem", certType, key)),
		"subjectCn":            certConfig.Subject.CN,
		"caPath":               path.Join(configDir, "setup", fmt.Sprintf("ca-%s.pem", certConfig.Issuer)),
		"certPath":             path.Join(configDir, "setup", fmt.Sprintf("%s-%s.pem", certType, key)),
		"pkcs11slotId":         configurations.GetPkcs11SlotIdMapping(issuerConfig.Yubikey.Slot),
		"rootCaPath":           path.Join(configDir, "setup", "ca-root.pem"),
		"bundleCertPath":       configurations.Configurations().SecretsConfiguration().GetCertPath(certType, key),
		"configPrivateKeyPath": configurations.Configurations().SecretsConfiguration().GetPrivateKeyPath(certType, key),
		"sans": map[string](interface{}){
			"ips": getSanList(certConfig.SANs.IPs),
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
		certificatesConfig.Ca[configurations.CaRoleRoot][0])
	genIntermediateCa(
		configurations.CaRoleDeploy,
		configDir,
		serverConfigTemplatePath,
		certificatesConfig,
		certificatesConfig.Ca[configurations.CaRoleDeploy][0])
	genIntermediateCa(
		configurations.CaRoleService,
		configDir,
		serverConfigTemplatePath,
		certificatesConfig,
		certificatesConfig.Ca[configurations.CaRoleService][0])

	// deploy
	genKeyPair(
		"deploy",
		configDir,
		serverConfigTemplatePath,
		certificatesConfig.Keys.Deploy[0])
	genEndCert(
		configurations.SecretTypeClient,
		configurations.SecretNameDeployAzureServicePrincipal,
		configDir,
		serverConfigTemplatePath,
		certificatesConfig.Areas.Deploy.AzureServicePrincipal[0],
		certificatesConfig)
	genEndCert(
		configurations.SecretTypeServer,
		configurations.SecretNameDeploySds,
		configDir,
		serverConfigTemplatePath,
		certificatesConfig.Areas.Deploy.SdsServer[0],
		certificatesConfig)
	genEndCert(
		configurations.SecretTypeClient,
		configurations.SecretNameDeploySds,
		configDir,
		serverConfigTemplatePath,
		certificatesConfig.Areas.Deploy.SdsClient[0],
		certificatesConfig)

	// proxy
	genEndCert(
		configurations.SecretTypeServer,
		configurations.SecretNameProxy,
		configDir,
		serverConfigTemplatePath,
		certificatesConfig.Areas.Proxy.Server[0],
		certificatesConfig)
	genEndCert(
		configurations.SecretTypeClient,
		configurations.SecretNameProxy,
		configDir,
		serverConfigTemplatePath,
		certificatesConfig.Areas.Proxy.Client[0],
		certificatesConfig)

	// api
	genEndCert(
		configurations.SecretTypeServer,
		configurations.SecretNameApi,
		configDir,
		serverConfigTemplatePath,
		certificatesConfig.Areas.Backend.Api[0],
		certificatesConfig)
	// site
	genEndCert(
		configurations.SecretTypeServer,
		configurations.SecretNameSite,
		configDir,
		serverConfigTemplatePath,
		certificatesConfig.Areas.Backend.Site[0],
		certificatesConfig)
}
