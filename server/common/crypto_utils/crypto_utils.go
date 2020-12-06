package crypto_utils

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

func ParseECPrivateKeyFromPemFile(filename string) (privateKey *ecdsa.PrivateKey, err error) {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	pemBlock, _ := pem.Decode(fileBytes)
	return x509.ParseECPrivateKey(pemBlock.Bytes)
}

func MarshalPKCS8PrivateKeyPemBlock(key interface{}) (encoded []byte, err error) {
	der, err := x509.MarshalPKCS8PrivateKey(key)
	encoded = pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: der,
	})
	return
}
