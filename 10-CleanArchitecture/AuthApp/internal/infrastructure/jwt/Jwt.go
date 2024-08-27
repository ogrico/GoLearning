package jwt

import (
	"errors"
	"fmt"
	"time"

	"AuthApp/internal/domain"
	"AuthApp/internal/infrastructure/config"

	"github.com/golang-jwt/jwt/v5"
)

// ValidateJwt valida un token JWT
func ValidateJwt(tokenString string) (*domain.Claims, error) {
	cfg := config.GetConfig()
	claims := &domain.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		// Validar el método de firma
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de firma inesperado: %v", token.Header["alg"])
		}
		return []byte(cfg.JWTKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("token inválido: %v", err)
	}

	// Verificar si el token es válido
	if !token.Valid {
		return nil, errors.New("token no válido")
	}

	// Verificar si el token ha expirado
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("token expirado")
	}

	return claims, nil
}
