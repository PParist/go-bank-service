package repositories

import "github.com/PParist/go-bank-service/entities"

type JwtRepository interface {
	GetRole() (*[]entities.Role, error)
}
