package models

import (
	"address-book-server-v3/internal/common/types"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id []byte `gorm:"primaryKey;autoIncrement"`

	Email    string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(255);not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Addresses []Address `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,strict_email"`
	Password string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,strict_email"`
	Password string `json:"password" validate:"required"`
}

type UserCmdOutputData struct {
	Id       types.UserId
	Email    string
	Password string
}

func NewUserCmdOutputData(user *User) *UserCmdOutputData {
	return &UserCmdOutputData{
		Id: types.UserId(user.Id),
		Email: user.Email,
		Password: user.Password,
	}
}
