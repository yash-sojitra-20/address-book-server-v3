package config

import (
	"address-book-server-v3/internal/common/types"

	"bitbucket.org/vayana/walt-go/osutil"
	"github.com/samber/mo"
)

type SMTPConfig struct {
	SMTP_USER string
	SMTP_PASS string
	SMTP_HOST string
	SMTP_PORT string
}

func newSMTPConfig() mo.Result[*SMTPConfig] {
	smtp_user, err := osutil.GetEnvVar(types.SMTP_USER).Get()
	if err != nil {
		return mo.Err[*SMTPConfig](err)
	}
	smtp_pass, err := osutil.GetEnvVar(types.SMTP_PASS).Get()
	if err != nil {
		return mo.Err[*SMTPConfig](err)
	}
	smtp_host, err := osutil.GetEnvVar(types.SMTP_HOST).Get()
	if err != nil {
		return mo.Err[*SMTPConfig](err)
	}
	smtp_port, err := osutil.GetEnvVar(types.SMTP_PORT).Get()
	if err != nil {
		return mo.Err[*SMTPConfig](err)
	}

	return mo.Ok(&SMTPConfig{
		SMTP_USER: *smtp_user,
		SMTP_PASS: *smtp_pass,
		SMTP_HOST: *smtp_host,
		SMTP_PORT: *smtp_port,
	})
}

func (config *SMTPConfig) GetSMTPUser() string {
	return config.SMTP_USER
}

func (config *SMTPConfig) GetSMTPPass() string {
	return config.SMTP_PASS
}

func (config *SMTPConfig) GetSMTPHost() string {
	return config.SMTP_HOST
}

func (config *SMTPConfig) GetSMTPPort() string {
	return config.SMTP_PORT
}

