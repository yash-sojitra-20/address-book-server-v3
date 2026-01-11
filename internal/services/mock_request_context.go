package services

import (
	"address-book-server-v3/internal/common/types"
	"address-book-server-v3/internal/common/utils"

	"github.com/gin-gonic/gin"
)

type MockRequestCtx struct {
	utils.RequestCtx
}

func (m *MockRequestCtx) GetGinCtx() *gin.Context {
	return &gin.Context{}
}

func (m *MockRequestCtx) GetCorrelationId() string {
	return "test-correlation-id"
}

func (m *MockRequestCtx) GetIP() types.Ip {
	return "127.0.0.1"
}

func NewMockRequestCtx() utils.RequestCtx {
	return &MockRequestCtx{}
}