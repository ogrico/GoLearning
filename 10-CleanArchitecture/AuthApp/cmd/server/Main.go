package main

import (
	"database/sql"
	"log"

	"AuthApp/internal/infrastructure/config" //"github.com/gin-gonic/gin"
	"AuthApp/internal/infrastructure/http"
	"AuthApp/internal/infrastructure/persistence/mysql"
	"AuthApp/internal/usecase"
)

func main() {
	// Para despliegue en producción habilitar gin.ReleaseMode
	/*
		gin.SetMode(gin.ReleaseMode)
	*/
	config := config.GetConfig()

	db, err := sql.Open("mysql", config.DNS)
	if err != nil {
		log.Fatalf("Error en la conexión con la DB: %v", err)
	}

	userRepo := &mysql.UserRepository{DB: db}
	authUseCase := &usecase.AuthUseCase{
		UserRepo: userRepo,
	}

	authHandler := &http.AuthHandler{AuthUseCase: authUseCase}
	router := http.SetupRouter(authHandler)

	if err := router.Run(":9090"); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
