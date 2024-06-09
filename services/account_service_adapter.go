package service

import (
	"fmt"

	"github.com/PParist/go-bank-service/entities"
	"github.com/PParist/go-bank-service/errorhandler"
	"github.com/PParist/go-bank-service/logs"
	"github.com/PParist/go-bank-service/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var validate = validator.New()

type accountService struct {
	repo repositories.AccountRepository
}

func NewAccountService(repo repositories.AccountRepository) AccountService {
	return &accountService{repo: repo}
}

func (s *accountService) CreateAccount(user_uid string, account entities.NewAccountRequest) (*entities.AccountRespons, error) {

	if err := validate.Struct(account); err != nil {
		fmt.Println("validate")
		for _, err := range err.(validator.ValidationErrors) {
			logs.Error(err)
			return nil, errorhandler.NewErrorBadRequest(fmt.Sprintf("field '%s' failed validation with tag '%s'", err.Field(), err.Tag()))
		}
	}
	createAccount := entities.Account{
		User_uid:     user_uid,
		Account_uid:  uuid.New().String(),
		Account_Type: account.Account_Type,
		Balance:      account.Balance,
		IsActive:     true,
	}

	if result, err := s.repo.Create(createAccount); err != nil {
		logs.Error(err)
		return nil, errorhandler.NewErrorInternalServerError(err.Error())
	} else {
		fmt.Printf("err is %#v", *result)
		if err := validate.Struct(result); err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				logs.Error(err)
				return nil, errorhandler.NewErrorInternalServerError(fmt.Sprintf("field '%s' failed validation with tag '%s'", err.Field(), err.Tag()))
			}
		}
		responsAccount := entities.AccountRespons{
			ID:           result.ID,
			User_uid:     user_uid,
			Account_uid:  result.Account_uid,
			CreatedAt:    result.CreatedAt,
			Account_Type: result.Account_Type,
			Balance:      result.Balance,
			IsActive:     result.IsActive,
		}
		return &responsAccount, nil
	}
}

func (s *accountService) GetAccounts() (*[]entities.AccountRespons, error) {
	if result, err := s.repo.GetAll(); err != nil {
		logs.Error(err)
		return nil, errorhandler.NewErrorInternalServerError(err.Error())
	} else {
		accountsRespons := []entities.AccountRespons{}
		for _, account := range *result {
			accountRespons := entities.AccountRespons{
				ID:           account.ID,
				User_uid:     account.User_uid,
				Account_uid:  account.Account_uid,
				CreatedAt:    account.CreatedAt,
				Account_Type: account.Account_Type,
				Balance:      account.Balance,
				IsActive:     account.IsActive,
			}
			accountsRespons = append(accountsRespons, accountRespons)
		}
		return &accountsRespons, nil
	}
}

func (s *accountService) GetAccountByUserUID(user_uid string) (*[]entities.AccountRespons, error) {
	validate := validator.New()
	if err := validate.Var(user_uid, "required,uuid4"); err != nil {
		logs.Error(err)
		return nil, errorhandler.NewErrorBadRequest("invalid format user_uid")
	}
	if result, err := s.repo.GetByUserUID(user_uid); err != nil {
		logs.Error(err)
		if err == gorm.ErrRecordNotFound {
			return nil, errorhandler.NewErrorNotFound(err.Error())
		}
		return nil, errorhandler.NewErrorInternalServerError(err.Error())
	} else {
		accountsRespons := []entities.AccountRespons{}
		for _, account := range *result {
			accountRespons := entities.AccountRespons{
				ID:           account.ID,
				User_uid:     account.User_uid,
				Account_uid:  account.Account_uid,
				CreatedAt:    account.CreatedAt,
				Account_Type: account.Account_Type,
				Balance:      account.Balance,
				IsActive:     account.IsActive,
			}
			accountsRespons = append(accountsRespons, accountRespons)
		}
		return &accountsRespons, nil
	}
}

func (s *accountService) UpdateAccount(accountUid string, account entities.AccountUpdateRequest) (*entities.AccountRespons, error) {
	if err := validate.Struct(account); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			logs.Error(err)
			return nil, errorhandler.NewErrorBadRequest(fmt.Sprintf("field '%s' failed validation with tag '%s'", err.Field(), err.Tag()))
		}
	}
	updateFields := make(map[string]interface{})

	updateFields["balance"] = account.Balance
	updateFields["is_active"] = account.IsActive

	if result, err := s.repo.UpdateByUID(accountUid, updateFields); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorhandler.NewErrorNotFound(err.Error())
		}
		return nil, errorhandler.NewErrorInternalServerError(err.Error())
	} else {
		respons := entities.AccountRespons{
			ID:           result.ID,
			User_uid:     result.User_uid,
			Account_uid:  result.Account_uid,
			CreatedAt:    result.CreatedAt,
			Account_Type: result.Account_Type,
			Balance:      result.Balance,
			IsActive:     result.IsActive,
		}
		return &respons, nil
	}
}

func (s *accountService) DeleteAccountByUID(accountUid string) error {
	if err := validate.Var(accountUid, "required,uuid4"); err != nil {
		logs.Error(err)
		return errorhandler.NewErrorBadRequest("invalid format account_uid")
	}
	if err := s.repo.DeleteByUID(accountUid); err != nil {
		if err == gorm.ErrRecordNotFound {
			return errorhandler.NewErrorNotFound(err.Error())
		}
		return errorhandler.NewErrorInternalServerError(err.Error())
	}
	return nil
}
