include .env

build:
	go build --o bin/piccolo cmd/piccolo/main.go

lint:
	go vet ./...

run:
	air

pg_schema:
	docker compose exec postgres pg_dump -U $(DB_USER) --exclude-table-data=schema_migrations --schema-only $(DB_NAME) > db/schema.sql

migrate_create:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrate_up:
	migrate -database $(DB_LOCALHOST_URL)?sslmode=disable -path db/migrations up

migrate_down:
	migrate -database $(DB_LOCALHOST_URL)?sslmode=disable -path db/migrations down 1
