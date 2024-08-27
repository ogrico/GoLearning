package usecase

import (
	"AuthApp/internal/domain"
	"AuthApp/internal/infrastructure/config"
	"AuthApp/internal/repository"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthUseCase struct {
	UserRepo repository.UserRepository
}

var cfg = config.GetConfig()

// GenerateToken genera un token JWT usando los datos del usuario.
func GenerateToken(user domain.User) (string, error) {

	claims := &domain.Claims{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmar el token con la clave secreta
	signedToken, err := token.SignedString([]byte(cfg.JWTKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Authenticate valida credenciales y devuelve el token
func (uc *AuthUseCase) Authenticate(username, password string) (string, error) {
	user, err := uc.UserRepo.FindByUsername(username)
	if err != nil {
		return "", err
	}

	if user.Password != password {
		return "", errors.New("datos incorrectos")
	}

	return GenerateToken(user)

}

func ValidateJwt(tokenString string) (*domain.Claims, error) {
	claims := &domain.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		// Validar el método de firma
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de firma inesperado: %v", token.Header["alg"])
		}
		return cfg.JWTKey, nil
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
