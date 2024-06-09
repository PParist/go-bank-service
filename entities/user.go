package entities

import (
	"time"

	"gorm.io/gorm"
)

type UserRequest struct {
	Username string    `gorm:"unique" validate:"required,min=3,max=32" json:"username"`
	Password string    `validate:"required,min=8" json:"password"`
	Email    string    `gorm:"unique" validate:"required,email" json:"email"`
	Role     string    `gorm:"foreignKey:Role;references:RoleName" json:"role"`
	Profile  []Profile `gorm:"foreignKey:User_uid;references:User_uid" json:"profiles"`
}

type UseResponse struct {
	Id       uint      `validate:"required" json:"id"`
	User_uid string    `gorm:"unique" validate:"required,uuid4" json:"user_uid"`
	Username string    `gorm:"unique" validate:"required,min=3,max=32" json:"username"`
	Password string    `validate:"required,min=8" json:"password"`
	Email    string    `gorm:"unique" validate:"required,email" json:"email"`
	Profile  []Profile `gorm:"foreignKey:User_uid;references:User_uid" json:"profiles"`
	Accounts []Account `gorm:"foreignKey:User_uid;references:User_uid" json:"account"`
}

type User struct {
	gorm.Model
	User_uid  string    `gorm:"unique" validate:"required,uuid4" json:"user_uid"`
	Username  string    `gorm:"unique" validate:"required,min=3,max=32" json:"username"`
	Password  string    `validate:"required,min=8" json:"password"`
	Email     string    `gorm:"unique" validate:"required,email" json:"email"`
	Profile   []Profile `gorm:"foreignKey:User_uid;references:User_uid" json:"profiles"`
	Accounts  []Account `gorm:"foreignKey:User_uid;references:User_uid" json:"account"`
	User_Role string    `gorm:"foreignKey:User_Role;references:RoleName" validate:"required" json:"role"`
}

type Profile struct {
	gorm.Model
	Profile_uid string    `gorm:"unique" validate:"required,uuid4" json:"profile_uid"`
	User_uid    string    `gorm:"unique" validate:"required,uuid4" json:"user_uid"`
	FirstName   string    `validate:"required" json:"firstname"`
	LastName    string    `validate:"required" json:"lastname"`
	Gender      string    `validate:"required" json:"gender"`
	Birthday    time.Time `validate:"required" json:"birthday"`
	Address     string    `validate:"required" json:"address"`
	Image       string    `json:"image"`
}

type UserRole struct {
	gorm.Model
	Role_uid string `gorm:"unique" validate:"required,uuid4"`
	Name     string ` json:"name"`
}
