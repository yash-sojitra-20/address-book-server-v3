package services

import (
	"address-book-server-v3/internal/common/utils"
	"address-book-server-v3/internal/core/application"

	"bitbucket.org/vayana/walt-go/logger"
)

type CommandContext interface {
	application.Application
	utils.RequestCtx

	IsCmdContext()
	GetLogger() *logger.Logger
}
type _CommandContext struct {
	application.Application
	utils.RequestCtx
	logger *logger.Logger
}

func (*_CommandContext) IsCmdContext() {}

func (c *_CommandContext) GetLogger() *logger.Logger {
	return c.logger
}

func NewCommandContext(
	application application.Application,
	reqCtx utils.RequestCtx,
	correlationLogger *logger.Logger,
) CommandContext {
	return &_CommandContext{application, reqCtx, correlationLogger}
}