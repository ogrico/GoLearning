package application

import (
	"github.com/ogrico/GoLearning/08-ArquitecturaHexagonal/Demo/internal/domain"
	"github.com/ogrico/GoLearning/08-ArquitecturaHexagonal/Demo/internal/ports"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *domain.User) error {
	return s.repo.Create(user)
}
