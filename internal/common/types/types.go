package types

import "github.com/google/uuid"

const (
	DB_HOSTNAME = "DB_HOSTNAME"
	DB_PORT     = "DB_PORT"
	DB_USERNAME = "DB_USERNAME"
	DB_PASSWORD = "DB_PASSWORD"
	DB_NAME     = "DB_NAME"
	DB_URL      = "DB_URL"
	DB_TYPE     = "DB_TYPE"

	SMTP_USER = "SMTP_USER"
	SMTP_PASS = "SMTP_PASS"
	SMTP_HOST = "SMTP_HOST"
	SMTP_PORT = "SMTP_PORT"

	APP_PORT   = "APP_PORT"
	SECRET_KEY = "SECRET_KEY"
	APP_URL    = "APP_URL"
	LOG_HOST   = "LOG_HOST"
)

var SYSTEM_USER_ID uuid.UUID = uuid.MustParse("6a1f0001-19b2-11f0-87f7-1cc10cfb5018")

type UserId uuid.UUID

type AddressId uuid.UUID

type Ip string
