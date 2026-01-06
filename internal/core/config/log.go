package config

import (
	"address-book-server-v3/internal/common/types"
	"io"
	"net"
	"os"

	"bitbucket.org/vayana/walt-go/logger"
	"bitbucket.org/vayana/walt-go/osutil"
	"github.com/samber/mo"
)

type LoggerConfig struct {
	url string
}

func (lc *LoggerConfig) GetUrl() string {
	return lc.url
}

func newLoggerConfig() mo.Result[*LoggerConfig] {

	url, err := osutil.GetEnvVar(types.LOG_HOST).Get()
	if err != nil {
		return mo.Err[*LoggerConfig](err)
	}
	return mo.Ok(&LoggerConfig{
		url: *url,
	})
}

func NewLogger(url string, isProd bool) (*logger.Logger, error) {
	var writer io.Writer
	if isProd {
		var err error
		writer, err = net.Dial("udp", url)
		if err != nil {
			return nil, err
		}
	} else {
		writer = os.Stdout
	}
	logger := logger.NewLogger(writer, "address-book-server-v3", isProd)
	return logger, nil
}
