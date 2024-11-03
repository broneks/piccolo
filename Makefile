include .env

build:
	go build --o bin/piccolo cmd/piccolo/main.go

run:
	go run cmd/piccolo/main.go

migrate_up:
	migrate -database $(DB_LOCALHOST_URL)?sslmode=disable -path db/migrations up

migrate_down:
	migrate -database $(DB_LOCALHOST_URL)?sslmode=disable -path db/migrations down 1
