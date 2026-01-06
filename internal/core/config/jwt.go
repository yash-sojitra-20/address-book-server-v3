package config

import (
	"address-book-server-v3/internal/common/types"

	"bitbucket.org/vayana/walt-go/osutil"
	"github.com/samber/mo"
)

type JwtConfig struct {
	secretKey string
}

func newJwtConfig() mo.Result[*JwtConfig] {

	secretKey, err := osutil.GetEnvVar(types.SECRET_KEY).Get()
	if err != nil {
		return mo.Err[*JwtConfig](err)
	}

	config := JwtConfig{
		secretKey: *secretKey,
	}
	return mo.Ok(&config)
}

func (config *JwtConfig) GetSecretKey() string {
	return config.secretKey
}
