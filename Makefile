# Load .env file
include .env
export

migrate-up:
	migrate -path database/migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path database/migrations -database "$(DB_URL)" down

migrate-version:
	migrate -path database/migrations -database "$(DB_URL)" version
