package config

import (
	"address-book-server-v3/internal/common/types"

	"bitbucket.org/vayana/walt-go/osutil"
	"github.com/samber/mo"
)

type ServerConfig struct {
	port   int
	appUrl string
}

func newServerConfig() mo.Result[*ServerConfig] {
	port, err := osutil.GetIntEnvVar(types.APP_PORT).Get()
	if err != nil {
		return mo.Err[*ServerConfig](err)
	}

	appUrl, err := osutil.GetEnvVar(types.APP_URL).Get()
	if err != nil {
		return mo.Err[*ServerConfig](err)
	}

	return mo.Ok(&ServerConfig{
		port:   int(*port),
		appUrl: *appUrl,
	})
}

func (config *ServerConfig) GetPort() int {
	return config.port
}

func (config *ServerConfig) GetAppUrl() string {
	return config.appUrl
}
