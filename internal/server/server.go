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
	routes.AddRoutes(server.router, application)
	server.router.Use(logmiddleware.GinMiddlewareLogger(server.GetLogger()))
	return server
}

// Start starts the HTTP server.
func (s *Server) Start() *http.Server {
	s.router.Use(logmiddleware.GinMiddlewareLogger(s.GetLogger()))

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
