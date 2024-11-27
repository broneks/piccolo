include .env

build:
	go build --o bin/piccolo cmd/piccolo

docker_build_local:
	docker build --build-arg ENV=local -t piccolo-app .

lint:
	go vet ./...

start_watch:
	air

start:
	./bin/piccolo

pg_schema:
	docker compose exec postgres pg_dump -U $(DB_USER) --exclude-table-data=schema_migrations --schema-only $(DB_NAME) > db/schema.sql

migrate_create:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrate_up:
	migrate -database $(DATABASE_URL) -path db/migrations up

migrate_down:
	migrate -database $(DATABASE_URL) -path db/migrations down 1
