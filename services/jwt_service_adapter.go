package service

import (
	"fmt"

	"github.com/PParist/go-bank-service/entities"
	"github.com/PParist/go-bank-service/errorhandler"
	"github.com/PParist/go-bank-service/logs"
	"github.com/PParist/go-bank-service/repositories"
)

type jwtService struct {
	repo repositories.JwtRepository
}

func NewJWTService(repo repositories.JwtRepository) JwtService {
	return &jwtService{repo: repo}
}

func (s *jwtService) GetAllRole() (*[]entities.Role, error) {
	if roles, err := s.repo.GetRole(); err != nil {
		logs.Error(err)
		if len(*roles) <= 0 {
			return nil, errorhandler.NewErrorNotFound("role not found")
		}
		return nil, errorhandler.NewErrorInternalServerError(err.Error())
	} else {
		return roles, nil
	}

}

func (s *jwtService) ValidateRole(roleHeader string, roles []string) error {
	var isMacth bool
	if result, err := s.repo.GetRole(); err != nil {
		logs.Error(err)
		if len(*result) <= 0 {
			return errorhandler.NewErrorNotFound("role not found")
		}
		return errorhandler.NewErrorInternalServerError(err.Error())

	} else {
		//Convert []role to map role
		roleMap := make(map[string]string)
		for _, v := range *result {
			roleMap[v.RoleName] = v.RoleName
		}

		//Check role in map
		role, exists := roleMap[roleHeader]
		if !exists {
			fmt.Println(roleHeader)
			return errorhandler.NewErrorForbidden("Invalid token")
		}
		for _, r := range roles {
			if role == r {
				isMacth = true
			}
		}
	}

	if !isMacth {
		return errorhandler.NewErrorForbidden("Permission denied")
	}

	return nil
}
