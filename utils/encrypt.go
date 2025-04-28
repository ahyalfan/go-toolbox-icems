package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"io"

	"github.com/ahyalfan/go-toolbox-icems/auth"
)

// EncryptText encrypts the given plaintext using AES encryption with CFB mode.
// It expects the `key` to be a valid AES key of appropriate size (16, 24, or 32 bytes).
// The function generates a random Initialization Vector (IV) for each encryption, which
// is prepended to the ciphertext. The result is returned as a concatenation of the IV
// and the encrypted data.
//
// Args:
//   - input (byte slice): The plaintext data to encrypt.
//   - key (byte slice): The AES encryption key (should be 16, 24, or 32 bytes long).
//
// Returns:
//   - ciphertext (byte slice): The result of the AES encryption, which includes the IV
//     as the first `aes.BlockSize` bytes followed by the encrypted data.
//   - error: Returns an error if any occurs during the encryption process, such as
//     an invalid key or issues generating the IV.
//
// Example usage:
//
//	ciphertext, err := EncryptText([]byte("my secret data"), []byte("myAESkey12345678"))
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println("Encrypted data:", ciphertext)
func EncryptText(input []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(input))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], input)

	return ciphertext, nil
}

// EncryptTextString encrypts the given plaintext string using AES encryption with CFB mode.
// It expects the `keyString` to be a valid AES key of appropriate size (16, 24, or 32 bytes).
// The function generates a random Initialization Vector (IV) for each encryption, which
// is prepended to the ciphertext. The result is returned as a concatenation of the IV
// and the encrypted data, which is a byte slice.
//
// Args:
//   - inputString (string): The plaintext string to encrypt.
//   - keyString (string): The AES encryption key as a string (should be 16, 24, or 32 bytes long).
//
// Returns:
//   - ciphertext (byte slice): The result of the AES encryption, which includes the IV
//     as the first `aes.BlockSize` bytes followed by the encrypted data.
//   - error: Returns an error if any occurs during the encryption process, such as
//     an invalid key or issues generating the IV.
//
// Example usage:
//
//	ciphertext, err := EncryptTextString("my secret data", "myAESkey12345678")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println("Encrypted data:", ciphertext)
func EncryptTextString(inputString string, keyString string) ([]byte, error) {
	key := []byte(keyString)
	input := []byte(inputString)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(input))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], input)

	return ciphertext, nil
}

// EncryptTextStringRsa encrypts a plaintext string using RSA public key encryption.
// The function reads a public key from the specified PEM file path and uses it to encrypt
// the input string with PKCS#1 v1.5 padding. The resulting ciphertext is encoded in base64
// for safe storage or transmission.
//
// Args:
//   - inputString (string): The plaintext string to be encrypted.
//   - path (string): The path to the RSA public key PEM file.
//
// Returns:
//   - string: The base64-encoded ciphertext string.
//   - error: An error if encryption or key loading fails.
//
// Example usage:
//
//	encryptedText, err := EncryptTextStringRsa("my secret message", "keys/public.pem")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println("Encrypted (base64):", encryptedText)
func EncryptTextStringRsa(inputString string, path string) (string, error) {
	publicKey, err := auth.LoadPublicKey(path)
	if err != nil {
		return "", err
	}
	input := []byte(inputString)
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, input)
	if err != nil {
		return "", err
	}

	encodeString := base64.StdEncoding.EncodeToString(cipherText)
	return encodeString, nil
}
