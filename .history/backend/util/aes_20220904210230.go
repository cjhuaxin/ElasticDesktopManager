package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func AesEncrypt(plainText, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	plainText = pkcs7Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(plainText))
	blockMode.CryptBlocks(crypted, plainText)

	return crypted, nil
}

func AesDecrypt(ciphertext []byte, key, iv []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    blockSize := block.BlockSize()
    blockMode := cipher.NewCBCDecrypter(block, iv[:blockSize])
    origData := make([]byte, len(ciphertext))
    blockMode.CryptBlocks(origData, ciphertext)
    origData = pkcs7UnPadding(origData)
    return origData, nil
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
