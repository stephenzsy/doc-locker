package configurations

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
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
	configDir   string
	serverSetup runOnceUtil
	deployment  runOnceUtil
}

func (c *configurations) ServerSetup() (*ServerSetupConfiguration, error) {
	c.serverSetup.once.Do(func() {
		c.serverSetup.data, c.serverSetup.err = newSetupConfiguration(c.configDir)
	})

	return c.serverSetup.data.(*ServerSetupConfiguration), c.serverSetup.err
}

func (c *configurations) Deployment() (*DeploymentConfigurationFile, error) {
	c.deployment.once.Do(func() {
		c.deployment.data, c.deployment.err = newDeploymentConfiguration(c.configDir)
	})

	return c.deployment.data.(*DeploymentConfigurationFile), c.deployment.err
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
