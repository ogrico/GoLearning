package http

import (
	"encoding/json"
	"net/http"

	"github.com/ogrico/GoLearning/08-ArquitecturaHexagonal/Demo/internal/application"
	"github.com/ogrico/GoLearning/08-ArquitecturaHexagonal/Demo/internal/domain"
)

func CreateUserHandler(service *application.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user domain.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := service.CreateUser(&user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}
