package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func AesEncrypt(key []byte, plainText string) (encoded string, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	plainTextBytes := pkcs7Padding([]byte(plainText), blockSize)
	//Make the cipher text a byte array of size BlockSize + the length of the message
	cipherText := make([]byte, aes.BlockSize+len(plainTextBytes))
	//iv is the cipherText up to the blocksize (16)
	iv := cipherText[:aes.BlockSize]
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(plainTextBytes))
	blockMode.CryptBlocks(crypted, plainTextBytes)

	return base64.RawStdEncoding.EncodeToString(crypted), nil
}

func AesDecrypt(key []byte, encrypted string) (decoded string, err error) {
	//Remove base64 encoding:
	cipherText, err := base64.RawStdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	//Create a new AES cipher with the key and encrypted message
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	//iv is the cipherText up to the blocksize (16)
	iv := cipherText[:aes.BlockSize]
	blockMode := cipher.NewCBCDecrypter(block, iv[:blockSize])
	origData := make([]byte, len(encrypted))
	blockMode.CryptBlocks(origData, cipherText)
	origData = pkcs7UnPadding(origData)

	return string(origData), nil
}

func pkcs7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(cipherText, padtext...)
}

func pkcs7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])

	return origData[:(length - unpadding)]
}
