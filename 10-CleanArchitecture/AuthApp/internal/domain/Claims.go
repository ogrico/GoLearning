package domain

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
