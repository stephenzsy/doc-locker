package configurations

import (
	"encoding/json"
	"fmt"
)

type KeyType string
type RsaKeySize int
type EcCurve string

const (
	RSA KeyType = "RSA"
	EC  KeyType = "EC"
)

func (s *KeyType) UnmarshalJSON(data []byte) error {
	a := (*string)(s)
	err := json.Unmarshal(data, a)
	if err != nil {
		return err
	}

	// Validate the valid enum values
	switch *s {
	case RSA, EC:
		return nil
	default:
		return fmt.Errorf("invalid value for KeyType: %s", *a)
	}
}

const (
	RsaKeySize3072 RsaKeySize = 3072
	RsaKeySize4096 RsaKeySize = 4096
)

func (s *RsaKeySize) UnmarshalJSON(data []byte) error {
	a := (*int)(s)
	err := json.Unmarshal(data, a)
	if err != nil {
		return err
	}

	// Validate the valid enum values
	switch *s {
	case RsaKeySize3072, RsaKeySize4096:
		return nil
	default:
		return fmt.Errorf("invalid value for RsaKeySize: %d", *a)
	}
}

const (
	EcCurveP256 EcCurve = "P-256"
	EcCurveP384 EcCurve = "P-384"
)

func (s *EcCurve) UnmarshalJSON(data []byte) error {
	a := (*string)(s)
	err := json.Unmarshal(data, a)
	if err != nil {
		return err
	}

	// Validate the valid enum values
	switch *s {
	case EcCurveP256, EcCurveP384:
		return nil
	default:
		return fmt.Errorf("invalid value for EcCurve: %s", *a)
	}
}

type KeyProperties struct {
	KeyType    KeyType    `json:"kty"`
	RsaKeySize RsaKeySize `json:"keySize,omitempty"`
	EcCurve    EcCurve    `json:"crv,omitempty"`
}
