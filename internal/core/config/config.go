package config

import (
	"address-book-server-v3/internal/common/types"
	"log"

	"bitbucket.org/vayana/walt-go/errors"
	"bitbucket.org/vayana/walt-go/fault"
	wgconfig "bitbucket.org/vayana/walt-gorm.go/config"
	wgerrors "bitbucket.org/vayana/walt-gorm.go/errors"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type AppConfig struct {
	*wgconfig.DatabaseConfig
	*ServerConfig
	*LoggerConfig
	*JwtConfig
}

func NewAppConfig() *AppConfig {
	dbConfig, err := wgconfig.NewDBConfig(
		types.DB_HOSTNAME,
		types.DB_PORT,
		types.DB_USERNAME,
		types.DB_PASSWORD,
		types.DB_NAME,
		types.DB_TYPE,
	).Get()
	if err != nil {
		log.Fatal("dbConfig", getErr(err, wgerrors.FaultBundle))
	}

	serverConfig, err := newServerConfig().Get()
	if err != nil {
		log.Fatal("serverConfig", getErr(err, errors.FaultBundle))
	}

	loggerConfig, err := newLoggerConfig().Get()
	if err != nil {
		log.Fatal("loggerConfig", getErr(err, errors.FaultBundle))
	}

	jwtConfig, err := newJwtConfig().Get()
	if err != nil {
		log.Fatal("jwtConfig", getErr(err, errors.FaultBundle))
	}

	log.Println("Config build success")
	return &AppConfig{
		dbConfig,
		serverConfig,
		loggerConfig,
		jwtConfig,
	}
}

func getErr(err error, bundle *i18n.Bundle) string {
	f, ok := err.(fault.Fault)
	if ok {
		return f.ToMessageAwareFault(bundle).Message("en")
	}
	return err.Error()
}
