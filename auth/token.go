package auth

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidSigningMethod = errors.New("invalid token signing method")
	ErrParseToken           = errors.New("error parsing token")
	ErrInvalidTokenClaims   = errors.New("invalid token claims")
	ErrLoadPublicKey        = errors.New("error loading public key")
	ErrLoadPrivateKey       = errors.New("error loading private key")
	ErrSignToken            = errors.New("error signing token")
)

// VerifyToken verifies and parses a JWT token using the provided public key path
// and returns the parsed claims of type T.
func VerifyToken[T jwt.Claims](pathPublicKey, tokenStr string, claimsFactory func() T) (T, error) {
	var zero T

	claims := claimsFactory()

	publicKey, err := LoadPublicKey(pathPublicKey)
	if err != nil {
		return zero, fmt.Errorf("error loading public key: %w", err)
		// return zero, ErrLoadPublicKey
	}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, ErrInvalidSigningMethod
		}
		return publicKey, nil
	})
	if err != nil {
		return zero, fmt.Errorf("error parsing token: %w", err)
		// return zero, ErrParseToken
	}

	parsedClaims, ok := token.Claims.(T)
	if !ok || !token.Valid {
		return zero, ErrInvalidTokenClaims
	}

	return parsedClaims, nil
}

// GenerateToken creates a signed JWT token using RS256 and the provided claims.
func GenerateToken[T jwt.Claims](pathPirvateKey string, claimsFactory func() T) (string, error) {
	claims := claimsFactory()
	// Membaca private key RSA
	privateKey, err := LoadPrivateKey(pathPirvateKey) // Fungsi untuk load private key RSA
	if err != nil {
		return "", fmt.Errorf("error loading private key: %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return tokenString, err
}
