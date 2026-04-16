package server

import (
	// "address-book-server-v3/internal/common/fault"
	// "address-book-server-v3/internal/common/types"
	"address-book-server-v3/internal/core/application"
	"address-book-server-v3/internal/routes"
	"fmt"
	"log"
	"net/http"
	"time"

	logmiddleware "bitbucket.org/vayana/walt-go/logger/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	application.Application
	router *gin.Engine
}

func NewServer(application application.Application) *Server {
	server := &Server{Application: application, router: gin.New()}

	return server
}

func (server *Server) AddRoutes() *Server {
	application := server.Application
	
	// Add middleware BEFORE defining routes
	server.router.Use(CORSMiddleware(nil))
	server.router.Use(logmiddleware.GinMiddlewareLogger(server.GetLogger()))
	
	routes.AddRoutes(server.router, application)
	return server
}

// Start starts the HTTP server.
func (s *Server) Start() *http.Server {
	// Logger middleware is now added in AddRoutes, so we don't need to add it here again.
	server := &http.Server{
		Addr:              ":" + fmt.Sprint(s.GetConfig().GetPort()),
		Handler:           s.router,
		ReadHeaderTimeout: time.Second * 200,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// f := fault.ConfigError(err)
			log.Fatal("unable to start the server: " + err.Error())
		}
	}()

	return server
}

// CORSMiddleware configures CORS using gin-contrib/cors library
func CORSMiddleware(allowedOrigins []string) gin.HandlerFunc {
	if len(allowedOrigins) == 0 {
		allowedOrigins = []string{"*"}
	}

	return cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Call-ID"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	})
}
