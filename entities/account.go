package entities

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	User_uid     string  `validate:"required,uuid4" json:"user_uid"`
	Account_uid  string  `gorm:"unique" validate:"required,uuid4" json:"account_uid"`
	Account_Type string  `validate:"required" json:"account_type"`
	Balance      float64 `validate:"required" json:"balance"`
	IsActive     bool    `validate:"required" json:"is_active"`
}

type NewAccountRequest struct {
	Account_Type string  `validate:"required" json:"account_type"`
	Balance      float64 `validate:"required" json:"balance"`
}

type AccountUpdateRequest struct {
	Balance  float64 `validate:"required" json:"balance"`
	IsActive bool    `validate:"required" json:"is_active"`
}
type AccountDeleteRequest struct {
	Account_uid string `gorm:"unique" validate:"required,uuid4" json:"account_uid"`
}

type AccountRespons struct {
	ID           uint      `gorm:"primarykey" json:"id"`
	User_uid     string    `validate:"required,uuid4" json:"user_uid"`
	Account_uid  string    `gorm:"unique" validate:"required,uuid4" json:"account_uid"`
	CreatedAt    time.Time `validate:"required" json:"create_at"`
	Account_Type string    `validate:"required" json:"account_type"`
	Balance      float64   `validate:"required" json:"balance"`
	IsActive     bool      `validate:"is_active"`
}

type Transition struct {
	gorm.Model
	Account_uid      int64     `validate:"required,uuid4" json:"account_uid"`
	Amount           float64   // จำนวนเงินที่ทำการทำธุรกรรม
	Transaction_Type string    // ประเภทของการทำธุรกรรม
	Transaction_Date time.Time // วันที่และเวลาของการทำธุรกรรม
	Description      string    // คำอธิบายเพิ่มเติมเกี่ยวกับการทำธุรกรรม
}
