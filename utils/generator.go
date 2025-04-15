package utils

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func GeneratorRandomString(size int) string {
	random := rand.New(rand.NewSource(time.Now().Unix()))
	charset := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	latters := make([]rune, size)
	for i := range latters {
		latters[i] = charset[random.Intn(len(charset))]
	}
	return string(latters)
}

func GeneratorRandomNumberString(size int) string {
	random := rand.New(rand.NewSource(time.Now().Unix()))
	charset := []rune("0123456789")
	latters := make([]rune, size)
	for i := range latters {
		latters[i] = charset[random.Intn(len(charset))]
	}
	return string(latters)
}

func GeneratorRandomNumber() int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return random.Intn(1000000)
}

func GenerateUUID() uuid.UUID {
	uuidData, err := uuid.NewRandom()
	if err != nil {
		return uuid.New()
	}
	return uuidData
}
