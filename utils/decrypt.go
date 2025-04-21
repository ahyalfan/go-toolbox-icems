package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

// DecryptText decrypts the given ciphertext using AES decryption in CFB mode.
// It expects the key to be the same size as the block size (e.g., 16 bytes for AES-128).
// The input data should include the IV (initialization vector) as the first 16 bytes
// followed by the actual ciphertext.
//
// It returns the decrypted plaintext and an error if any occurred during decryption.
//
// Example usage:
//
//	decrypted, err := DecryptText(ciphertext, key)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(string(decrypted))
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

// DecryptTextString decrypts the given ciphertext and returns the plaintext as a string.
// The key is provided as a string, and it's converted to a byte slice internally.
// The input data should include the IV (initialization vector) as the first 16 bytes
// followed by the actual ciphertext.
//
// It returns the decrypted plaintext as a string and an error if any occurred during decryption.
//
// Example usage:
//
//	decryptedText, err := DecryptTextString(ciphertext, "mysecretkey")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(decryptedText)
func DecryptTextString(input []byte, keyString string) (string, error) {
	key := []byte(keyString)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(input) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := input[:aes.BlockSize]
	ciphertext := input[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)

	return string(plaintext), nil
}
