package types

import "github.com/google/uuid"

const (
	DB_HOSTNAME       = "DB_HOSTNAME"
	DB_PORT           = "DB_PORT"
	DB_USERNAME       = "DB_USERNAME"
	DB_PASSWORD       = "DB_PASSWORD"
	DB_NAME           = "DB_NAME"
	DB_URL			  = "DB_URL"

	SMTP_USER         = "SMTP_USER"
	SMTP_PASS         = "SMTP_PASS"
	SMTP_HOST         = "SMTP_HOST"
	SMTP_PORT         = "SMTP_PORT"
	
	SERVER_PORT       = "SERVER_PORT"
	JWT_SECRET        = "JWT_SECRET"
	APP_URL           = "APP_URL"
)

// var SYSTEM_USER_ID uuid.UUID = uuid.MustParse("6a1f0001-19b2-11f0-87f7-1cc10cfb5018")

type UserId uuid.UUID

type AddressId uuid.UUID

type Ip string