package ports

import "github.com/ogrico/GoLearning/08-ArquitecturaHexagonal/Demo/internal/domain"

type UserRepository interface {
	GetByID(id string) (*domain.User, error)
	Create(user *domain.User) error
	Update(user *domain.User) error
	Delete(id string) error
}

type ProductRepository interface {
	GetAccountByUserID(userID string) (*domain.Account, error)
	GetCreditByUserID(userID string) (*domain.Credit, error)
	GetCreditCardByUserID(userID string) (*domain.CreditCard, error)
}
