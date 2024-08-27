package mysql

import (
	"AuthApp/internal/domain"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type UserRepository struct {
	DB *sql.DB
}

// FindByUsername busca al usuario
func (r *UserRepository) FindByUsername(username string) (domain.User, error) {
	var user domain.User
	err := r.DB.QueryRow("SELECT id, username, password FROM user WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Create permite crear usuarios
func (r *UserRepository) Create(user domain.User) error {
	_, err := r.DB.Exec("INSERT INTO user (id, username, password) VALUES (?, ?, ?)", user.ID, user.Username, user.Password)
	return err
}
