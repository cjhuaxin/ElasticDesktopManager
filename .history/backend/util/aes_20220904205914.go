package util

import(
	"crypto/aes"
)

func AesEncrypt(plainText, key, iv []byte)([]byte,error){
	block, err := aes.NewCipher(key)
	if err != nil {
        return nil, err
    }

	blockSize := block.BlockSize()
	plaintext = PKCS7Padding(plaintext, blockSize)
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
    padding := blockSize - len(ciphertext)%blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, padtext...)
}
