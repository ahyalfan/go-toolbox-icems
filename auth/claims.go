package auth

import "github.com/golang-jwt/jwt/v5"

type UserClaimsSpesifikRole struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	RoleID   string `json:"role_id"`
	RoleName string `json:"role_name"`
	AppID    string `json:"app_id"`
	jwt.RegisteredClaims
}
