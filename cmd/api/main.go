package main

import (
	"address-book-server-v3/internal/core/application"
	"address-book-server-v3/internal/core/config"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Error loading .env file")
	}

	appConfig := config.NewAppConfig()
	application := application.NewApplication(appConfig)

	fmt.Println("application:", application)

}
