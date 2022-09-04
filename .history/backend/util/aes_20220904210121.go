package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func AesEncrypt(plainText, key, iv []byte)([]byte,error){
	block, err := aes.NewCipher(key)
	if err != nil {
        return nil, err
    }

	blockSize := block.BlockSize()
	plainText = PKCS7Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(plainText))
}

func PKCS7Padding(cipherText []byte, blockSize int) []byte {
    padding := blockSize - len(cipherText)%blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(cipherText, padtext...)
}
