package repositories

import "github.com/PParist/go-bank-service/entities"

type AuthRepository interface {
	Login(string) (*entities.User, error)
}
