package models

import (
	"address-book-server-v3/internal/common/types"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id types.UserId `gorm:"primaryKey;autoIncrement"`

	Email    string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(255);not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Addresses []Address `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.Id = types.UserId(uuid.New())
	return
}

type RegisterRequest struct {
	Body *RegisterRequestBody
}

type RegisterRequestBody struct {
	Email    string `json:"email" validate:"required,strict_email"`
	Password string `json:"password" validate:"required"`
}

type RegisterResponse struct {
	Id uuid.UUID
	Email string
}

func NewRegisterResponse(id uuid.UUID, email string) *RegisterResponse {
	return &RegisterResponse{
		Id: id,
		Email: email,
	}
}

type LoginRequestBody struct {
	Email    string `json:"email" validate:"required,strict_email"`
	Password string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Body *LoginRequestBody
}

type LoginResponse struct {
	Token *string
}

func NewLoginResponse(token string) *LoginResponse {
	return &LoginResponse{
		Token: &token,
	}
}

type LoginCmdOutputData struct {
	Token *string
}

func NewLoginCmdOutputData(token string) *LoginCmdOutputData {
	return &LoginCmdOutputData{
		Token: &token,
	}
}

type UserCmdOutputData struct {
	Id       types.UserId
	Email    string
}

func (UserCmdOutputData) IsCmdOutput() {}

func (LoginCmdOutputData) IsCmdOutput() {}

func NewUserCmdOutputData(user *User) *UserCmdOutputData {
	return &UserCmdOutputData{
		Id: types.UserId(user.Id),
		Email: user.Email,
	}
}
