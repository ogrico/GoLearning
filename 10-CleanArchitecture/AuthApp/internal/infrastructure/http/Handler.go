package http

import (
	"AuthApp/internal/infrastructure/jwt"
	"AuthApp/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthUseCase *usecase.AuthUseCase
}

// GetToken autentica y devuelve el token
func (h *AuthHandler) GetToken(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.AuthUseCase.Authenticate(loginData.Username, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// ValidateToken valida el token "Jwt"
func (h *AuthHandler) ValidateToken(c *gin.Context) {
	var request struct {
		Jwt string `json:"jwt"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims, err := jwt.ValidateJwt(request.Jwt)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"valid": true, "claims": claims})
}
