package usecase

import (
	"AuthApp/internal/domain"
	"AuthApp/internal/repository"
)

type UserUseCase struct {
	UserRepo repository.UserRepository
}

// CreateUser crea un nuevo usuario.
func (uc *UserUseCase) CreateUser(user domain.User) error {
	return uc.UserRepo.Create(user)
}
