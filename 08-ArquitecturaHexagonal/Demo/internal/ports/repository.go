package ports

import "Demo/internal/domain"

type UserRepository interface {
	GetAllUsers() ([]domain.User, error)
	GetUserByID(id string) (*domain.User, error)
	GetAllCredits() ([]domain.Credit, error)
	GetCreditByID(id string) (*domain.Credit, error)
}
