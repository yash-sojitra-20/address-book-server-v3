package main

import (
	"address-book-server-v3/internal/core/application"
	"address-book-server-v3/internal/core/config"
	"address-book-server-v3/internal/server"
	"log"

	"bitbucket.org/vayana/walt-gin-gonic/utils"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Error loading .env file")
	}

	appConfig := config.NewAppConfig()
	application := application.NewApplication(appConfig)

	server := server.NewServer(application).AddRoutes().Start()

	<-utils.WaitForTermination(server)
}
