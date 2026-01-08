package controllers

import (
	"address-book-server-v3/internal/common/utils"
	"address-book-server-v3/internal/core/application"
	"address-book-server-v3/internal/models"
	"address-book-server-v3/internal/services"

	"bitbucket.org/vayana/walt-go/command"
	"github.com/google/uuid"
	"github.com/samber/mo"
)


func NewRegisterRequest(application application.Application, reqCtx utils.RequestCtx) mo.Result[*models.RegisterRequest] {
	body, err := utils.GetDataFromRequestBody[models.RegisterRequestBody](reqCtx.GetGinCtx()).Get()

	if err != nil {
		return mo.Err[*models.RegisterRequest](err)
	}

	return mo.Ok(&models.RegisterRequest{
		Body: body,
	})
}

func RegisterRequestController(application application.Application, reqCtx utils.RequestCtx, request *models.RegisterRequest) mo.Result[*models.RegisterResponse] {
	bundle := application.GetBundle()
	logger := utils.NewApplicationBaseLogger(application.GetLogger(), reqCtx.GetIP())

	cmdCtx := services.NewCommandContext(application, reqCtx, logger)
	
	cmdOutput, err := command.ExecuteCommand(cmdCtx, services.NewRegisterUserCmd(request.Body.Email, request.Body.Password)).Get()
	if err != nil {
		logger.Error(utils.PrepareMsg(err, bundle))
		return mo.Err[*models.RegisterResponse](err)
	}

	response := models.NewRegisterResponse(uuid.UUID(cmdOutput.Id), cmdOutput.Email)

	return mo.Ok(response)
}

func NewLoginRequest(application application.Application, reqCtx utils.RequestCtx) mo.Result[*models.LoginRequest] {
	body, err := utils.GetDataFromRequestBody[models.LoginRequestBody](reqCtx.GetGinCtx()).Get()

	if err != nil {
		return mo.Err[*models.LoginRequest](err)
	}

	return mo.Ok(&models.LoginRequest{
		Body: body,
	})
}

func LoginRequestController(application application.Application, reqCtx utils.RequestCtx, request *models.LoginRequest) mo.Result[*models.LoginResponse] {
	bundle := application.GetBundle()
	logger := utils.NewApplicationBaseLogger(application.GetLogger(), reqCtx.GetIP())

	cmdCtx := services.NewCommandContext(application, reqCtx, logger)
	
	cmdOutput, err := command.ExecuteCommand(cmdCtx, services.NewLoginUserCmd(request.Body.Email, request.Body.Password)).Get()
	if err != nil {
		logger.Error(utils.PrepareMsg(err, bundle))
		return mo.Err[*models.LoginResponse](err)
	}

	response := models.NewLoginResponse(*cmdOutput.Token)
	return mo.Ok(response)
}