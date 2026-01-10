package models

import (
	"address-book-server-v3/internal/common/types"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id []byte `gorm:"type:binary(16);primaryKey;autoIncrement"`

	Email    string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(255);not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Addresses []Address `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}

// func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
// 	id := uuid.New()
// 	user.Id = id[:]
// 	return
// }

type RegisterRequest struct {
	Body *RegisterRequestBody
}

type RegisterRequestBody struct {
	Email    string `json:"email" validate:"required,strict_email"`
	Password string `json:"password" validate:"required,password"`
}

type RegisterResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

func NewRegisterResponse(id types.UserId, email string) *RegisterResponse {
	return &RegisterResponse{
		Id:    id.String(),
		Email: email,
	}
}

type LoginRequestBody struct {
	Email    string `json:"email" validate:"required,strict_email"`
	Password string `json:"password" validate:"required,password"`
}

type LoginRequest struct {
	Body *LoginRequestBody
}

type LoginResponse struct {
	Token string `json:"token"`
}

func NewLoginResponse(token string) *LoginResponse {
	return &LoginResponse{
		Token: token,
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
	Id    types.UserId
	Email string
}

func (UserCmdOutputData) IsCmdOutput() {}

func (LoginCmdOutputData) IsCmdOutput() {}

func NewUserCmdOutputData(user *User) *UserCmdOutputData {
	_id, _ := uuid.FromBytes(user.Id)

	return &UserCmdOutputData{
		Id:    types.UserId(_id),
		Email: user.Email,
	}
}
