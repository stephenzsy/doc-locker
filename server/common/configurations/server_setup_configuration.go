package configurations

import (
	"encoding/json"
	"fmt"
	"path"

	"github.com/stephenzsy/doc-locker/server/common/app_context"
	"github.com/stephenzsy/doc-locker/server/common/auth"
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

type CertificateConfiguration struct {
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
	CertificateConfiguration
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
			KeyPair                  []CertificateConfiguration `json:"keyPair"`
			SdsAzureServicePrincipal []CertificateConfiguration `json:"sdsAzureServicePrincipal"`
			SdsServer                []CertificateConfiguration `json:"sdsServer"`
			SdsClientEnvoy           []CertificateConfiguration `json:"sdsClientEnvoy"`
		} `json:"deploy"`
		Proxy struct {
			Server []CertificateConfiguration `json:"server"`
			Client []CertificateConfiguration `json:"client"`
		} `json:"proxy"`
		Backend struct {
			Api  []CertificateConfiguration `json:"api"`
			Site []CertificateConfiguration `json:"site"`
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
	configDir      string
	ProxyListener  ListenerConfig                       `json:"proxyListener"`
	ServerListener ListenerConfig                       `json:"serverListener"`
	SdsListener    ListenerConfig                       `json:"sdsListener"`
	Certificates   ServerSetupCertificatesConfiguration `json:"certificates"`
	Cloud          ServerSetupCloudConfiguration        `json:"cloud"`
}

func (c ServerSetupConfiguration) TmpPath() string {
	return path.Join(c.configDir, "tmp")
}

func (c ServerSetupConfiguration) ScriptsPath() string {
	return path.Join(c.configDir, "scripts")
}

func newServerSetupConfiguration(ctx app_context.AppContext) (config ServerSetupConfiguration, err error) {

	configDir, err := GetConfigurationsDir(ctx)
	if err != nil {
		return
	}
	configRootDir, err := GetConfigurationsRootDir(ctx)
	if err != nil {
		return
	}

	config = ServerSetupConfiguration{
		configDir: configDir,
	}
	configPath := path.Join(configRootDir, "setup")
	err = loadConfigFromFile(path.Join(configPath, "server.json"), &config)
	return
}

func GetServerSetupConfiguration(ctx app_context.AppContext) (config ServerSetupConfiguration, err error) {

	if err = app_context.VerifyElevated(ctx); err != nil {
		return
	}
	if err = app_context.VerifyCallerId(ctx, auth.SystemCallerIdBootstrap); err != nil {
		return
	}

	return newServerSetupConfiguration(ctx)
}
