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
	Slot84 YubikeySlotId = "84"
)

type CertificateConfig struct {
	Subject struct {
		CN string `json:"CN"`
	} `json:"subject"`
	Serial string `json:"serial"`
	SANs   struct {
		IPs []string `json:"ips"`
	} `json:"sans"`
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
	case Slot84:
		return "7"
	}
	return ""
}

type ServerSetupCertificatesConfiguration struct {
	LibPaths struct {
		Pkcs11 string `json:"pkcs11"`
		Ykcs11 string `json:"ykcs11"`
	} `json:"libPaths"`
	Ca struct {
		Root    []YubikeyStoredCertificateConfiguration `json:"root"`
		Deploy  []YubikeyStoredCertificateConfiguration `json:"deploy"`
		Service []YubikeyStoredCertificateConfiguration `json:"service"`
	} `json:"ca"`
	Keys struct {
		Deploy []CertificateConfig `json:"deploy"`
	} `json:"keys"`
	Areas struct {
		Deploy struct {
			AzureServicePrincipal []CertificateConfig `json:"azureServicePrincipal"`
			SdsServer             []CertificateConfig `json:"sdsServer"`
			SdsClient             []CertificateConfig `json:"sdsClient"`
		} `json:"deploy"`
		Proxy struct {
			Server []CertificateConfig `json:"server"`
		} `json:"proxy"`
	} `json:"areas"`
}

type ServerSetupCloudAzureConfiguration struct {
	AadOauthEndpoint    string `json:"aadOauthEndpoint"`
	AadResourceKeyVault string `json:"aadResourceKeyVault"`
	AadTenantId         string `json:"aadTenantId"`
	ApplicationId       string `json:"applicationId"`
	KeyVaultBaseUrl     string `json:"keyVaultBaseUrl"`
}

type ServerSetupCloudConfiguration struct {
	Azure ServerSetupCloudAzureConfiguration `json:"azure"`
}

type ServerSetupConfiguration struct {
	ProxyListener  ListenerConfig                       `json:"proxyListener"`
	ServerListener ListenerConfig                       `json:"serverListener"`
	SdsListener    ListenerConfig                       `json:"sdsListener"`
	Certificates   ServerSetupCertificatesConfiguration `json:"certificates"`
	Cloud          ServerSetupCloudConfiguration        `json:"cloud"`
}

func newSetupConfiguration(configDir string) (*ServerSetupConfiguration, error) {
	config := ServerSetupConfiguration{}
	err := loadConfigFromFile(path.Join(configDir, "setup", "server.json"), &config)
	return &config, err
}
