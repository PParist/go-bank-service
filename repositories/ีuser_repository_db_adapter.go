package repositories

import (
	"fmt"

	"github.com/PParist/go-bank-service/entities"
	"github.com/PParist/go-bank-service/logs"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type gormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) UserRepository {
	return &gormUserRepository{db: db}
}

func (r *gormUserRepository) Create(user entities.User) error {

	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var existinguser entities.User
	result := tx.Where("username = ?", user.Username).Or("email = ?", user.Email).First(&existinguser)
	if result.Error != nil {
		if result := r.db.Create(&user); result.Error != nil {
			return result.Error
		}
		return nil
	}

	tx.Rollback()
	logs.Error(result.Error)
	return fmt.Errorf("user already exists")
}
func (r *gormUserRepository) GetAll() (*[]entities.User, error) {
	var users = []entities.User{}
	if result := r.db.Preload(clause.Associations).Find(&users); result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}
func (r *gormUserRepository) GetByID(id int) (*entities.User, error) {
	user := entities.User{}
	if result := r.db.Where("id = ?", id).Preload("Profile").First(&user); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (r *gormUserRepository) UpdateByID(id int, updateFields map[string]interface{}) (*entities.User, error) {

	tx := r.db.Begin()
	user := new(entities.User)
	result := tx.Model(user).Where("id = ?", id).Updates(&updateFields)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	result = tx.Where("id = ?", id).Preload("Profile").First(user)
	if result.Error != nil {
		tx.Rollback()
	}
	return user, tx.Commit().Error
}
func (r *gormUserRepository) DeleteByID(id int) error {
	user := entities.User{}
	if result := r.db.Where("id = ?", id).Delete(&user); result.Error != nil {
		return result.Error
	}
	return nil
}
