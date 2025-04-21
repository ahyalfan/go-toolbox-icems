package utils

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// GeneratorRandomString generates a random string of the specified size.
// The string will be composed of uppercase and lowercase letters (A-Z, a-z)
// and digits (0-9).
//
// Example usage:
//
//	randomString := GeneratorRandomString(10) // returns something like "a1B2c3D4E5"
func GeneratorRandomString(size int) string {
	random := rand.New(rand.NewSource(time.Now().Unix()))
	charset := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	latters := make([]rune, size)
	for i := range latters {
		latters[i] = charset[random.Intn(len(charset))]
	}
	return string(latters)
}

// GeneratorRandomNumberString generates a random string of digits (0-9) of the specified size.
//
// Example usage:
//
//	randomNumber := GeneratorRandomNumberString(6) // returns something like "123456"
func GeneratorRandomNumberString(size int) string {
	random := rand.New(rand.NewSource(time.Now().Unix()))
	charset := []rune("0123456789")
	latters := make([]rune, size)
	for i := range latters {
		latters[i] = charset[random.Intn(len(charset))]
	}
	return string(latters)
}

// GeneratorRandomNumber generates a random integer between 0 and 1,000,000.
//
// Example usage:
//
//	randomNum := GeneratorRandomNumber() // returns something like 987654
func GeneratorRandomNumber() int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return random.Intn(1000000)
}

// GenerateUUID generates a new random UUID using the uuid package.
// If an error occurs while generating the UUID, it returns a new UUID using uuid.New().
//
// Example usage:
//
//	newUUID := GenerateUUID() // returns something like "b1c44c23-1bc9-4d52-bef1-ffed803d5689"
func GenerateUUID() uuid.UUID {
	uuidData, err := uuid.NewRandom()
	if err != nil {
		return uuid.New()
	}
	return uuidData
}
