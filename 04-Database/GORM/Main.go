package main

import (
	"github.com/gorilla/mux"
	"gorm/handlers"
	"log"
	"net/http"
)

func main() {
	//Rutas Http
	router := mux.NewRouter()
	router.HandleFunc("/api/autor", handlers.GetAutors).Methods("GET")
	log.Fatal(http.ListenAndServe(":9090", router))
	println("Server is running on 9090")
}
