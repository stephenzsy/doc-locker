package configurations

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"sync"

	"github.com/stephenzsy/doc-locker/server/common/app_context"
	"github.com/stephenzsy/doc-locker/server/common/auth"
)

func loadConfigFromFile(filePath string, configData interface{}) error {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, configData)
	return err
}

type runOnceUtil struct {
	data interface{}
	err  error
	once sync.Once
}

type configurations struct {
	configDir     string
	deployment    runOnceUtil
	secertsConfig SecretsConfiguration
}

func (c *configurations) SecretsConfiguration() *SecretsConfiguration {
	return &c.secertsConfig
}

var (
	config     *configurations
	configOnce sync.Once
)

func GetConfigurationsRootDir(ctx app_context.AppContext) (dir string, err error) {
	if err = app_context.VerifyElevated(ctx); err != nil {
		return
	}
	if err = app_context.VerifyCallerId(ctx, auth.SystemCallerIdBootstrap, auth.ServiceCallerIdSds); err != nil {
		return
	}
	dir = os.Getenv("DOCLOCKER_CONFIG_DIR")
	if dir == "" {
		err = errors.New("invalid environment variable: DOCLOCKER_CONFIG_DIR")
	}
	return
}

func GetConfigurationsDir(ctx app_context.AppContext) (dir string, err error) {
	if err = app_context.VerifyElevated(ctx); err != nil {
		return
	}
	if err = app_context.VerifyCallerId(ctx, auth.SystemCallerIdBootstrap, auth.ServiceCallerIdSds); err != nil {
		return
	}
	rootDir, err := GetConfigurationsRootDir(ctx)
	dir = path.Join(rootDir, ctx.Deployment().Id())
	return
}
