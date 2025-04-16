package utils_test

import (
	"fmt"
	"testing"

	"github.com/ahyalfan/go-toolbox-icems/utils"
	"github.com/stretchr/testify/assert"
)

var key = []byte("1234567890123456") // 16, 24, atau 32 byte

func TestEncryptDecryptSucces(t *testing.T) {
	input := []byte("nik dari keluaragra yogi")
	e, err := utils.EncryptText(input, key)
	assert.Nil(t, err)
	assert.NotNil(t, e)

	hasil, err := utils.DecryptText(e, key)
	assert.Nil(t, err)
	assert.Equal(t, string(input), string(hasil))
}

func TestEncryptDecryptErrorKey(t *testing.T) {
	input := []byte("nik dari keluaragra yogi")
	keyNew := []byte("12345678901234567")
	e, err := utils.EncryptText(input, keyNew)
	fmt.Println(err)
	assert.Nil(t, e)
	assert.NotNil(t, err)

	hasil, err := utils.DecryptText(e, keyNew)
	assert.Nil(t, hasil)
	assert.NotNil(t, err)
}

func TestEncryptDecryptTextStringSucces(t *testing.T) {
	input := "nik dari keluaragra yogi"
	e, err := utils.EncryptTextString(input, string(key))
	fmt.Println(err)
	assert.Nil(t, err)
	assert.NotNil(t, e)

	hasil, err := utils.DecryptTextString(e, string(key))
	fmt.Println(hasil)
	assert.Nil(t, err)
	assert.Equal(t, input, hasil)
}
