package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// Define your JWT secret key (keep it safe and load from config/env in production)
var jwtSecretKey = []byte("jwttokenforsnsms")

// Claims struct for JWT payload
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateJWTToken creates a signed JWT token with username and role
func GenerateJWTToken(username, role string) (string, error) {
	claims := Claims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 1 day expiry
			Issuer:    "snsms",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}
