package security

import (
	"crypto/aes"
	"crypto/rand"
)

func GenerateAes256Key() ([]byte, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	return b, err
}

func GenerateAesIv() ([]byte, error) {
	b := make([]byte, aes.BlockSize)
	_, err := rand.Read(b)
	return b, err
}

func GetAesCipherTextLength(src []byte) int {
	l := len(src)
	if l%aes.BlockSize == 0 {
		return l
	}
	return (l/aes.BlockSize + 1) * aes.BlockSize
}
