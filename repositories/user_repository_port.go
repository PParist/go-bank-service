package repositories

import "github.com/PParist/go-bank-service/entities"

type UserRepository interface {
	Create(entities.User) error
	GetAll() (*[]entities.User, error)
	GetByID(id int) (*entities.User, error)
	UpdateByID(int, map[string]interface{}) (*entities.User, error)
	DeleteByID(id int) error
}
