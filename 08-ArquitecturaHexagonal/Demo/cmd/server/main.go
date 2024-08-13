package main

import (
	"net/http"

	"github.com/ogrico/GoLearning/08-ArquitecturaHexagonal/Demo/internal/adapters/http"
	"github.com/ogrico/GoLearning/08-ArquitecturaHexagonal/Demo/internal/adapters/repository"
	"github.com/ogrico/GoLearning/08-ArquitecturaHexagonal/Demo/internal/application"
)

func main() {
	userRepo := repository.NewJSONUserRepository("../../data.json")
	userService := application.NewUserService(userRepo)
	http.HandleFunc("/users", CreateUserHandler(userService))

	http.ListenAndServe(":8080", nil)
}
