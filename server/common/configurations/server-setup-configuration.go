package configurations

import "path"

type ListenerConfig struct {
	Address string `json:"address"`
	Port    uint   `json:"port"`
}

type YubikeySlotId string

const (
	Slot82 YubikeySlotId = "82"
	Slot83 YubikeySlotId = "83"
)

type CertificateConfig struct {
	Subject struct {
		CN string `json:"CN"`
	} `json:"subject"`
	Serial string `json:"serial"`
}

type YubikeyStoredCertificateConfiguration struct {
	CertificateConfig
	Yubikey struct {
		Slot YubikeySlotId `json:"slot"`
	}
}

func GetPkcs11SlotIdMapping(slot YubikeySlotId) string {
	switch slot {
	case Slot82:
		return "5"
	case Slot83:
		return "6"
	}
	return ""
}

type ServerSetupCertificatesConfiguration struct {
	LibPaths struct {
		Pkcs11 string `json:"pkcs11"`
		Ykcs11 string `json:"ykcs11"`
	} `json:"libPaths"`
	Ca struct {
		Root   []YubikeyStoredCertificateConfiguration `json:"root"`
		Deploy []YubikeyStoredCertificateConfiguration `json:"deploy"`
	} `json:"ca"`
	Keys struct {
		Deploy []CertificateConfig `json:"deploy"`
	} `json:"keys"`
	Client struct {
		Deploy struct {
			AzureServicePrincipal []CertificateConfig `json:"azureServicePrincipal"`
		} `json:"deploy"`
	} `json:"client"`
}

type ServerSetupConfiguration struct {
	data struct {
		ServerListener ListenerConfig                       `json:"serverListener"`
		ProxyListener  ListenerConfig                       `json:"proxyListener"`
		Certificates   ServerSetupCertificatesConfiguration `json:"certificates"`
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

func (c *ServerSetupConfiguration) Certificates() ServerSetupCertificatesConfiguration {
	return c.data.Certificates
}
