package auth_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/ahyalfan/go-toolbox-icems/auth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var token string

func TestVerifyAuthToken_CustomClaims(t *testing.T) {
	email := "rahasia@gmail.com"
	privateKey := "./../../private_key.pem"
	token, err := auth.GenerateToken(privateKey, func() *auth.UserClaimsSpesifikRole {
		return &auth.UserClaimsSpesifikRole{
			ID:       uuid.NewString(),
			Email:    email,
			RoleName: "tes",
			RegisteredClaims: jwt.RegisteredClaims{
				Subject:   email,
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(10) * time.Minute)),
			},
		}
	})
	fmt.Println(token)
	fmt.Println(err)
	assert.Nil(t, err)
	assert.NotEqual(t, "", token)

	publicKey := "./../../public_key.pem"

	v, err := auth.VerifyToken(
		publicKey,
		token,
		func() *auth.UserClaimsSpesifikRole {
			return &auth.UserClaimsSpesifikRole{}
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, v)
}
