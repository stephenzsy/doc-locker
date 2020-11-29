package configurations

import "path"

type listener struct {
	Address string `json:"address"`
	Port    uint   `json:"port"`
}

type serverSetupConfiguration struct {
	data struct {
		ServerListener listener `json:"serverListener"`
	}
}

func newSetupConfiguration(configDir string) *serverSetupConfiguration {
	config := serverSetupConfiguration{}
	loadConfigFromFile(path.Join(configDir, "setup", "server.json"), &config.data)
	return &config
}

func (c *serverSetupConfiguration) ServerListener() *listener {
	return &c.data.ServerListener
}
