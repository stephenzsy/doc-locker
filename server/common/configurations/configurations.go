package configurations

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

func loadConfigFromFile(filePath string, configData interface{}) error {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bytes, configData)
	return err
}

type configurations struct {
	configDir        string
	serverSetup      *ServerSetupConfiguration
	serverSetupError error
	serverSetupOnce  sync.Once
}

func (c *configurations) ServerSetup() (*ServerSetupConfiguration, error) {
	c.serverSetupOnce.Do(func() {
		c.serverSetup, c.serverSetupError = newSetupConfiguration(c.configDir)
	})

	return c.serverSetup, c.serverSetupError
}

var (
	config     *configurations
	configOnce sync.Once
)

func newConfigurations() *configurations {
	configDir := os.Getenv("DOCLOCKER_CONFIG_DIR")
	return &configurations{
		configDir: configDir,
	}
}

func Configurations() *configurations {
	configOnce.Do(func() {
		config = newConfigurations()
	})

	return config
}

func (c *configurations) ConfigRootDir() string {
	return c.configDir
}
