package service

import "github.com/PParist/go-bank-service/entities"

type UserService interface {
	CreateUser(*entities.User) error
	GetUsers() ([]entities.User, error)
	GetUserByID(int) (*entities.User, error)
	UpdateUserByID(int, entities.User) (*entities.User, error)
	DeleteUser(int) error
}
