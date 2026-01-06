package application

import (
	"address-book-server-v3/internal/common/types"
	"address-book-server-v3/internal/core/config"
	"log"

	"bitbucket.org/vayana/walt-go/logger"
	auditlog "bitbucket.org/vayana/walt-gorm.go/audit"
	wgormconfig "bitbucket.org/vayana/walt-gorm.go/config"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gorm.io/gorm"
)

type Application interface {
	GetLogger() *logger.Logger
	GetDb() *gorm.DB
	GetBundle() *i18n.Bundle
	GetConfig() *config.AppConfig
}

type application struct {
	config *config.AppConfig
	logger *logger.Logger
	db     *gorm.DB
	bundle *i18n.Bundle
}

func NewApplication(appConfig *config.AppConfig) Application {
	db, err := wgormconfig.ConnectToDatabase(appConfig.DatabaseConfig).Get()
	if err != nil {
		log.Fatal("dbConn", err)
	}

	logger, err := config.NewLogger(appConfig.LoggerConfig.GetUrl(), false)
	if err != nil {
		log.Fatal(err)
	}

	bundle, err := config.NewFaultWrapper(db, logger)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize auditlog
	auditTables := []string{}
	ignoreFields := []string{"created_at", "updated_at", "created_by", "updated_by"}

	err = auditlog.Initialize(
		db,
		logger,
		auditTables,
		ignoreFields,
		types.SYSTEM_USER_ID,
	)
	if err != nil {
		log.Fatal(err)
	}

	return &application{
		config: appConfig,
		logger: logger,
		db:     db,
		bundle: bundle,
	}
}

func (application *application) GetConfig() *config.AppConfig {
	return application.config
}

func (application *application) GetLogger() *logger.Logger {
	return application.logger
}

func (application *application) GetDb() *gorm.DB {
	return application.db
}

func (application *application) GetBundle() *i18n.Bundle {
	return application.bundle
}
