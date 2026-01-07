package routes

import (
	"address-book-server-v3/internal/core/application"
	
	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine, application application.Application) {
// 	v1 := router.Group("/v1")

// 	authenticationRoutes := v1.Group("/auth")
// 	{
// 		authenticationRoutes.POST("/register", utils.HandleRequest(application, controllers.AuthenticatedUserController, controllers.NewAuthenticateUserRequest))
// 	}

// 	// All Private Routes

// 	// v1.Use(utils.HandleMiddleware(application, middlewares.JWTVerificationMiddlewareByUserId))
// 	// addressRoutes := v1.Group("/address")
// 	// {
// 	// 	addressRoutes.POST("", utils.HandleRequest(application, controllers.AuthenticatedUserController, controllers.NewAuthenticateUserRequest, nil))
// 	// }
	
}
