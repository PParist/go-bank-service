package repositories

import (
	"github.com/PParist/go-bank-service/entities"
	"gorm.io/gorm"
)

type authGormRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authGormRepository{db: db}
}

func (r *authGormRepository) Login(username string) (*entities.User, error) {

	user := new(entities.User)
	if result := r.db.Where("username = ?", username).First(user); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
