package configurations

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"

	"github.com/stephenzsy/doc-locker/server/common/app_context"
)

func loadConfigFromFile(filePath string, configData interface{}) error {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, configData)
	return err
}

func GetConfigurationsRootDir(ctx app_context.AppContext) (dir string, err error) {
	if err = app_context.VerifyElevated(ctx); err != nil {
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
	rootDir, err := GetConfigurationsRootDir(ctx)
	dir = path.Join(rootDir, ctx.Deployment().Id())
	return
}
