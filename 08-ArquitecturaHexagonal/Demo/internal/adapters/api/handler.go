package api

import (
	"encoding/json"
	"net/http"

	"Demo/internal/application"

	"github.com/gorilla/mux"
)

type Handler struct {
	userService *application.UserService
}

func NewHandler(userService *application.UserService) *Handler {
	return &Handler{userService: userService}
}

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Extrae el ID de la URL
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Si el usuario se encuentra, responde con los datos en formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
