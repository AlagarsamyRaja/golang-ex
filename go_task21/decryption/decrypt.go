package decryption

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func Decrypt(text, key []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(text) < nonceSize {
		return nil, fmt.Errorf("ciphertext is short . Check the cipher text")
	}

	nonce, encryptedText := text[:nonceSize], text[nonceSize:]

	decryptedText,err:= gcm.Open(nil, nonce, encryptedText, nil)
	if err != nil {
		return nil,fmt.Errorf("%s\n",err)
	}

	return decryptedText,nil

}