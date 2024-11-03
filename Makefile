include .env

build:
	go build --o bin/piccolo cmd/piccolo/main.go

run:
	go run cmd/piccolo/main.go

migrate_up:
	migrate -verbose -database $(DB_LOCALHOST_URL)?sslmode=disable -path db/migrations up
