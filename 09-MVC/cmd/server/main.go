package main

import (
	"MVC/internal/controller"
	"MVC/internal/database"
	"MVC/internal/repository"
	"MVC/internal/service"
	"log"
	"net/http"
)

func main() {
	db := database.Connect()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	http.HandleFunc("/users", userController.GetAllUsers)

	log.Println("Servidor corriendo en el puerto 9094...")
	if err := http.ListenAndServe(":9094", nil); err != nil {
		log.Fatal(err)
	}
}
