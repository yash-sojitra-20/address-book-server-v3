# Load .env file
include .env
export

migrate-up:
	migrate -path database/migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path database/migrations -database "$(DB_URL)" down

migrate-version:
	migrate -path database/migrations -database "$(DB_URL)" version


include test.env

migrate-up-test:
	migrate -path database/migrations -database "$(TEST_DB_URL)" up

migrate-version-test:
	migrate -path database/migrations -database "$(TEST_DB_URL)" version

migrate-down-test:
	migrate -path database/migrations -database "$(TEST_DB_URL)" down