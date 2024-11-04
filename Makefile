include .env

build:
	go build --o bin/piccolo cmd/piccolo/main.go

run:
	air

migrate_create:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrate_up:
	migrate -database $(DB_LOCALHOST_URL)?sslmode=disable -path db/migrations up

migrate_down:
	migrate -database $(DB_LOCALHOST_URL)?sslmode=disable -path db/migrations down 1
