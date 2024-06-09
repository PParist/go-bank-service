package service

import "github.com/PParist/go-bank-service/entities"

type JwtService interface {
	GetAllRole() (*[]entities.Role, error)
	ValidateRole(string, []string) error
}
