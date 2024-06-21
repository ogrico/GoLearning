package handlers

import (
	"gorm/db"
	"gorm/models"
	"net/http"
)

func GetAutors(rw http.ResponseWriter, req *http.Request) {
	var autores []models.Autor
	db.Database().Find(&autores)
	sendData(rw, autores, http.StatusOK)
}
