package controllers

import (
	"address-book-server-v3/internal/common/utils"
	"address-book-server-v3/internal/core/application"
	"address-book-server-v3/internal/core/middlewares"

	"address-book-server-v3/test"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type ControllerTestSuite struct {
	suite.Suite
	app    application.Application
	router *gin.Engine
}

// func (s *ControllerTestSuite) SetupSuite() {

// }

func (s *ControllerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)

	// Create test app
	s.app = test.NewTestingSuite(nil, nil).Application

	// Initialize router
	s.router = gin.New()

	// Register routes
	AddTestRoutes(s.router, s.app)

	// Clean DB
	s.cleanup()
}

func (s *ControllerTestSuite) TearDownSuite() {
	s.cleanup()
}

func (s *ControllerTestSuite) cleanup() {
	s.app.GetDb().Exec("DELETE FROM users")
	s.app.GetDb().Exec("DELETE FROM addresses")
}

func TestControllerSuite(t *testing.T) {
	suite.Run(t, new(ControllerTestSuite))
}

func AddTestRoutes(router *gin.Engine, application application.Application) {
	api := router.Group("/api")
	v3 := api.Group("/v3")

	authenticationRoutes := v3.Group("/auth")
	{
		authenticationRoutes.POST("/register", utils.HandleRequest(application, RegisterRequestController, NewRegisterRequest))

		authenticationRoutes.POST("/login", utils.HandleRequest(application, LoginRequestController, NewLoginRequest))
	}

	// All Private Routes

	v3.Use(utils.HandleMiddleware(application, middlewares.AuthMiddleware))

	addressRoutes := v3.Group("/addresses")
	{
		addressRoutes.POST("", utils.HandleRequest(application, CreateAddrRequestController, NewCreateAddrRequest))

		addressRoutes.GET("", utils.HandleRequest(application, ListAllAddrRequestController, NewListAllAddrRequest))

		addressRoutes.GET("/:id", utils.HandleRequest(application, GetByIdRequestController, NewGetByIdRequest))

		addressRoutes.PUT("/:id", utils.HandleRequest(application, UpdateAddrRequestController, NewUpdateAddrRequest))

		addressRoutes.DELETE("/:id", utils.HandleRequest(application, DeleteAddrRequestController, NewDeleteAddrRequest))

		addressRoutes.POST("/export", utils.HandleRequest(application, ExportCustomAddrRequestController, NewExportCustomAddrRequest))

		addressRoutes.GET("/filter", utils.HandleRequest(application, FilterAddrRequestController, NewFilterAddrRequest))
	}
}
