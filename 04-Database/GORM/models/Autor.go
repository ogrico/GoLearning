package models

type Autor struct {
	ID           int64  `json:"id"`
	Nombre       string `json:"nombre"`
	Nacionalidad string `json:"nacionalidad"`
}

func (Autor) TableName() string {
	return "autores"
}
