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

func genEnvoy(
	ctx app_context.AppContext,
	templatesDir string,
	serverSetupConfig configurations.ServerSetupConfiguration) (e error) {
	templateFilename := path.Join(templatesDir, "envoy.yaml.mustache")
	configDir, e := configurations.GetConfigurationsDir(ctx)
	if e != nil {
		return
	}
	rendered, e := mustache.RenderFile(templateFilename, map[string]interface{}{
		"sdsServer": map[string]interface{}{
			"address":   serverSetupConfig.SdsListener.Address,
			"portValue": serverSetupConfig.SdsListener.Port,
		},
		"sdsClient": map[string]interface{}{
			"certPath": path.Join(configDir, "certs", "client-cert-deploy-sds.pem"),
			"keyPath":  path.Join(configDir, "certsk", "client-key-deploy-sds.pem"),
		},
		"server": map[string]interface{}{
			"address":   serverSetupConfig.ServerListener.Address,
			"portValue": serverSetupConfig.ServerListener.Port,
		},
		"proxy": map[string]interface{}{
			"address":     serverSetupConfig.ProxyListener.Address,
			"portValue":   serverSetupConfig.ProxyListener.Port,
			"sdsCertName": fmt.Sprintf("%s-%s", configurations.SecretTypeServer, configurations.SecretNameProxy),
		},
	})
	if e != nil {
		return
	}
	e = ioutil.WriteFile(path.Join(configDir, "envoy", "envoy-config.yaml"), []byte(rendered), 0644)
	if e != nil {
		return
	}
	return
}

func main() {
	serverConfigTemplatePath := os.Getenv("DOCLOCKER_SETUP_TEMPLATES_DIR")
	if serverConfigTemplatePath == "" {
		log.Fatal("environment name is null: DOCLOCKER_SETUP_TEMPLATES_DIR")
	}
	serverConfigTemplatePath = path.Join(serverConfigTemplatePath, "server-config")
	serviceContext, e := app_context.NewAppServiceContext(context.Background(), app_context.WellKnownCallerdBootstrap)
	if e != nil {
		log.Fatal(e)
	}
	serviceContext = serviceContext.Elevate()
	serverSetupConfig, e := configurations.GetServerSetupConfiguration(serviceContext)
	if e != nil {
		log.Fatal(e)
	}
	if e = genEnvoy(serviceContext, serverConfigTemplatePath, serverSetupConfig); e != nil {
		log.Fatal(e)
	}
}
