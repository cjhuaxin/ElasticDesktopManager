package util

func AesEncrypt(plainText, key, iv []byte)([]byte,error){
	block, err := aes.NewCipher(key)
}
