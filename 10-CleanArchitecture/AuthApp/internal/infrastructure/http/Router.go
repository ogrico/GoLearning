package http

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(authHandler *AuthHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/login/getToken", authHandler.GetToken)
	r.POST("/login/validateToken", authHandler.ValidateToken)

	return r
}
