package crypto

import "encoding/base64"

// Encrypt aes + base64
func Encrypt(plaintext []byte, key string) (ciphertext2 string, err error) {
	iv := "0000000000000000"
	var ciphertext1 []byte
	ciphertext1, err = AESEncrypt(plaintext, []byte(iv), []byte(key))
	if err != nil {
		return
	}
	ciphertext2 = base64.StdEncoding.EncodeToString(ciphertext1)
	return
}

// Decrypt aes + base64
func Decrypt(ciphertext2, key string) (plaintext []byte, err error) {
	var ciphertext1 []byte
	ciphertext1, err = base64.StdEncoding.DecodeString(ciphertext2)
	if err != nil {
		return
	}
	iv := "0000000000000000"
	plaintext, err = AESDecrypt(ciphertext1, []byte(iv), []byte(key))
	return
}
