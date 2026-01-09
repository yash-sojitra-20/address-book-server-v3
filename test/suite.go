package test

import (
	"address-book-server-v3/internal/core/application"
	"address-book-server-v3/internal/core/config"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type TestingSuite struct {
	application.Application
}

func NewTestingSuite(path *string, testdb *gorm.DB) *TestingSuite {
	envPath := "../../test.env"
	if path != nil {
		envPath = *path
	}

	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading test.env file")
	}

	application := application.NewApplicationForTesting(config.NewAppConfig(), testdb)

	return &TestingSuite{Application: application}
}