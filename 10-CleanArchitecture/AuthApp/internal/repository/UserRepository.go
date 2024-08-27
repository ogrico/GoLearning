package repository

import "AuthApp/internal/domain"

type UserRepository interface {
	FindByUsername(username string) (domain.User, error)
	Create(user domain.User) error
}
