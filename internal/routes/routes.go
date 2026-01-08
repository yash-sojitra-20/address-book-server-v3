package routes

import (
	"address-book-server-v3/internal/common/utils"
	"address-book-server-v3/internal/controllers"
	"address-book-server-v3/internal/core/application"

	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine, application application.Application) {
	api := router.Group("/api")
	v3 := api.Group("/v3")

	authenticationRoutes := v3.Group("/auth")
	{
		authenticationRoutes.POST("/register", utils.HandleRequest(application, controllers.RegisterRequestController, controllers.NewRegisterRequest))

		authenticationRoutes.POST("/login", utils.HandleRequest(application, controllers.LoginRequestController, controllers.NewLoginRequest))
	}

	// All Private Routes

	addressRoutes := v3.Group("/address")
	{
		addressRoutes.POST("", utils.HandleRequest(application, controllers.CreateAddrRequestController, controllers. NewCreateAddrRequest))

		addressRoutes.GET("", utils.HandleRequest(application, controllers.ListAllAddrRequestController, controllers. NewListAllAddrRequest))
		
		addressRoutes.GET("/:id", utils.HandleRequest(application, controllers.GetByIdRequestController, controllers. NewGetByIdRequest))

		addressRoutes.PUT("/:id", utils.HandleRequest(application, controllers.UpdateAddrRequestController, controllers. NewUpdateAddrRequest))

		addressRoutes.DELETE("/:id", utils.HandleRequest(application, controllers.DeleteAddrRequestController, controllers. NewDeleteAddrRequest))

		addressRoutes.POST("/export", utils.HandleRequest(application, controllers.ExportCustomAddrRequestController, controllers. NewExportCustomAddrRequest))

		addressRoutes.GET("/filter", utils.HandleRequest(application, controllers.FilterAddrRequestController, controllers. NewFilterAddrRequest))
	}
	
}
