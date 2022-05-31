package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func Aes256Encode(plaintext string, key string, iv string, blockSize int) string {
	bKey := []byte(key)
	bIV := []byte(iv)
	bPlaintext := PKCS5Padding([]byte(plaintext), blockSize, len(plaintext))
	block, _ := aes.NewCipher(bKey)
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func Aes256Decode(encrypted string, key string, iv string, blockSize int) string {
	bKey := []byte(key)
	bIV := []byte(iv)
	block, _ := aes.NewCipher(bKey)
	mode := cipher.NewCBCDecrypter(block, bIV)
	encryptedDecode, _ := base64.StdEncoding.DecodeString(encrypted)
	if len(encryptedDecode) == 0 || len(encryptedDecode)%aes.BlockSize != 0 {
		return "error"
	}
	result := make([]byte, len(encryptedDecode))
	mode.CryptBlocks(result, encryptedDecode)
	return string(PKCS5Unpadding([]byte(result)))
}

func PKCS5Unpadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
