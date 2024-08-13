package repository

import (
	"encoding/json"
	"io/ioutil"

	"github.com/ogrico/GoLearning/08-ArquitecturaHexagonal/Demo/internal/domain"
)

type JSONUserRepository struct {
	filePath string
}

// Create implements ports.UserRepository.
func (r *JSONUserRepository) Create(user *domain.User) error {
	panic("unimplemented")
}

// Delete implements ports.UserRepository.
func (r *JSONUserRepository) Delete(id string) error {
	panic("unimplemented")
}

// Update implements ports.UserRepository.
func (r *JSONUserRepository) Update(user *domain.User) error {
	panic("unimplemented")
}

func NewJSONUserRepository(filePath string) *JSONUserRepository {
	return &JSONUserRepository{filePath: filePath}
}

func (r *JSONUserRepository) GetByID(id string) (*domain.User, error) {
	data, err := ioutil.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}
	var users []domain.User
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, nil
}
