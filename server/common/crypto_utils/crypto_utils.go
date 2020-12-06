package crypto_utils

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rsa"
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

func ParseRsaPrivateKeyFromPemFile(filename string) (privateKey *rsa.PrivateKey, err error) {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	pemBlock, _ := pem.Decode(fileBytes)
	return x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
}

func MarshalPKCS8PrivateKeyPemBlock(key interface{}) (encoded []byte, err error) {
	der, err := x509.MarshalPKCS8PrivateKey(key)
	encoded = pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: der,
	})
	return
}

func MarshalCertificatesPemBlock(certificates ...*x509.Certificate) []byte {
	buffer := bytes.Buffer{}
	for _, c := range certificates {
		buffer.Write(pem.EncodeToMemory(&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: c.Raw,
		}))
	}
	return buffer.Bytes()
}

func MarshalRsaPrivateKeyPemBlock(key *ecdsa.PrivateKey) (encoded []byte, err error) {
	der, err := x509.MarshalECPrivateKey(key)
	encoded = pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: der,
	})
	return
}

func MarshalECPrivateKeyPemBlock(key *ecdsa.PrivateKey) (encoded []byte, err error) {
	der, err := x509.MarshalECPrivateKey(key)
	encoded = pem.EncodeToMemory(&pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: der,
	})
	return
}
