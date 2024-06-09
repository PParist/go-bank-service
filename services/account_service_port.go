package service

import "github.com/PParist/go-bank-service/entities"

type AccountService interface {
	CreateAccount(string, entities.NewAccountRequest) (*entities.AccountRespons, error)
	GetAccounts() (*[]entities.AccountRespons, error)
	GetAccountByUserUID(string) (*[]entities.AccountRespons, error)
	UpdateAccount(string, entities.AccountUpdateRequest) (*entities.AccountRespons, error)
	DeleteAccountByUID(string) error
}
