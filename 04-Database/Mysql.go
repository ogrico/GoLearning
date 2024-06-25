package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type Autor struct {
	ID           int    `json:"id"`
	Nombre       string `json:"nombre"`
	Nacionalidad string `json:"nacionalidad"`
}

func getAll(db *sql.DB) ([]Autor, error) {
	rows, err := db.Query("SELECT id, nombre, nacionalidad FROM autores")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var autores []Autor
	for rows.Next() {
		var autor Autor
		if err := rows.Scan(&autor.ID, &autor.Nombre, &autor.Nacionalidad); err != nil {
			return nil, err
		}
		autores = append(autores, autor)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return autores, nil
}

func newAutor(db *sql.DB, nombre string, nacionalidad string) (Autor, error) {

	result, err := db.Exec("INSERT INTO autores (nombre, nacionalidad) VALUES (?, ?)", nombre, nacionalidad)
	if err != nil {
		return Autor{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Autor{}, err
	}

	nuevoAutor := Autor{
		ID:           int(id),
		Nombre:       nombre,
		Nacionalidad: nacionalidad,
	}

	return nuevoAutor, nil
}

func main() {
	fmt.Println("Hello World")
	dsn := "geovany:123456@tcp(127.0.0.1:3306)/go"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	defer db.Close()

	autores, err := getAll(db)
	if err != nil {
		log.Fatal(err)
	}

	for _, autor := range autores {
		fmt.Printf("ID: %d, Nombre: %s, Nacionalidad: %s\n", autor.ID, autor.Nombre, autor.Nacionalidad)
	}

	nuevoAutor, err := newAutor(db, "Test3", "Colombiano")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Autor creado:", nuevoAutor)

}
