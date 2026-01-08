package utils

import (
	"address-book-server-v3/internal/common/fault"
	"address-book-server-v3/internal/common/types"
	"address-book-server-v3/internal/core/application"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"regexp"
	"runtime"
	"runtime/debug"

	// "runtime"
	// "runtime/debug"
	"strings"
	"unicode"

	"bitbucket.org/vayana/walt-gin-gonic/request"
	wgf "bitbucket.org/vayana/walt-go/fault"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/samber/mo"
)

type RequestCtx interface {
	GetIP() types.Ip
	GetGinCtx() *gin.Context
	GetUserId() *types.UserId
	GetEmail() *string
}

type _RequestCtx struct {
	GinCtx *gin.Context
	IP     types.Ip
	UserId *types.UserId
	Email *string
}

func NewRequestCtx(
	ginCtx *gin.Context,
) RequestCtx {
	ip := ginCtx.ClientIP()
	userId, _ := GetUserId(ginCtx).Get()
	email, _ := GetEmail(ginCtx).Get()
	return &_RequestCtx{
		GinCtx: ginCtx,
		IP:     types.Ip(ip),
		UserId: userId,
		Email: email,
	}
}

func (r *_RequestCtx) GetIP() types.Ip {
	return r.IP
}

func (r *_RequestCtx) GetGinCtx() *gin.Context {
	return r.GinCtx
}

func (r *_RequestCtx) GetUserId() *types.UserId {
	return r.UserId
}

func (r *_RequestCtx) GetEmail() *string {
	return r.Email
}

type RequestPayload[R any] func(application.Application, RequestCtx) mo.Result[*R]

type ApiRequestHandler[T any, R any] func(application.Application, RequestCtx, *R) mo.Result[*T]

func HandleRequest[RQ any, RS any](application application.Application, handler ApiRequestHandler[RS, RQ], payloadBuilder RequestPayload[RQ]) gin.HandlerFunc {
	bundle := application.GetBundle()
	logger := application.GetLogger()

	return func(c *gin.Context) {
		var bodyBytes []byte
		var err error
		if c.Request != nil && c.Request.Body != nil {
			bodyBytes, err = io.ReadAll(c.Request.Body)
			if err != nil {
				logger.Error(fmt.Sprint("Error in request body", err))
			} else {
				c.Request.Body = io.NopCloser(bytes.NewReader(bodyBytes))
			}
		}
		var res mo.Result[*RS]
		defer func() {
			if c.Request != nil && c.Request.Body != nil {
				c.Request.Body.Close() // #nosec G104
			}
			switch exception := recover(); exception {
			case nil:
				if res.IsError() {
					err := res.Error().(wgf.Fault)
					logger.Error(PrepareMsg(err, bundle))
					status := getStatusCode(err.ResponseErrType())
					res := getErrorResponse(err, bundle)
					c.JSON(status, res)
				} else {
					responseData, _ := res.Get()
					if responseData == nil {
						c.JSON(200, map[string]any{
							"data": nil,
						})
					} else {
						c.JSON(200, map[string]any{
							"data": *responseData,
						})
					}
				}
			default:
				f := fault.InternalServerError(fmt.Errorf("error %+v", exception))
				logger.Error(PrepareMsg(f, bundle))
				stack := debug.Stack()
				logger.Error(fmt.Sprintf("Stack trace for exception : %v \n %v", exception, string(stack)))
				if _, file, line, ok := runtime.Caller(1); ok {
					logger.Error(fmt.Sprintf("Recovered from panic in file %s at line %d: %v\n", file, line, exception))
				} else {
					logger.Error(fmt.Sprintln("Recovered from panic but couldn't retrieve file name and line number"))
				}

				// Send Email
				// fullURL := c.Request.URL.String()
				// host := c.Request.Host

				// emailBody := fmt.Sprintf("<br/>Url: %s<br/> Body: %s <br/><br/> Error: %+v <br/> Stack Trace: %+v", host+"/"+fullURL, string(bodyBytes), f.Cause().Error(), string(stack))
				// // Send Email
				// if _, err := application.GetNotificationService().SendNotification(types.EmailTypeInternalServerError, emailBody, string(f.Code()), true, logger).Get(); err != nil {
				// 	logger.Error(fmt.Sprintf("Error sending email: %v", PrepareMsg(err, bundle)))
				// }

				c.AbortWithStatusJSON(500, getErrorResponse(f, bundle))
			}
		}()

		res = (func() mo.Result[*RS] {			
			reqCtx := NewRequestCtx(c)
			req, err := payloadBuilder(application, reqCtx).Get()
			if err != nil {
				return mo.Err[*RS](err)
			}
			return handler(application, reqCtx, req)
		})()

	}
}

func getStatusCode(response wgf.ResponseErrType) int {
	switch response {
	case fault.BadRequest:
		return http.StatusBadRequest
	case fault.Unauthorized:
		return http.StatusUnauthorized
	case fault.Forbidden:
		return http.StatusForbidden
	case fault.NotFound:
		return http.StatusNotFound
	case fault.AlreadyExists:
		return http.StatusConflict
	}
	return http.StatusInternalServerError
}

func prettifyFieldName(field string) string {
	var result []rune
	for i, r := range field {
		if i > 0 && unicode.IsUpper(r) {
			result = append(result, ' ')
		}
		result = append(result, r)
	}
	return string(result)
}

func parseValidatorErrorString(errString string, fieldMessages map[string]string) []string {
	var messages []string

	// Regex to extract key details
	pattern := regexp.MustCompile(`Key: '([^']+)\.([^']+)' Error:Field validation for '([^']+)' failed on the '([^']+)' tag`)
	matches := pattern.FindAllStringSubmatch(errString, -1)

	for _, match := range matches {
		if len(match) == 5 {
			field := match[2] // E.g., Region
			tag := match[4]   // E.g., required, min, regioncode
			key := field + "." + tag

			switch tag {
			case "required":
				// Generic required message
				messages = append(messages, fmt.Sprintf("%s is required.", prettifyFieldName(field)))
			default:
				if msg, ok := fieldMessages[key]; ok {
					messages = append(messages, msg)
				} else {
					messages = append(messages, fmt.Sprintf("Validation failed for field '%s' with rule '%s'", field, tag))
				}
			}
		}
	}

	return messages
}

func RenderErrorResponse(f wgf.Fault, bundle fault.FaultWrapper, fieldMessages map[string]string) wgf.Fault {
	e := parseValidatorErrorString(f.Cause().Error(), fieldMessages)
	return fault.InvalidRequestError(fmt.Errorf("%s", strings.Join(e, ", ")))

}

func getErrorResponse(f wgf.Fault, bundle *i18n.Bundle) map[string]any {
	var errors []string
	for _, cause := range f.Causes() {
		errors = append(errors, cause.Error())
	}
	return map[string]any{
		"errors": map[string]any{
			"message":     f.ToMessageAwareFault(bundle).Message("en"),
			"errorCode":   f.Code().String(),
			"otherErrors": errors,
		},
	}
}

type ApiMiddlewareHandler func(application.Application, *gin.Context) mo.Result[*bool]

func HandleMiddleware(application application.Application, middlewareHandler ApiMiddlewareHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		var res mo.Result[*bool]
		bundle := application.GetBundle()
		logger := application.GetLogger()

		defer func() {
			switch exception := recover(); exception {
			case nil:
				if res.IsError() {
					err := res.Error().(wgf.Fault)
					status := getStatusCode(err.ResponseErrType())
					res := getErrorResponse(err, bundle)
					c.AbortWithStatusJSON(status, res)
					c.Request.Body.Close() // #nosec G104
				} else {
					c.Next()
				}
			default:
				f := fault.InternalServerError(fmt.Errorf("error %+v", exception))
				logger.Error(PrepareMsg(f, bundle))
				stack := debug.Stack()
				formattedStack := strings.ReplaceAll(string(stack), "\n", "\n")
				formattedStack = strings.ReplaceAll(formattedStack, "\t", "\t")
				logger.Error(fmt.Sprintf("Stack trace for exception : %v \n %s\n", exception, formattedStack))
				if _, file, line, ok := runtime.Caller(1); ok {
					logger.Error(fmt.Sprintf("Recovered from panic in file %s at line %d: %v\n", file, line, exception))
				} else {
					logger.Error(fmt.Sprintln("Recovered from panic but couldn't retrieve file name and line number"))
				}

				// // Send Email
				// // emailRepo := repositories.NewEmailRepo(repositories.NewRepoContext(application.GetDb(), logger, bundle))
				// // _, f = application.GetNotificationService().SendNotification(types.EMAIL_TYPE_INTERNAL_SERVER_ERROR, f.Cause().Error(), string(f.Code()), true, logger, emailRepo.ShouldSkipEmailByErrorCode).ToValue()
				// // if f != nil {
				// // 	logger.Error(fmt.Sprintf("Error sending email: %v", PrepareMsg(f, bundle)))
				// // }

				c.AbortWithStatusJSON(500, getErrorResponse(f, bundle))
				c.Request.Body.Close() // #nosec G104
			}
		}()
		res = middlewareHandler(application, c)
	}
}

func TrimStructStrings(v any) {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		// Skip unexported fields
		if !field.CanSet() {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			// Direct string field
			field.SetString(strings.TrimSpace(field.String()))

		case reflect.Ptr:
			// Pointer fields
			if field.IsNil() {
				continue
			}
			elem := field.Elem()
			switch elem.Kind() {
			case reflect.String:
				trimmed := strings.TrimSpace(elem.String())

				// Create a new value of the correct type
				newVal := reflect.New(elem.Type()).Elem()
				newVal.SetString(trimmed)
				field.Set(newVal.Addr())

			default:
				// Recurse for nested struct pointers or user-defined types
				if elem.Kind() == reflect.Struct {
					TrimStructStrings(field.Interface())
				}
			}

		case reflect.Struct:
			// Recurse into nested structs
			TrimStructStrings(field.Addr().Interface())
		}
	}
}

func GetDataFromRequestBody[T any](c *gin.Context) mo.Result[*T] {
	var t T
	if err := c.ShouldBind(&t); err != nil {
		return mo.Err[*T](fault.GetRequestDataError(err))
	}
	TrimStructStrings(&t)
	validate := validator.New()
	validate = registerCustomValidators(validate)
	if err := validate.Struct(t); err != nil {
		return mo.Err[*T](fault.InvalidRequestError(err))
	}
	return mo.Ok(&t)
}

func GetUserId(c *gin.Context) mo.Result[*types.UserId] {
	userIdResult := request.GetValueFromGinContext[types.UserId](c, "userId")
	if userIdResult.IsError() {
		err := userIdResult.Error()
		return mo.Err[*types.UserId](fault.GetUserIdError("userId", err))
	}
	return userIdResult
}


func GetEmail(c *gin.Context) mo.Result[*string] {
	email := request.GetValueFromGinContext[string](c, "email")
	if email.IsError() {
		err := email.Error()
		return mo.Err[*string](fault.GetUserIdError("email", err))
	}
	return email
}