package main

import (
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

func genRootCa(configDir string,
	templatesDir string,
	certificatesConfig configurations.ServerSetupCertificatesConfiguration) {
	templateFilename := path.Join(templatesDir, "generate-root-ca_sh.mustache")
	caCertificateConfig := certificatesConfig.Ca.Root[0]

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
	genEnvoy(configurations.Configurations().ConfigRootDir(), serverConfigTemplatePath, serverSetupConfig)
	genRootCa(configurations.Configurations().ConfigRootDir(), serverConfigTemplatePath, serverSetupConfig.Certificates())
}
