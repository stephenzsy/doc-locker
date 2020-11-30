package security

import "crypto/rand"

func GenerateAes256Key() ([]byte, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	return b, err
}
