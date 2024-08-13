package application

import (
	"Demo/internal/domain"
	"Demo/internal/ports"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]domain.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) GetUserByID(id string) (*domain.User, error) {
	return s.repo.GetUserByID(id)
}
