package repositories

import "github.com/PParist/go-bank-service/entities"

type AccountRepository interface {
	Create(entities.Account) (*entities.Account, error)
	GetAll() (*[]entities.Account, error)
	GetByUserUID(string) (*[]entities.Account, error)
	UpdateByUID(string, map[string]interface{}) (*entities.Account, error)
	DeleteByUID(string) error
}
