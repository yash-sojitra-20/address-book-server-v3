DB_URL=mysql://swaraj:swaraj%40123@tcp(localhost:3306)/addressbook_v3

migrate-up:
	migrate -path database/migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path database/migrations -database "$(DB_URL)" down

migrate-version:
	migrate -path database/migrations -database "$(DB_URL)" version
