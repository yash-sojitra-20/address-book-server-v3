package models

import (
	"address-book-server-v3/internal/common/types"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	Id []byte `gorm:"type:binary(16);primaryKey;autoIncrement"`

	UserId []byte `gorm:"type:binary(16);index;not null"`

	FirstName string `gorm:"type:varchar(100);not null"`
	LastName  string `gorm:"type:varchar(100)"`
	Email     string `gorm:"type:varchar(255);index;not null"`
	Phone     string `gorm:"type:varchar(20)" validate:"omitempty"`

	AddressLine1 string `gorm:"type:varchar(255);not null"`
	AddressLine2 string `gorm:"type:varchar(255)"`
	City         string `gorm:"type:varchar(100)"`
	State        string `gorm:"type:varchar(100)"`
	Country      string `gorm:"type:varchar(100)"`
	Pincode      string `gorm:"type:varchar(20)"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (address *Address) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.New()
	address.Id = id[:]
	return
}

type CreateAddressRequest struct {
	Body *CreateAddressRequestBody
}

type CreateAddressRequestBody struct {
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name"`
	Email        string `json:"email" validate:"required,strict_email"`
	Phone        string `json:"phone" validate:"omitempty,phone"`
	AddressLine1 string `json:"address_line1" validate:"required"`
	AddressLine2 string `json:"address_line2"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	Pincode      string `json:"pincode" validate:"omitempty,pincode"`
}

type GetByIdRequest struct {
	Body *GetByIdRequestBody
}

type GetByIdRequestBody struct {
	AddressId types.AddressId
}

type UpdateAddressRequest struct {
	AddressId types.AddressId
	Body *UpdateAddressRequestBody
}

type UpdateAddressRequestBody struct {
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
	Email        *string `json:"email" validate:"omitempty,strict_email"`
	Phone        *string `json:"phone" validate:"omitempty,phone"`
	AddressLine1 *string `json:"address_line1"`
	AddressLine2 *string `json:"address_line2"`
	City         *string `json:"city"`
	State        *string `json:"state"`
	Country      *string `json:"country"`
	Pincode      *string `json:"pincode" validate:"omitempty,pincode"`
}

type AddressCmdOutputData struct {
	Id           types.AddressId
	UserId       types.UserId
	FirstName    string
	LastName     string
	Email        string
	Phone        string
	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	Country      string
	Pincode      string
}

func NewAddressCmdOutputData(address *Address) *AddressCmdOutputData {
	return &AddressCmdOutputData{
		Id:           types.AddressId(address.Id),
		UserId:       types.UserId(address.UserId),
		FirstName:    address.FirstName,
		LastName:     address.LastName,
		Email:        address.Email,
		Phone:        address.Phone,
		AddressLine1: address.AddressLine1,
		AddressLine2: address.AddressLine2,
		City:         address.City,
		State:        address.State,
		Country:      address.Country,
		Pincode:      address.Pincode,
	}
}

func (AddressCmdOutputData) IsCmdOutput() {}

type ListAddressCmdOutputData struct {
	Addresses []AddressCmdOutputData
}

func NewListAddressCmdOutputData(addresses []AddressCmdOutputData) *ListAddressCmdOutputData {
	return &ListAddressCmdOutputData{
		Addresses: addresses,
	}
}

func (ListAddressCmdOutputData) IsCmdOutput() {}

type DeleteRequest struct {
	Body *DeleteRequestBody
}

type DeleteRequestBody struct {
	AddressId types.AddressId
}

type DeleteCmdOutputData struct {
	Message string
}

func NewDeleteCmdOutputData(message string) *DeleteCmdOutputData {
	return &DeleteCmdOutputData{
		Message: message,
	}
}

func (DeleteCmdOutputData) IsCmdOutput() {}

type ExportAddressRequestBody struct {
	Fields []string `json:"fields" validate:"required,min=1"`
	Email  string   `json:"email" validate:"required,strict_email"`
}

type ExportAddressRequest struct {
	Body *ExportAddressRequestBody
}

type ExportAsyncAddrCmdOutoutData struct {
	Message string
}

func NewExportAsyncAddrCmdOutputData(message string) *ExportAsyncAddrCmdOutoutData {
	return &ExportAsyncAddrCmdOutoutData{
		Message: message,
	}
}

func (ExportAsyncAddrCmdOutoutData) IsCmdOutput() {}

type FilterAddrQueryBody struct {
	Page    int    `form:"page"`
	Limit   int    `form:"limit"`
	Search  string `form:"search"`
	City    string `form:"city"`
	State   string `form:"state"`
	Country string `form:"country"`
	Pincode string `form:"pincode"`
}

type FilterAddrQuery struct {
	Body *FilterAddrQueryBody
}

type FilterAddrCmdOutputData struct {
	Data  []AddressCmdOutputData
	Total int64
}

func NewFilterAddrCmdOutputData(data []AddressCmdOutputData, total int64) *FilterAddrCmdOutputData {
	return &FilterAddrCmdOutputData{
		Data: data,
		Total: total,
	}
}

func (FilterAddrCmdOutputData) IsCmdOutput() {}

type ListAllAddrRequest struct{}
