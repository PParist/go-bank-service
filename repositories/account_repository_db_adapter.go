package repositories

import (
	"fmt"

	"github.com/PParist/go-bank-service/entities"
	"gorm.io/gorm"
)

type accountRepositoryDB struct {
	db *gorm.DB
}

func NewAccountRepositoryDB(db *gorm.DB) AccountRepository {
	return &accountRepositoryDB{db: db}
}

func (r *accountRepositoryDB) Create(account entities.Account) (*entities.Account, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var existingAccount entities.Account
	if err := tx.Where("account_uid = ?", account.Account_uid).First(&existingAccount); err.Error == nil {
		tx.Rollback()
		return nil, fmt.Errorf("account with UID %s already exists", account.Account_uid)
	}

	if err := tx.Create(&account).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &account, nil
}
func (r *accountRepositoryDB) GetAll() (*[]entities.Account, error) {
	account := []entities.Account{}
	if result := r.db.Find(&account); result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}
func (r *accountRepositoryDB) GetByUserUID(user_uid string) (*[]entities.Account, error) {
	account := []entities.Account{}
	if result := r.db.Where("user_uid = ?", user_uid).Find(&account); result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}
func (r *accountRepositoryDB) UpdateByUID(accountUid string, updateFields map[string]interface{}) (*entities.Account, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	account := entities.Account{}
	result := tx.Model(account).Where("account_uid = ?", accountUid).Updates(&updateFields)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	result = tx.Where("account_uid = ?", accountUid).Find(&account)
	if result.Error != nil {
		tx.Rollback()
	}
	return &account, tx.Commit().Error
}
func (r *accountRepositoryDB) DeleteByUID(account_uid string) error {
	var account entities.Account
	if err := r.db.Where("account_uid = ?", account_uid).First(&account).Error; err != nil {
		return err
	}
	if result := r.db.Where("account_uid = ?", account_uid).Delete(&account); result.Error != nil {
		return result.Error
	}
	return nil
}
