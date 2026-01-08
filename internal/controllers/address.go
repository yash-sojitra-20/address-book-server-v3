package controllers

import (
	"address-book-server-v3/internal/common/types"
	"address-book-server-v3/internal/common/utils"
	"address-book-server-v3/internal/core/application"
	"address-book-server-v3/internal/models"
	"address-book-server-v3/internal/services"

	"bitbucket.org/vayana/walt-go/command"
	"github.com/google/uuid"
	"github.com/samber/mo"
)

func NewCreateAddrRequest(application application.Application, reqCtx utils.RequestCtx) mo.Result[*models.CreateAddressRequest] {
	body, err := utils.GetDataFromRequestBody[models.CreateAddressRequestBody](reqCtx.GetGinCtx()).Get()

	if err != nil {
		return mo.Err[*models.CreateAddressRequest](err)
	}

	return mo.Ok(&models.CreateAddressRequest{
		Body: body,
	})
}

func CreateAddrRequestController(application application.Application, reqCtx utils.RequestCtx, request *models.CreateAddressRequest) mo.Result[*models.AddressCmdOutputData] {
	bundle := application.GetBundle()
	logger := utils.NewApplicationBaseLogger(application.GetLogger(), reqCtx.GetIP())

	cmdCtx := services.NewCommandContext(application, reqCtx, logger)
	
	cmdOutput, err := command.ExecuteCommand(cmdCtx, services.NewCreateAddrCmd(*request, *cmdCtx.GetUserId())).Get()
	if err != nil {
		logger.Error(utils.PrepareMsg(err, bundle))
		return mo.Err[*models.AddressCmdOutputData](err)
	}

	return mo.Ok(cmdOutput)
}

func NewListAllAddrRequest(application application.Application, reqCtx utils.RequestCtx) mo.Result[*models.ListAllAddrRequest] {
	_, err := utils.GetDataFromRequestBody[*models.ListAllAddrRequest](reqCtx.GetGinCtx()).Get()

	if err != nil {
		return mo.Err[*models.ListAllAddrRequest](err)
	}
	return mo.Ok[*models.ListAllAddrRequest](nil)
}

func ListAllAddrRequestController(application application.Application, reqCtx utils.RequestCtx, request *models.ListAllAddrRequest) mo.Result[*models.ListAddressCmdOutputData] {
	bundle := application.GetBundle()
	logger := utils.NewApplicationBaseLogger(application.GetLogger(), reqCtx.GetIP())

	cmdCtx := services.NewCommandContext(application, reqCtx, logger)
	
	cmdOutput, err := command.ExecuteCommand(cmdCtx, services.NewGetAllAddrCmd(*cmdCtx.GetUserId())).Get()
	if err != nil {
		logger.Error(utils.PrepareMsg(err, bundle))
		return mo.Err[*models.ListAddressCmdOutputData](err)
	}

	return mo.Ok(cmdOutput)
}

func NewGetByIdRequest(application application.Application, reqCtx utils.RequestCtx) mo.Result[*models.GetByIdRequest] {
	addrId, err := uuid.Parse(reqCtx.GetGinCtx().Param("id"))
	if err != nil {
		return mo.Err[*models.GetByIdRequest](err)
	}
	
	return mo.Ok(&models.GetByIdRequest{
		Body: &models.GetByIdRequestBody{
			AddressId: types.AddressId(addrId),
		},
	})
}

func GetByIdRequestController(application application.Application, reqCtx utils.RequestCtx, request *models.GetByIdRequest) mo.Result[*models.AddressCmdOutputData] {
	bundle := application.GetBundle()
	logger := utils.NewApplicationBaseLogger(application.GetLogger(), reqCtx.GetIP())

	cmdCtx := services.NewCommandContext(application, reqCtx, logger)
	
	cmdOutput, err := command.ExecuteCommand(cmdCtx, services.NewGetByIdAddrCmd(types.AddressId(request.Body.AddressId) ,*cmdCtx.GetUserId())).Get()
	if err != nil {
		logger.Error(utils.PrepareMsg(err, bundle))
		return mo.Err[*models.AddressCmdOutputData](err)
	}

	return mo.Ok(cmdOutput)
}

func NewUpdateAddrRequest(application application.Application, reqCtx utils.RequestCtx) mo.Result[*models.UpdateAddressRequest] {
	body, err := utils.GetDataFromRequestBody[models.UpdateAddressRequestBody](reqCtx.GetGinCtx()).Get()
	if err != nil {
		return mo.Err[*models.UpdateAddressRequest](err)
	}

	addrId, err := uuid.Parse(reqCtx.GetGinCtx().Param("id"))
	if err != nil {
		return mo.Err[*models.UpdateAddressRequest](err)
	}

	return mo.Ok(&models.UpdateAddressRequest{
		Body: body,
		AddressId: types.AddressId(addrId),
	})
}

func UpdateAddrRequestController(application application.Application, reqCtx utils.RequestCtx, request *models.UpdateAddressRequest) mo.Result[*models.AddressCmdOutputData] {
	bundle := application.GetBundle()
	logger := utils.NewApplicationBaseLogger(application.GetLogger(), reqCtx.GetIP())

	cmdCtx := services.NewCommandContext(application, reqCtx, logger)
	
	cmdOutput, err := command.ExecuteCommand(cmdCtx, services.NewUpdateAddrCmd(request.AddressId, *cmdCtx.GetUserId(), *request)).Get()
	if err != nil {
		logger.Error(utils.PrepareMsg(err, bundle))
		return mo.Err[*models.AddressCmdOutputData](err)
	}

	return mo.Ok(cmdOutput)
}

func NewDeleteAddrRequest(application application.Application, reqCtx utils.RequestCtx) mo.Result[*models.DeleteRequest] {
	addrId, err := uuid.Parse(reqCtx.GetGinCtx().Param("id"))
	if err != nil {
		return mo.Err[*models.DeleteRequest](err)
	}

	return mo.Ok(&models.DeleteRequest{
		Body: &models.DeleteRequestBody{
			AddressId: types.AddressId(addrId),
		},
	})
}

func DeleteAddrRequestController(application application.Application, reqCtx utils.RequestCtx, request *models.DeleteRequest) mo.Result[*models.DeleteCmdOutputData] {
	bundle := application.GetBundle()
	logger := utils.NewApplicationBaseLogger(application.GetLogger(), reqCtx.GetIP())

	cmdCtx := services.NewCommandContext(application, reqCtx, logger)
	
	cmdOutput, err := command.ExecuteCommand(cmdCtx, services.NewDeleteAddrCmd(request.Body.AddressId, *cmdCtx.GetUserId())).Get()
	if err != nil {
		logger.Error(utils.PrepareMsg(err, bundle))
		return mo.Err[*models.DeleteCmdOutputData](err)
	}

	return mo.Ok(cmdOutput)
}

func NewExportCustomAddrRequest(application application.Application, reqCtx utils.RequestCtx) mo.Result[*models.ExportAddressRequest] {
	body, err := utils.GetDataFromRequestBody[models.ExportAddressRequestBody](reqCtx.GetGinCtx()).Get()

	if err != nil {
		return mo.Err[*models.ExportAddressRequest](err)
	}

	return mo.Ok(&models.ExportAddressRequest{
		Body: body,
	})
}

func ExportCustomAddrRequestController(application application.Application, reqCtx utils.RequestCtx, request *models.ExportAddressRequest) mo.Result[*models.ExportAsyncAddrCmdOutoutData] {
	bundle := application.GetBundle()
	logger := utils.NewApplicationBaseLogger(application.GetLogger(), reqCtx.GetIP())

	cmdCtx := services.NewCommandContext(application, reqCtx, logger)
	
	cmdOutput, err := command.ExecuteCommand(cmdCtx, services.NewExportAsyncAddrCmd(*cmdCtx.GetUserId(), request.Body.Fields, request.Body.Email)).Get()
	if err != nil {
		logger.Error(utils.PrepareMsg(err, bundle))
		return mo.Err[*models.ExportAsyncAddrCmdOutoutData](err)
	}

	return mo.Ok(cmdOutput)
}

func NewFilterAddrRequest(application application.Application, reqCtx utils.RequestCtx) mo.Result[*models.FilterAddrQuery] {
	body, err := utils.GetDataFromRequestBody[models.FilterAddrQueryBody](reqCtx.GetGinCtx()).Get()

	if err != nil {
		return mo.Err[*models.FilterAddrQuery](err)
	}

	return mo.Ok(&models.FilterAddrQuery{
		Body: body,
	})
}

func FilterAddrRequestController(application application.Application, reqCtx utils.RequestCtx, request *models.FilterAddrQuery) mo.Result[*models.FilterAddrCmdOutputData] {
	bundle := application.GetBundle()
	logger := utils.NewApplicationBaseLogger(application.GetLogger(), reqCtx.GetIP())

	cmdCtx := services.NewCommandContext(application, reqCtx, logger)
	
	cmdOutput, err := command.ExecuteCommand(cmdCtx, services.NewFilterAddrCmd(*request, *cmdCtx.GetUserId())).Get()
	if err != nil {
		logger.Error(utils.PrepareMsg(err, bundle))
		return mo.Err[*models.FilterAddrCmdOutputData](err)
	}

	return mo.Ok(cmdOutput)
}