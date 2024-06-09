package entities

import "gorm.io/gorm"

type UserLogin struct {
	Username string
	Password string
}

type UserToken struct {
	UserUid string `validate:"required"`
	Role    string `validate:"required"`
}

type Role struct {
	gorm.Model
	RoleName    string `gorm:"unique;type:varchar(255);not null" validate:"required" json:"rolename"`
	Discription string `gorm:"type:text" json:"discription"`
}

type Permission struct {
	gorm.Model
	PermissionName string `gorm:"unique;type:varchar(255);not null" validate:"required" json:"permissionname"`
	Discription    string `gorm:"type:text" json:"discription"`
}

type RolePermission struct {
	RoleID       uint
	Role         Role
	PermissionID uint
	Permission   Permission
}
