package repository

import (
	"database/sql"

	"MVC/internal/model"

	_ "github.com/go-sql-driver/mysql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetAllUsers() ([]model.User, error) {
	rows, err := r.db.Query("SELECT id, name, email FROM usuarios")
	if err != nil {
		return nil, err
	}
	defer rows.Close()	

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		//fmt.Print(user, "\n")
		users = append(users, user)
	}
	return users, nil
}
