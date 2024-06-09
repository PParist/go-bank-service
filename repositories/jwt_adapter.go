package repositories

import (
	"github.com/PParist/go-bank-service/entities"
	"gorm.io/gorm"
)

type jwtRepositoryDB struct {
	db *gorm.DB
}

func NewJWTRepository(db *gorm.DB) JwtRepository {
	return &jwtRepositoryDB{db: db}
}

func (r *jwtRepositoryDB) GetRole() (*[]entities.Role, error) {
	var roles = []entities.Role{}
	if result := r.db.Find(&roles); result.Error != nil {
		return nil, result.Error
	}
	return &roles, nil
}
