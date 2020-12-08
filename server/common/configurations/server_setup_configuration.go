package configurations

import (
	"encoding/json"
	"fmt"
	"path"
)

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

func (s *YubikeySlotId) UnmarshalJSON(data []byte) error {
	a := (*string)(s)
	err := json.Unmarshal(data, a)
	if err != nil {
		return err
	}

	// Validate the valid enum values
	switch *s {
	case Slot82, Slot83, Slot84:
		return nil
	default:
		return fmt.Errorf("invalid value for YubikeySlotId: %s", *a)
	}
}

type CaRole string

const (
	CaRoleRoot    CaRole = "root"
	CaRoleDeploy  CaRole = "deploy"
	CaRoleService CaRole = "service"
)

func (s *CaRole) UnmarshalJSON(data []byte) error {
	a := (*string)(s)
	err := json.Unmarshal(data, a)
	if err != nil {
		return err
	}

	// Validate the valid enum values
	switch *s {
	case CaRoleRoot, CaRoleDeploy, CaRoleService:
		return nil
	default:
		return fmt.Errorf("invalid value for CaRole: %s", *a)
	}
}

type CertificateConfig struct {
	Subject struct {
		CN string `json:"CN"`
	} `json:"subject"`
	Serial string `json:"serial"`
	SANs   struct {
		IPs []string `json:"ips"`
	} `json:"sans"`
	KeyProperties KeyProperties `json:"keyProps"`
	Issuer        CaRole        `json:"issuer"`
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
	Ca    map[CaRole][]YubikeyStoredCertificateConfiguration `json:"ca"`
	Areas struct {
		Deploy struct {
			KeyPair   []CertificateConfig `json:"keyPair"`
			SdsServer []CertificateConfig `json:"sdsServer"`
			SdsClient []CertificateConfig `json:"sdsClient"`
		} `json:"deploy"`
		Proxy struct {
			Server []CertificateConfig `json:"server"`
			Client []CertificateConfig `json:"client"`
		} `json:"proxy"`
		Backend struct {
			Api  []CertificateConfig `json:"api"`
			Site []CertificateConfig `json:"site"`
		} `json:"backend"`
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
