package service

import (
	"fmt"

	"github.com/PParist/go-bank-service/entities"
	"github.com/PParist/go-bank-service/errorhandler"
	"github.com/PParist/go-bank-service/logs"
	"github.com/PParist/go-bank-service/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (r *userService) CreateUser(_user *entities.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(_user.Password), bcrypt.DefaultCost)
	if err != nil {
		logs.Error(err)
		return errorhandler.NewErrorBadRequest("invalid password")
	}
	_user.Password = string(hashedPassword)
	_user.User_uid = uuid.New().String()
	validate := validator.New()
	// Validate the user struct
	if err := validate.Struct(_user); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			logs.Error(err)
			return errorhandler.NewErrorBadRequest(fmt.Sprintf("field '%s' failed validation with tag '%s'", err.Field(), err.Tag()))
		}
	}
	if err := r.repo.Create(*_user); err != nil {
		logs.Error(err)
		return errorhandler.NewErrorInternalServerError(err.Error())
	}
	return nil
}
func (r *userService) GetUsers() ([]entities.User, error) {
	if users, err := r.repo.GetAll(); err != nil {
		logs.Error(err)
		return nil, errorhandler.NewErrorInternalServerError(err.Error())
	} else {
		//TODO:กรณีไม่ต้องการให้ primary adapter รู้จัก data หลังบ้านหรือต้องการ return data แค่เฉพาะส่วนให้ทำ DTO
		// userRespons := []entities.UseResponse{}
		// for _, user := range *users {
		// 	respons := entities.UseResponse{
		// 		Id:       user.ID,
		// 		User_uid: user.User_uid,
		// 		Username: user.Username,
		// 		Password: user.Password,
		// 		Email:    user.Email,
		// 	}
		// 	userRespons = append(userRespons, respons)
		// }
		return *users, nil
	}

}
func (r *userService) GetUserByID(id int) (*entities.User, error) {
	if user, err := r.repo.GetByID(id); err != nil {
		//logs.Error(err)
		if err == gorm.ErrRecordNotFound {
			return nil, errorhandler.NewErrorNotFound("user not found")
		}
		return nil, errorhandler.NewErrorInternalServerError(err.Error())
	} else {
		//TODO:กรณีไม่ต้องการให้ primary adapter รู้จัก data หลังบ้านหรือต้องการ return data แค่เฉพาะส่วนให้ทำ DTO
		return user, nil
	}

}
func (r *userService) UpdateUserByID(id int, user entities.User) (*entities.User, error) {
	// TODO: DTO สร้าง map เพื่อเก็บฟิลด์และค่าที่ต้องการอัปเดต
	updateFields := make(map[string]interface{})
	// เพิ่มฟิลด์และค่าที่ต้องการอัปเดตเข้าไปใน map
	if user.Username != "" {
		updateFields["username"] = user.Username
	}

	if user.Password != "" {

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			logs.Error(err)
			return nil, errorhandler.NewErrorBadRequest("invalid password")
		}

		user.Password = string(hashedPassword)
		updateFields["password"] = user.Password
	}

	if user.Email != "" {
		updateFields["email"] = user.Email
	}

	if user.Profile != nil {
		updateFields["profile"] = user.Profile
	}
	if user.User_Role != "" {
		// if user.User_Role == "Admin" && roleToken != "Admin" {
		// 	user.User_Role = ""
		// }
		updateFields["user_role"] = user.User_Role
	}

	if users, err := r.repo.UpdateByID(id, updateFields); err != nil {
		//logs.Error(err)
		if err == gorm.ErrRecordNotFound {
			return nil, errorhandler.NewErrorNotFound("user not found")
		}
		return nil, errorhandler.NewErrorInternalServerError(err.Error())
	} else {
		//TODO:กรณีไม่ต้องการให้ primary adapter รู้จัก data หลังบ้านหรือต้องการ return data แค่เฉพาะส่วนให้ทำ DTO
		// user := &entities.User{}
		// for key, v := range *users {
		// 	switch key {
		// 	case "username":
		// 		user.Username = v.(string)
		// 	case "password":
		// 		user.Password = v.(string)
		// 	case "email":
		// 		user.Email = v.(string)
		// 	case "profile":
		// 		user.Profile = v.([]entities.Profile)
		// 	}
		// }
		return users, nil
	}
}
func (r *userService) DeleteUser(id int) error {
	if err := r.repo.DeleteByID(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			return errorhandler.NewErrorNotFound("user not found")
		}
		return errorhandler.NewErrorInternalServerError(err.Error())
	} else {
		//TODO:กรณีไม่ต้องการให้ primary adapter รู้จัก data หลังบ้านหรือต้องการ return data แค่เฉพาะส่วนให้ทำ DTO
		return nil
	}

}
