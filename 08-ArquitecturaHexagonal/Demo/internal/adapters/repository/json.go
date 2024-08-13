// internal/adapters/repository/json.go
package repository

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"Demo/internal/domain"
)

type JSONRepository struct {
	filePath string
	data     map[string]interface{}
}

func NewJSONRepository(filePath string) *JSONRepository {
	return &JSONRepository{filePath: filePath}
}

func (repo *JSONRepository) loadData() error {
	file, err := os.Open(repo.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var data map[string]interface{}
	if err := json.Unmarshal(byteValue, &data); err != nil {
		return err
	}
	repo.data = data
	return nil
}

func (repo *JSONRepository) GetAllUsers() ([]domain.User, error) {
	if err := repo.loadData(); err != nil {
		return nil, err
	}

	usersJson, ok := repo.data["users"].([]interface{})
	if !ok {
		return nil, errors.New("users data format error")
	}

	var users []domain.User
	for _, userJson := range usersJson {
		userMap := userJson.(map[string]interface{})
		user := domain.User{
			ID:    userMap["id"].(string),
			Name:  userMap["name"].(string),
			Email: userMap["email"].(string),
		}
		users = append(users, user)
	}
	return users, nil
}

func (repo *JSONRepository) GetUserByID(id string) (*domain.User, error) {
	users, err := repo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (repo *JSONRepository) GetAllCredits() ([]domain.Credit, error) {
	// Similar to GetAllUsers, but parsing the credits data
	// Implement the logic here
	return nil, nil
}

func (repo *JSONRepository) GetCreditByID(id string) (*domain.Credit, error) {
	// Similar to GetUserByID, but parsing the credits data
	// Implement the logic here
	return nil, nil
}
