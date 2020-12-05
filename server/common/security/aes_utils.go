package security

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

func GenerateAes256Key() ([]byte, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	return b, err
}

func generateAesIv(blockSize int) ([]byte, error) {
	b := make([]byte, blockSize)
	_, err := rand.Read(b)
	return b, err
}

func pkcs5Padding(ciphertext *[]byte, blockSize int) []byte {
	padding := blockSize - len(*ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(*ciphertext, padtext...)
}

func pkcs5Trim(content *[]byte, blockSize int) []byte {
	l := len(*content)
	padding := int((*content)[l-1])
	return (*content)[:l-padding]
}

func AESEncrypt(encryptionKey []byte, content *[]byte) (encrypted []byte, err error) {
	cipherBlock, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return
	}

	iv, err := generateAesIv(cipherBlock.BlockSize())
	if err != nil {
		return
	}

	encrypter := cipher.NewCBCEncrypter(cipherBlock, iv)
	padded := pkcs5Padding(content, cipherBlock.BlockSize())
	cipherText := make([]byte, len(padded))
	encrypter.CryptBlocks(cipherText, padded)
	encrypted = append(iv, cipherText...)
	return
}

func AESDecrypt(encryptionKey []byte, encrypted *[]byte) (content []byte, err error) {
	cipherBlock, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return
	}

	iv, cipherText := (*encrypted)[:cipherBlock.BlockSize()], (*encrypted)[cipherBlock.BlockSize():]

	decrypter := cipher.NewCBCDecrypter(cipherBlock, iv)
	padded := make([]byte, len(cipherText))
	decrypter.CryptBlocks(padded, cipherText)
	content = pkcs5Trim(&padded, cipherBlock.BlockSize())
	return
}
