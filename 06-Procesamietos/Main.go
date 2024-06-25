package main

import (
	"fmt"
	"log"
	"net/http"
	"procesamientos/router"
)

func main() {
	fmt.Println("Hello World")
	router := router.NewRouter()
	fmt.Println("Server is running on 9090")
	log.Fatal(http.ListenAndServe(":9090", router))
}
