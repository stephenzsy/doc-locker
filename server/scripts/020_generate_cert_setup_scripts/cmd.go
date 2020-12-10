package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/cbroglie/mustache"
	"github.com/stephenzsy/doc-locker/server/common/app_context"
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

type scriptGenUtil struct {
	serverSetupConfig configurations.ServerSetupConfiguration
	secretsConfig     configurations.SecretsConfiguration
}

func (this scriptGenUtil) genRootCa(
	seq string,
	caCertificateConfig configurations.YubikeyStoredCertificateConfiguration) {

	templateFilename := path.Join(serverConfigTemplatesPathBase, "generate-root-ca.sh.mustache")

	rendered, e := mustache.RenderFile(templateFilename, map[string](interface{}){
		"keyProps":       getKeyPropsContext(caCertificateConfig.KeyProperties),
		"slot":           string(caCertificateConfig.Yubikey.Slot),
		"cnfPath":        path.Join(this.serverSetupConfig.TmpPath(), fmt.Sprintf("ca-%s.cnf", configurations.CaRoleRoot)),
		"privateKeyPath": path.Join(this.serverSetupConfig.TmpPath(), fmt.Sprintf("ca-%s-private-key.pem", configurations.CaRoleRoot)),
		"subjectCn":      caCertificateConfig.Subject.CN,
		"serial":         caCertificateConfig.Serial,
		"certPath":       this.secretsConfig.GetCaPath(configurations.CaRoleRoot),
	})
	if e != nil {
		log.Fatal(e)
	}
	e = ioutil.WriteFile(path.Join(this.serverSetupConfig.ScriptsPath(), fmt.Sprintf("%s-generate-%s-ca.sh", seq, configurations.CaRoleRoot)), []byte(rendered), 0700)
	if e != nil {
		log.Fatal(e)
	}
}

func (this scriptGenUtil) genIntermediateCa(
	seq string,
	caRole configurations.CaRole,
	caCertificateConfig configurations.YubikeyStoredCertificateConfiguration) {

	templateFilename := path.Join(serverConfigTemplatesPathBase, "generate-int-ca.sh.mustache")

	rendered, e := mustache.RenderFile(templateFilename, map[string](interface{}){
		"keyProps":       getKeyPropsContext(caCertificateConfig.KeyProperties),
		"slot":           string(caCertificateConfig.Yubikey.Slot),
		"cnfPath":        path.Join(this.serverSetupConfig.TmpPath(), fmt.Sprintf("ca-%s.cnf", caRole)),
		"csrPath":        path.Join(this.serverSetupConfig.TmpPath(), fmt.Sprintf("ca-%s.csr", caRole)),
		"privateKeyPath": path.Join(this.serverSetupConfig.TmpPath(), fmt.Sprintf("ca-%s-private-key.pem", caRole)),
		"subjectCn":      caCertificateConfig.Subject.CN,
		"caPath":         this.secretsConfig.GetCaPath(caCertificateConfig.Issuer),
		"certPath":       path.Join(this.serverSetupConfig.TmpPath(), fmt.Sprintf("%s-%s.pem", configurations.SecretTypeCa, caRole)),
		"bundleCertPath": this.secretsConfig.GetCaPath(caRole),
		"pkcs11slotId":   configurations.GetPkcs11SlotIdMapping(this.serverSetupConfig.Certificates.Ca[configurations.CaRoleRoot][0].Yubikey.Slot),
		"libPaths": map[string]string{
			"pkcs11": this.serverSetupConfig.Certificates.LibPaths.Pkcs11,
			"ykcs11": this.serverSetupConfig.Certificates.LibPaths.Ykcs11,
		},
	})
	if e != nil {
		log.Fatal(e)
	}
	e = ioutil.WriteFile(path.Join(this.serverSetupConfig.ScriptsPath(), fmt.Sprintf("%s-generate-%s-ca.sh", seq, caRole)), []byte(rendered), 0700)
	if e != nil {
		log.Fatal(e)
	}
}

func (this scriptGenUtil) genEndCert(
	seq string,
	secretType configurations.SecretType,
	secretName configurations.SecretName,
	certConfig configurations.CertificateConfiguration) {

	templateFilename := path.Join(serverConfigTemplatesPathBase, "generate-end-cert.sh.mustache")
	issuerConfig := this.serverSetupConfig.Certificates.Ca[certConfig.Issuer][0]

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
		"keyProps":       getKeyPropsContext(certConfig.KeyProperties),
		"useTlsServer":   secretType == configurations.SecretTypeServer,
		"useTlsClient":   secretType == configurations.SecretTypeClient,
		"useKeyPair":     secretType == configurations.SecretTypeKeyPair,
		"csrCnfPath":     path.Join(this.serverSetupConfig.TmpPath(), fmt.Sprintf("%s-cert-%s-csr.cnf", secretType, secretName)),
		"crtCnfPath":     path.Join(this.serverSetupConfig.TmpPath(), fmt.Sprintf("%s-cert-%s-crt.cnf", secretType, secretName)),
		"csrPath":        path.Join(this.serverSetupConfig.TmpPath(), fmt.Sprintf("%s-cert-%s.csr", secretType, secretName)),
		"privateKeyPath": this.secretsConfig.GetPrivateKeyPath(secretType, secretName),
		"subjectCn":      certConfig.Subject.CN,
		"caPath":         this.secretsConfig.GetCaPath(certConfig.Issuer),
		"certPath":       path.Join(this.serverSetupConfig.TmpPath(), fmt.Sprintf("%s-%s.pem", secretType, secretName)),
		"pkcs11slotId":   configurations.GetPkcs11SlotIdMapping(issuerConfig.Yubikey.Slot),
		"bundleCertPath": this.secretsConfig.GetCertPath(secretType, secretName),
		"hasSANs":        len(certConfig.SANs.IPs) > 0,
		"sans": map[string](interface{}){
			"ips": getSanList(certConfig.SANs.IPs),
		},
		"libPaths": map[string]string{
			"pkcs11": this.serverSetupConfig.Certificates.LibPaths.Pkcs11,
			"ykcs11": this.serverSetupConfig.Certificates.LibPaths.Ykcs11,
		},
	})
	if e != nil {
		log.Fatal(e)
	}
	e = ioutil.WriteFile(path.Join(this.serverSetupConfig.ScriptsPath(), fmt.Sprintf("%s-generate-%s-cert-%s.sh", seq, secretType, secretName)), []byte(rendered), 0755)
	if e != nil {
		log.Fatal(e)
	}
}

var serverConfigTemplatesPathBase string

func main() {
	serverConfigTemplatePath := os.Getenv("DOCLOCKER_SETUP_TEMPLATES_DIR")
	if serverConfigTemplatePath == "" {
		log.Fatal("environment name is null: DOCLOCKER_SETUP_TEMPLATES_DIR")
	}
	serverConfigTemplatesPathBase = path.Join(serverConfigTemplatePath, "server-config")
	serviceContext, e := app_context.NewAppServiceContext(context.Background(), app_context.WellKnownCallerdBootstrap)
	if e != nil {
		log.Fatal(e)
	}
	serviceContext = serviceContext.Elevate()
	serverSetupConfig, e := configurations.GetServerSetupConfiguration(serviceContext)
	if e != nil {
		log.Fatal(e)
	}
	secretsConfig, e := configurations.GetSecretsConfiguration(serviceContext)
	if e != nil {
		log.Fatal(e)
	}
	certificatesConfig := serverSetupConfig.Certificates
	sgu := scriptGenUtil{
		serverSetupConfig: serverSetupConfig,
		secretsConfig:     secretsConfig,
	}
	sgu.genRootCa(
		"0010",
		certificatesConfig.Ca[configurations.CaRoleRoot][0])
	sgu.genIntermediateCa(
		"0020",
		configurations.CaRoleDeploy,
		certificatesConfig.Ca[configurations.CaRoleDeploy][0])
	sgu.genIntermediateCa(
		"0030",
		configurations.CaRoleService,
		certificatesConfig.Ca[configurations.CaRoleService][0])

	// deploy
	sgu.genEndCert(
		"0110",
		configurations.SecretTypeKeyPair,
		configurations.SecretNameDeploy,
		certificatesConfig.Areas.Deploy.KeyPair[0])
	sgu.genEndCert(
		"0111",
		configurations.SecretTypeClient,
		configurations.SecretNameDeploySdsAzureServicePrincipal,
		certificatesConfig.Areas.Deploy.SdsAzureServicePrincipal[0])
	sgu.genEndCert(
		"0120",
		configurations.SecretTypeServer,
		configurations.SecretNameDeploySds,
		certificatesConfig.Areas.Deploy.SdsServer[0])
	sgu.genEndCert(
		"0121",
		configurations.SecretTypeClient,
		configurations.SecretNameDeploySdsEnvoy,
		certificatesConfig.Areas.Deploy.SdsClientEnvoy[0])

	// proxy
	sgu.genEndCert(
		"0210",
		configurations.SecretTypeServer,
		configurations.SecretNameProxy,
		certificatesConfig.Areas.Proxy.Server[0])
	sgu.genEndCert(
		"0211",
		configurations.SecretTypeClient,
		configurations.SecretNameProxy,
		certificatesConfig.Areas.Proxy.Client[0])

	// api
	sgu.genEndCert(
		"0310",
		configurations.SecretTypeServer,
		configurations.SecretNameApi,
		certificatesConfig.Areas.Backend.Api[0])

	// site
	sgu.genEndCert(
		"0410",
		configurations.SecretTypeServer,
		configurations.SecretNameSite,
		certificatesConfig.Areas.Backend.Site[0])
}
