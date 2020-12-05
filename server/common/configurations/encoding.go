package configurations

import (
	"encoding/hex"
	"encoding/json"
	"strings"
)

type HexString []byte

func (s *HexString) MarshalJSON() ([]byte, error) {
	return json.Marshal(strings.ToUpper(hex.EncodeToString(*s)))
}

func (s *HexString) UnmarshalJson(data []byte) (err error) {
	var encodedStr string
	if err = json.Unmarshal(data, &encodedStr); err != nil {
		return
	}
	*s, err = hex.DecodeString(encodedStr)
	return
}
