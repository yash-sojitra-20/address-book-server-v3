package services

import (
	"address-book-server-v3/internal/common/types"
	"address-book-server-v3/internal/common/utils"
	"address-book-server-v3/internal/models"
	"address-book-server-v3/internal/repositories"
	"fmt"

	"bitbucket.org/vayana/walt-go/command"
	"bitbucket.org/vayana/walt-gorm.go/transaction"
	"github.com/google/uuid"
	"github.com/samber/mo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// type IAddressService interface {
// 	Create(userID uint64, address *models.Address) error
// 	GetAll(userID uint64) ([]models.AddressResponse, error)
// 	GetByID(userID, addressID uint64) (*models.AddressResponse, error)
// 	Update(userID uint64, id uint64, updated *models.Address) error
// 	Delete(userID uint64, id uint64) error
// 	ExportAddressesCustomAsync(userID uint64, fields []string, sendTo string)
// 	GetFilteredAddresses(userID uint64, listAddressQuery models.ListAddressQuery) ([]models.AddressResponse, int64, error)
// }

type CreateAddrCmd struct {
	userId               types.UserId
	createAddressRequest models.CreateAddressRequest
}

func NewCreateAddrCmd(createAddressRequest models.CreateAddressRequest, userId types.UserId) *CreateAddrCmd {
	return &CreateAddrCmd{createAddressRequest: createAddressRequest, userId: userId}
}

func (cmd *CreateAddrCmd) Execute(c command.CmdContext) mo.Result[*models.AddressCmdOutputData] {
	ctx := c.(CommandContext)

	operation := func(db *gorm.DB) mo.Result[*models.AddressCmdOutputData] {
		logger := ctx.GetLogger()

		repoCtx := repositories.NewRepoContext(db, logger)
		repo := repositories.NewAddressRepo(repoCtx)

		address := models.Address{
			UserId:       cmd.userId[:],
			FirstName:    cmd.createAddressRequest.Body.FirstName,
			LastName:     cmd.createAddressRequest.Body.LastName,
			Email:        cmd.createAddressRequest.Body.Email,
			Phone:        cmd.createAddressRequest.Body.Phone,
			AddressLine1: cmd.createAddressRequest.Body.AddressLine1,
			AddressLine2: cmd.createAddressRequest.Body.AddressLine2,
			City:         cmd.createAddressRequest.Body.City,
			State:        cmd.createAddressRequest.Body.State,
			Country:      cmd.createAddressRequest.Body.Country,
			Pincode:      cmd.createAddressRequest.Body.Pincode,
		}
		id := uuid.New()
		address.Id = id[:]

		addr, err := repo.Create(&address).Get()
		if err != nil {
			return mo.Err[*models.AddressCmdOutputData](err)
		}

		_addr := models.NewAddressCmdOutputData(addr)

		return mo.Ok(_addr)
	}

	return transaction.DoInTransaction(ctx.GetDb(), operation)
}

type GetAllAddrCmd struct {
	userId types.UserId
}

func NewGetAllAddrCmd(userId types.UserId) *GetAllAddrCmd {
	return &GetAllAddrCmd{userId: userId}
}

func (cmd *GetAllAddrCmd) Execute(c command.CmdContext) mo.Result[*models.ListAddressCmdOutputData] {
	ctx := c.(CommandContext)

	operation := func(db *gorm.DB) mo.Result[*models.ListAddressCmdOutputData] {
		logger := ctx.GetLogger()

		repoCtx := repositories.NewRepoContext(db, logger)
		repo := repositories.NewAddressRepo(repoCtx)

		addresses, err := repo.FindByUser(uuid.UUID(cmd.userId)).Get()
		if err != nil {
			return mo.Err[*models.ListAddressCmdOutputData](err)
		}

		var _addresses []models.AddressCmdOutputData
		for _, v := range *addresses {
			_addresses = append(_addresses, *models.NewAddressCmdOutputData(&v))
		}

		_addrs := models.NewListAddressCmdOutputData(_addresses)

		return mo.Ok(_addrs)
	}

	return transaction.DoInTransaction(ctx.GetDb(), operation)
}

type GetByIdAddrCmd struct {
	addressId types.AddressId
	userId    types.UserId
}

func NewGetByIdAddrCmd(addressId types.AddressId, userId types.UserId) *GetByIdAddrCmd {
	return &GetByIdAddrCmd{addressId: addressId, userId: userId}
}

func (cmd *GetByIdAddrCmd) Execute(c command.CmdContext) mo.Result[*models.AddressCmdOutputData] {
	ctx := c.(CommandContext)

	operation := func(db *gorm.DB) mo.Result[*models.AddressCmdOutputData] {
		logger := ctx.GetLogger()

		repoCtx := repositories.NewRepoContext(db, logger)
		repo := repositories.NewAddressRepo(repoCtx)

		address, err := repo.FindByID(uuid.UUID(cmd.addressId), uuid.UUID(cmd.userId)).Get()
		if err != nil {
			return mo.Err[*models.AddressCmdOutputData](err)
		}

		_addr := models.NewAddressCmdOutputData(address)

		return mo.Ok(_addr)
	}

	return transaction.DoInTransaction(ctx.GetDb(), operation)
}

type UpdateAddrCmd struct {
	addressId            types.AddressId
	userId               types.UserId
	updateAddressRequest models.UpdateAddressRequest
}

func NewUpdateAddrCmd(addressId types.AddressId, userId types.UserId, updateAddressRequest models.UpdateAddressRequest) *UpdateAddrCmd {
	return &UpdateAddrCmd{addressId: addressId, userId: userId, updateAddressRequest: updateAddressRequest}
}

func (cmd *UpdateAddrCmd) Execute(c command.CmdContext) mo.Result[*models.AddressCmdOutputData] {
	ctx := c.(CommandContext)

	operation := func(db *gorm.DB) mo.Result[*models.AddressCmdOutputData] {
		logger := ctx.GetLogger()

		repoCtx := repositories.NewRepoContext(db, logger)
		repo := repositories.NewAddressRepo(repoCtx)

		address, err := repo.FindByID(uuid.UUID(cmd.addressId), uuid.UUID(cmd.userId)).Get()
		if err != nil {
			return mo.Err[*models.AddressCmdOutputData](err)
		}

		if cmd.updateAddressRequest.Body.FirstName != nil {
			address.FirstName = *cmd.updateAddressRequest.Body.FirstName
		}
		if cmd.updateAddressRequest.Body.LastName != nil {
			address.LastName = *cmd.updateAddressRequest.Body.LastName
		}
		if cmd.updateAddressRequest.Body.Email != nil {
			address.Email = *cmd.updateAddressRequest.Body.Email
		}
		if cmd.updateAddressRequest.Body.Phone != nil {
			address.Phone = *cmd.updateAddressRequest.Body.Phone
		}
		if cmd.updateAddressRequest.Body.AddressLine1 != nil {
			address.AddressLine1 = *cmd.updateAddressRequest.Body.AddressLine1
		}
		if cmd.updateAddressRequest.Body.AddressLine2 != nil {
			address.AddressLine2 = *cmd.updateAddressRequest.Body.AddressLine2
		}
		if cmd.updateAddressRequest.Body.City != nil {
			// fmt.Println("=====================> came in baroda.... *cmd.updateAddressRequest.Body.City", *cmd.updateAddressRequest.Body.City)
			// fmt.Println("=====================> Before : address.City: ", address.City)
			address.City = *cmd.updateAddressRequest.Body.City
			// fmt.Println("=====================> After: address.City: ", address.City)
		}
		if cmd.updateAddressRequest.Body.State != nil {
			address.State = *cmd.updateAddressRequest.Body.State
		}
		if cmd.updateAddressRequest.Body.Country != nil {
			address.Country = *cmd.updateAddressRequest.Body.Country
		}
		if cmd.updateAddressRequest.Body.Pincode != nil {
			address.Pincode = *cmd.updateAddressRequest.Body.Pincode
		}

		address, err = repo.Update(address).Get()

		_addr := models.NewAddressCmdOutputData(address)

		return mo.Ok(_addr)
	}

	return transaction.DoInTransaction(ctx.GetDb(), operation)
}

type DeleteAddrCmd struct {
	addressId types.AddressId
	userId    types.UserId
}

func NewDeleteAddrCmd(addressId types.AddressId, userId types.UserId) *DeleteAddrCmd {
	return &DeleteAddrCmd{addressId: addressId, userId: userId}
}

func (cmd *DeleteAddrCmd) Execute(c command.CmdContext) mo.Result[*models.DeleteCmdOutputData] {
	ctx := c.(CommandContext)

	operation := func(db *gorm.DB) mo.Result[*models.DeleteCmdOutputData] {
		logger := ctx.GetLogger()

		repoCtx := repositories.NewRepoContext(db, logger)
		repo := repositories.NewAddressRepo(repoCtx)

		address, err := repo.FindByID(uuid.UUID(cmd.addressId), uuid.UUID(cmd.userId)).Get()
		if err != nil {
			return mo.Err[*models.DeleteCmdOutputData](err)
		}

		message, err := repo.Delete(address).Get()
		if err != nil {
			return mo.Err[*models.DeleteCmdOutputData](err)
		}

		_message := models.NewDeleteCmdOutputData(*message)

		return mo.Ok(_message)

	}

	return transaction.DoInTransaction(ctx.GetDb(), operation)
}

type ExportAsyncAddrCmd struct {
	UserId types.UserId
	Fields []string
	Email  string
}

func NewExportAsyncAddrCmd(userId types.UserId, fields []string, email string) *ExportAsyncAddrCmd {
	return &ExportAsyncAddrCmd{UserId: userId, Fields: fields, Email: email}
}

func (cmd *ExportAsyncAddrCmd) Execute(c command.CmdContext) mo.Result[*models.ExportAsyncAddrCmdOutoutData] {
	ctx := c.(CommandContext)

	var exportData *[]map[string]interface{}

	operation := func(db *gorm.DB) mo.Result[*models.ExportAsyncAddrCmdOutoutData] {
		logger := ctx.GetLogger()

		repoCtx := repositories.NewRepoContext(db, logger)
		repo := repositories.NewAddressRepo(repoCtx)

		data, err := repo.FindAllForExport(cmd.Fields, uuid.UUID(cmd.UserId)).Get()
		if err != nil {
			return mo.Err[*models.ExportAsyncAddrCmdOutoutData](err)
		}

		exportData = data

		message := models.NewExportAsyncAddrCmdOutputData("Export started")
		return mo.Ok(message)
	}

	result := transaction.DoInTransaction(ctx.GetDb(), operation)
	if result.IsError() {
		return result
	}

	// Async work
	go func() {
		defer func() {
			if r := recover(); r != nil {
				ctx.GetLogger().Error("panic in export", zap.Any("panic in export", r))
			}
		}()

		appCfg := ctx.GetConfig()

		fileDetails, err := utils.GenerateCustomAddressesCSV(
			uint64(uuid.UUID(cmd.UserId).ID()),
			cmd.Fields,
			*exportData,
		).Get()

		if err != nil {
			ctx.GetLogger().Error("csv generation failed", zap.Error(err))
			return
		}

		downloadURL := fmt.Sprintf(
			"%s/downloads/%s",
			appCfg.GetAppUrl(),
			fileDetails.FileName,
		)

		emailBody := fmt.Sprintf(
			"Attached is the custom address report.\n\nDownload link:\n%s",
			downloadURL,
		)

		err = utils.SendEmailWithAttachment(
			appCfg.SMTP_HOST,
			appCfg.SMTP_PORT,
			appCfg.SMTP_USER,
			appCfg.SMTP_PASS,
			cmd.Email,
			"Custom Address CSV Export",
			emailBody,
			fileDetails.FilePath,
		)

		if err != nil {
			ctx.GetLogger().Error("email send failed", zap.Error(err))
		}
	}()

	return result

}

type FilterAddrCmd struct {
	userId          types.UserId
	filterAddrQuery models.FilterAddrQuery
}

func NewFilterAddrCmd(filterAddrQuery models.FilterAddrQuery, userId types.UserId) *FilterAddrCmd {
	return &FilterAddrCmd{filterAddrQuery: filterAddrQuery, userId: userId}
}

func (cmd *FilterAddrCmd) Execute(c command.CmdContext) mo.Result[*models.FilterAddrCmdOutputData] {
	ctx := c.(CommandContext)

	operation := func(db *gorm.DB) mo.Result[*models.FilterAddrCmdOutputData] {
		logger := ctx.GetLogger()

		repoCtx := repositories.NewRepoContext(db, logger)
		repo := repositories.NewAddressRepo(repoCtx)

		if cmd.filterAddrQuery.Body.Page <= 0 {
			cmd.filterAddrQuery.Body.Page = 1
		}
		if cmd.filterAddrQuery.Body.Limit <= 0 {
			cmd.filterAddrQuery.Body.Limit = 10
		}

		result, err := repo.FindFiltered(uuid.UUID(cmd.userId), &cmd.filterAddrQuery).Get()

		if err != nil {
			return mo.Err[*models.FilterAddrCmdOutputData](err)
		}

		var data []models.AddressCmdOutputData
		for _, a := range result.Addresses {
			_addr := models.NewAddressCmdOutputData(&a)
			data = append(data, *_addr)
		}

		return mo.Ok(&models.FilterAddrCmdOutputData{
			Data:  data,
			Total: result.Total,
		})

	}

	return transaction.DoInTransaction(ctx.GetDb(), operation)
}
