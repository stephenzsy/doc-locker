package configurations

import "path"

type ListenerConfig struct {
	Address string `json:"address"`
	Port    uint   `json:"port"`
}

type ServerSetupConfiguration struct {
	data struct {
		ServerListener ListenerConfig `json:"serverListener"`
		ProxyListener  ListenerConfig `json:"proxyListener"`
	}
}

func newSetupConfiguration(configDir string) (*ServerSetupConfiguration, error) {
	config := ServerSetupConfiguration{}
	err := loadConfigFromFile(path.Join(configDir, "setup", "server.json"), &config.data)
	return &config, err
}

func (c *ServerSetupConfiguration) ServerListener() ListenerConfig {
	return c.data.ServerListener
}

func (c *ServerSetupConfiguration) ProxyListener() ListenerConfig {
	return c.data.ProxyListener
}
