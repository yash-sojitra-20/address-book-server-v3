package middlewares

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORSMiddleware configures CORS using gin-contrib/cors library
func CORSMiddleware(allowedOrigins []string) gin.HandlerFunc {
	if len(allowedOrigins) == 0 {
		allowedOrigins = []string{"*"}
	}

	return cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	})
}