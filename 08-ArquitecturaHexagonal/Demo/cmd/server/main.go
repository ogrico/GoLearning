package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"Demo/internal/adapters/api"
	"Demo/internal/adapters/repository"
	"Demo/internal/application"

	"github.com/gorilla/mux"
)

func main() {

	// Obtener el directorio actual
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Construir la ruta completa al archivo data.json
	jsonFilePath := filepath.Join(currentDir, "../../data.json")

	repo := repository.NewJSONRepository(jsonFilePath)
	userService := application.NewUserService(repo)
	handler := api.NewHandler(userService)

	r := mux.NewRouter()
	r.HandleFunc("/users", handler.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handler.GetUserByID).Methods("GET")
	fmt.Println("Server is running on 9090")
	log.Fatal(http.ListenAndServe(":9090", r))
}
