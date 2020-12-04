package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/cbroglie/mustache"
	"github.com/stephenzsy/doc-locker/server/common/configurations"
)

func genEnvoy(templatesDir string, serverSetupConfig *configurations.ServerSetupConfiguration) {
	templateFilename := path.Join(templatesDir, "envoy.yaml.mustache")
	configs := configurations.Configurations()
	configDir := configs.ConfigRootDir()
	rendered, e := mustache.RenderFile(templateFilename, map[string]interface{}{
		"sdsServer": map[string]interface{}{
			"address":   serverSetupConfig.SdsListener().Address,
			"portValue": serverSetupConfig.SdsListener().Port,
		},
		"sdsClient": map[string]interface{}{
			"certPath": path.Join(configDir, "certs", "client-cert-deploy-sds.pem"),
			"keyPath":  path.Join(configDir, "certsk", "client-key-deploy-sds.pem"),
		},
		"server": map[string]interface{}{
			"address":   serverSetupConfig.ServerListener().Address,
			"portValue": serverSetupConfig.ServerListener().Port,
		},
		"proxy": map[string]interface{}{
			"address":     serverSetupConfig.ProxyListener().Address,
			"portValue":   serverSetupConfig.ProxyListener().Port,
			"sdsCertName": configurations.SdsSecretNameProxyServer,
		},
	})
	if e != nil {
		log.Fatal(e)
	}
	e = ioutil.WriteFile(path.Join(configDir, "envoy", "envoy-config.yaml"), []byte(rendered), 0644)
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
	genEnvoy(
		serverConfigTemplatePath,
		serverSetupConfig)
}
