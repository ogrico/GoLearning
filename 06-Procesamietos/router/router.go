package router

import (
	"github.com/gorilla/mux"
	"procesamientos/handlers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/cuentas", handlers.GetCuentasBancarias).Methods("GET")
	router.HandleFunc("/api/cuentas/file", handlers.GetFile).Methods("POST")

	return router
}
