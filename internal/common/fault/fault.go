package fault

import (
	"address-book-server-v3/internal/common/types"
	"encoding/json"

	"bitbucket.org/vayana/walt-go/fault"
	// "github.com/google/uuid"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// Error Components
const (
	ErrService     fault.ErrComponent = "service"
	ErrRepo        fault.ErrComponent = "repository"
	ErrController  fault.ErrComponent = "controller"
	ErrApplication fault.ErrComponent = "application"
)

// Error Response Types
const (
	BadRequest     fault.ResponseErrType = "BadRequest"
	Forbidden      fault.ResponseErrType = "Forbidden"
	NotFound       fault.ResponseErrType = "NotFound"
	AlreadyExists  fault.ResponseErrType = "AlreadyExists"
	InternalServer fault.ResponseErrType = "InternalServerError"
	Unauthorized   fault.ResponseErrType = "Unauthorized"
)

// Error Codes
const (
	// configuration error codes
	ErrAppConfigErr               fault.ErrorCode = "CNF0000000000"
	ErrInternalServerError        fault.ErrorCode = "CNF0000000010"
	ErrDatabaseInternalError      fault.ErrorCode = "CNF0000000020"
	ErrFailedParsingRSAPublicKey  fault.ErrorCode = "CNF0000000070"
	ErrFailedParsingRSAPrivateKey fault.ErrorCode = "CNF0000000080"

	// request error codes
	ErrFailedToExtractDataFromRequest fault.ErrorCode = "REQ0000000000"
	ErrGetUserIdFromGinCtx            fault.ErrorCode = "REQ0000000020"
	ErrAuthTokenNotFound              fault.ErrorCode = "REQ0000000040" // #nosec G101
	ErrInvalidAuthToken               fault.ErrorCode = "REQ0000000050" // #nosec G101
	ErrInvalidRequestBody             fault.ErrorCode = "REQ0000000060"
	ErrGeneratePostRequest            fault.ErrorCode = "REQ0000000070"
	ErrExecutingRequest               fault.ErrorCode = "REQ0000000090"
	ErrReadingRespBody                fault.ErrorCode = "REQ0000000100"

	//Token Error
	ErrGenerateTokenFailed fault.ErrorCode = "GTK0000000000"

	//Otp Error
	ErrOtpRecordNotFound         fault.ErrorCode = "OTP0000000000"
	ErrInvalidOtpType            fault.ErrorCode = "OTP0000000030"
	ErrApplicationStatusMismatch fault.ErrorCode = "OTP0000000040"
	ErrOtpEmpty                  fault.ErrorCode = "OTP0000000050"

	// db error codes
	ErrRecordNotFound fault.ErrorCode = "REPO0000000000"

	// user error codes
	ErrUserNotFound              fault.ErrorCode = "USR0000000000"
	ErrUserExistWithEmailAlready fault.ErrorCode = "USR0000000010"
	ErrUserPermissionNotFound    fault.ErrorCode = "USR0000000030"
	ErrUserRoleNil               fault.ErrorCode = "USR0000000040"
	ErrInvalidPassword           fault.ErrorCode = "USR0000000050"

	//role error codes
	ErrRoleNotFound fault.ErrorCode = "ROLE000000000"

	// Felix error codes
	ErrSendNotification fault.ErrorCode = "FELIX00000000"

	// Application common error codes
	ErrMarshal fault.ErrorCode = "COMMON0000000"

	// Crypto error codes
	ErrEncryption fault.ErrorCode = "CRYPTO0000000"
	ErrDecryption fault.ErrorCode = "CRYPTO0000010"

	//permission error codes
	ErrPermissionNotFound fault.ErrorCode = "PMS0000000000"

	//role_permission error codes
	ErrRolePermissionNotFound fault.ErrorCode = "RP00000000000"
)

type FaultWrapper struct {
	BasicFault  fault.BasicFaultsCache
	FaultBundle i18n.Bundle
}

var localBasicFaultsCache = fault.NewBasicFaultCache(initLocalBasicFaults())

func initLocalBasicFaults() map[fault.ErrorCode]fault.BasicFault {
	faults := map[fault.ErrorCode]fault.BasicFault{}

	// Configuration Errors
	faults[ErrAppConfigErr] = fault.NewBasicFault(ErrAppConfigErr).SetComponent(ErrApplication).SetResponseType(InternalServer)
	faults[ErrInternalServerError] = fault.NewBasicFault(ErrInternalServerError).SetComponent(ErrApplication).SetResponseType(InternalServer)
	faults[ErrDatabaseInternalError] = fault.NewBasicFault(ErrDatabaseInternalError).SetComponent(ErrRepo).SetResponseType(InternalServer)
	faults[ErrFailedParsingRSAPublicKey] = fault.NewBasicFault(ErrFailedParsingRSAPublicKey).SetComponent(ErrApplication).SetResponseType(InternalServer)
	faults[ErrFailedParsingRSAPrivateKey] = fault.NewBasicFault(ErrFailedParsingRSAPrivateKey).SetComponent(ErrApplication).SetResponseType(InternalServer)

	// Request Errors
	faults[ErrFailedToExtractDataFromRequest] = fault.NewBasicFault(ErrFailedToExtractDataFromRequest).SetComponent(ErrController).SetResponseType(BadRequest)
	faults[ErrGetUserIdFromGinCtx] = fault.NewBasicFault(ErrGetUserIdFromGinCtx).SetComponent(ErrController).SetResponseType(BadRequest)
	faults[ErrAuthTokenNotFound] = fault.NewBasicFault(ErrAuthTokenNotFound).SetComponent(ErrController).SetResponseType(Unauthorized)
	faults[ErrInvalidAuthToken] = fault.NewBasicFault(ErrInvalidAuthToken).SetComponent(ErrController).SetResponseType(Unauthorized)
	faults[ErrInvalidRequestBody] = fault.NewBasicFault(ErrInvalidRequestBody).SetComponent(ErrController).SetResponseType(BadRequest)
	faults[ErrGeneratePostRequest] = fault.NewBasicFault(ErrGeneratePostRequest).SetComponent(ErrController).SetResponseType(InternalServer)
	faults[ErrExecutingRequest] = fault.NewBasicFault(ErrExecutingRequest).SetComponent(ErrController).SetResponseType(InternalServer)
	faults[ErrReadingRespBody] = fault.NewBasicFault(ErrReadingRespBody).SetComponent(ErrController).SetResponseType(InternalServer)

	// Token Errors
	faults[ErrGenerateTokenFailed] = fault.NewBasicFault(ErrGenerateTokenFailed).SetComponent(ErrService).SetResponseType(InternalServer)

	// OTP Errors
	faults[ErrOtpRecordNotFound] = fault.NewBasicFault(ErrOtpRecordNotFound).SetComponent(ErrRepo).SetResponseType(NotFound)
	faults[ErrInvalidOtpType] = fault.NewBasicFault(ErrInvalidOtpType).SetComponent(ErrService).SetResponseType(BadRequest)
	faults[ErrApplicationStatusMismatch] = fault.NewBasicFault(ErrApplicationStatusMismatch).SetComponent(ErrService).SetResponseType(BadRequest)
	faults[ErrOtpEmpty] = fault.NewBasicFault(ErrOtpEmpty).SetComponent(ErrService).SetResponseType(BadRequest)

	// DB Errors
	faults[ErrRecordNotFound] = fault.NewBasicFault(ErrRecordNotFound).SetComponent(ErrRepo).SetResponseType(NotFound)

	// User Errors
	faults[ErrUserNotFound] = fault.NewBasicFault(ErrUserNotFound).SetComponent(ErrRepo).SetResponseType(NotFound)
	faults[ErrUserExistWithEmailAlready] = fault.NewBasicFault(ErrUserExistWithEmailAlready).SetComponent(ErrRepo).SetResponseType(AlreadyExists)
	faults[ErrUserPermissionNotFound] = fault.NewBasicFault(ErrUserPermissionNotFound).SetComponent(ErrService).SetResponseType(Forbidden)
	faults[ErrUserRoleNil] = fault.NewBasicFault(ErrUserRoleNil).SetComponent(ErrService).SetResponseType(Forbidden)
	faults[ErrInvalidPassword] = fault.NewBasicFault(ErrInvalidPassword).SetComponent(ErrService).SetResponseType(Unauthorized)

	// Role Errors
	faults[ErrRoleNotFound] = fault.NewBasicFault(ErrRoleNotFound).SetComponent(ErrRepo).SetResponseType(NotFound)

	// Felix Errors
	faults[ErrSendNotification] = fault.NewBasicFault(ErrSendNotification).SetComponent(ErrService).SetResponseType(InternalServer)

	// Common Errors
	faults[ErrMarshal] = fault.NewBasicFault(ErrMarshal).SetComponent(ErrApplication).SetResponseType(InternalServer)

	// Crypto Errors
	faults[ErrEncryption] = fault.NewBasicFault(ErrEncryption).SetComponent(ErrService).SetResponseType(InternalServer)
	faults[ErrDecryption] = fault.NewBasicFault(ErrDecryption).SetComponent(ErrService).SetResponseType(InternalServer)

	// Permission Errors
	faults[ErrPermissionNotFound] = fault.NewBasicFault(ErrPermissionNotFound).SetComponent(ErrRepo).SetResponseType(NotFound)

	// Role-Permission Errors
	faults[ErrRolePermissionNotFound] = fault.NewBasicFault(ErrRolePermissionNotFound).SetComponent(ErrRepo).SetResponseType(NotFound)

	return faults
}

// Generic Errors
var InternalServerError = _InternalServerError(&localBasicFaultsCache)
var ConfigError = _ConfigError(&localBasicFaultsCache)
var DBError = _DBError(&localBasicFaultsCache)
var RecordNotFound = _RecordNotFound(&localBasicFaultsCache)
var GetRequestDataError = _GetRequestDataError(&localBasicFaultsCache)
var GetUserIdError = _GetUserIdError(&localBasicFaultsCache)
var GeneratePostRequestError = _GeneratePostRequestError(&localBasicFaultsCache)
var ReadingResponseBodyError = _ReadingResponseBodyError(&localBasicFaultsCache)
var ExecutingRequestError = _ExecutingRequestError(&localBasicFaultsCache)
var AuthTokenNotFoundError = _AuthTokenNotFoundError(&localBasicFaultsCache)
var AuthTokenInvalidError = _AuthTokenInvalidError(&localBasicFaultsCache)
var InvalidRequestError = _InvalidRequestError(&localBasicFaultsCache)
var FailedTokenGeneration = _FailedTokenGeneration(&localBasicFaultsCache)

var InvalidPassword = _InvalidPassword(&localBasicFaultsCache)

// Specific Errors
var UserNotFound = _UserNotFound(&localBasicFaultsCache)
var UserExistWithEmailAlready = _UserExistWithEmailAlready(&localBasicFaultsCache)


// Factory Functions

func _InternalServerError(cache *fault.BasicFaultsCache) func(error) fault.Fault {
	return func(cause error) fault.Fault {
		return cache.GetBasicFault(ErrInternalServerError).ToFault(nil, cause)
	}
}

func _ConfigError(cache *fault.BasicFaultsCache) func(error) fault.Fault {
	return func(cause error) fault.Fault {
		return cache.GetBasicFault(ErrAppConfigErr).ToFault(nil, cause)
	}
}

func _DBError(cache *fault.BasicFaultsCache) func(error) fault.Fault {
	return func(cause error) fault.Fault {
		return cache.GetBasicFault(ErrDatabaseInternalError).ToFault(nil, cause)
	}
}

func _RecordNotFound(cache *fault.BasicFaultsCache) func(map[string]any, error) fault.Fault {
	return func(parameters map[string]any, cause error) fault.Fault {
		b, _ := json.Marshal(parameters)
		data := map[string]any{"params": string(b)}
		return cache.GetBasicFault(ErrRecordNotFound).ToFault(data, cause)
	}
}

func _GetRequestDataError(cache *fault.BasicFaultsCache) func(error) fault.Fault {
	return func(cause error) fault.Fault {
		return cache.GetBasicFault(ErrFailedToExtractDataFromRequest).ToFault(nil, cause)
	}
}

func _GetUserIdError(cache *fault.BasicFaultsCache) func(string, error) fault.Fault {
	return func(id string, cause error) fault.Fault {
		data := map[string]any{"name": id}
		return cache.GetBasicFault(ErrGetUserIdFromGinCtx).ToFault(data, cause)
	}
}

func _GeneratePostRequestError(cache *fault.BasicFaultsCache) func(error) fault.Fault {
	return func(cause error) fault.Fault {
		return cache.GetBasicFault(ErrGeneratePostRequest).ToFault(nil, cause)
	}
}

func _ReadingResponseBodyError(cache *fault.BasicFaultsCache) func(error) fault.Fault {
	return func(cause error) fault.Fault {
		return cache.GetBasicFault(ErrReadingRespBody).ToFault(nil, cause)
	}
}

func _ExecutingRequestError(cache *fault.BasicFaultsCache) func(error) fault.Fault {
	return func(cause error) fault.Fault {
		return cache.GetBasicFault(ErrExecutingRequest).ToFault(nil, cause)
	}
}

func _AuthTokenNotFoundError(cache *fault.BasicFaultsCache) func() fault.Fault {
	return func() fault.Fault {
		return cache.GetBasicFault(ErrAuthTokenNotFound).ToFault(nil, nil)
	}
}

func _AuthTokenInvalidError(cache *fault.BasicFaultsCache) func(error) fault.Fault {
	return func(cause error) fault.Fault {
		return cache.GetBasicFault(ErrInvalidAuthToken).ToFault(nil, cause)
	}
}

func _InvalidRequestError(cache *fault.BasicFaultsCache) func(error) fault.Fault {
	return func(cause error) fault.Fault {
		return cache.GetBasicFault(ErrInvalidRequestBody).ToFault(nil, cause)
	}
}


func _FailedTokenGeneration(cache *fault.BasicFaultsCache) func(error) fault.Fault {
	return func(cause error) fault.Fault {
		return cache.GetBasicFault(ErrGenerateTokenFailed).ToFault(nil, cause)
	}
}




func _InvalidPassword(cache *fault.BasicFaultsCache) func(error) fault.Fault {
	return func(cause error) fault.Fault {
		return cache.GetBasicFault(ErrInvalidPassword).ToFault(nil, cause)
	}
}

func _UserNotFound(cache *fault.BasicFaultsCache) func(error, types.UserId) fault.Fault {
	return func(cause error, id types.UserId) fault.Fault {
		data := map[string]any{"user_id": id}
		return cache.GetBasicFault(ErrUserNotFound).ToFault(data, cause)
	}
}

func _UserExistWithEmailAlready(cache *fault.BasicFaultsCache) func(string, error) fault.Fault {
	return func(email string, cause error) fault.Fault {
		data := map[string]any{"email": email}
		return cache.GetBasicFault(ErrUserExistWithEmailAlready).ToFault(data, cause)
	}
}
