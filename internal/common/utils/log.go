package utils

import (
	"address-book-server-v3/internal/common/types"
	"fmt"
	"runtime"

	"bitbucket.org/vayana/walt-go/fault"
	"bitbucket.org/vayana/walt-go/logger"
	"github.com/google/uuid"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func NewApplicationBaseLogger(
	logger *logger.Logger,
	ip types.Ip,
) *logger.Logger {
	return logger.CorrelationLogger("correlationId", uuid.New().String()).With(map[string]any{
		"ip": ip,
	})
}

func PrepareMsg(err error, b *i18n.Bundle) string {
	pc, file, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc).Name()
	f, ok := err.(fault.Fault)
	if ok {
		return fmt.Sprintf(
			"Error Code: %+v, Error Message: %+v, Other Errors: %+v, Function Name: %s, File Name: %s:%d",
			f.Code(),
			f.ToMessageAwareFault(b).Message("en"),
			f.Causes(),
			fn, file, line,
		)
	}
	return fmt.Sprintf(
		"Error %+v, Function Name: %s, File Name: %s:%d",
		err,
		fn, file, line,
	)

}
