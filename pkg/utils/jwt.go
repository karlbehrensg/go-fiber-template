package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTClaims struct {
	*jwt.StandardClaims

	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
}

// CreateToken create token
func (c *JWTClaims) CreateToken() (string, error) {
	c.StandardClaims = &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// ValidateToken validate token
func (c *JWTClaims) ValidateToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, c, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return err
	}

	return nil
}
