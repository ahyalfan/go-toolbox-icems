package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func DecryptText(input []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(input) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	iv := input[:aes.BlockSize]
	ciphertext := input[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)

	return plaintext, nil
}
