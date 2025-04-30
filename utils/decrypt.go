package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/ahyalfan/go-toolbox-icems/auth"
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

func DecryptAny[T any](input []byte, keyString string) (T, error) {
	key := []byte(keyString)
	block, err := aes.NewCipher(key)
	if err != nil {
		var zero T
		return zero, err
	}

	if len(input) < aes.BlockSize {
		var zero T
		return zero, fmt.Errorf("ciphertext too short")
	}

	iv := input[:aes.BlockSize]
	ciphertext := input[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)

	var result T
	err = json.Unmarshal(plaintext, &result)
	if err != nil {
		var zero T
		return zero, err
	}

	return result, nil
}

// DecryptTextStringRsa decrypts a base64-encoded RSA-encrypted string using a private key.
// The function reads the private key from the specified PEM file and uses it to decrypt
// the data that was encrypted with the matching public key. It assumes the ciphertext uses
// PKCS#1 v1.5 padding.
//
// Args:
//   - inputEncode (string): The base64-encoded encrypted string.
//   - pathPrivateKey (string): The path to the RSA private key PEM file.
//
// Returns:
//   - string: The original plaintext string.
//   - error: An error if base64 decoding, key loading, or decryption fails.
//
// Example usage:
//
//	plaintext, err := DecryptTextStringRsa(encryptedText, "keys/private.pem")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println("Decrypted text:", plaintext)
func DecryptTextStringRsa(inputEncode string, pathPrivateKey string) (string, error) {
	input, err := base64.StdEncoding.DecodeString(inputEncode)
	if err != nil {
		return "", err
	}

	privateKey, err := auth.LoadPrivateKey(pathPrivateKey)
	if err != nil {
		return "", err
	}

	decryptedText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, input)
	if err != nil {
		return "", err
	}

	return string(decryptedText), nil
}
