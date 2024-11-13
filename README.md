# Piccolo

Private photos manager

### Dev Setup

1. Duplicate `.env.sample` as `.env` and fill all values
2. Install golang-migrate as an external dependency
    - run `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
3. Start the app using Docker
    - run `docker compose up -d`
4. Open [http://localhost:8001/api/health](http://localhost:8001/api/health) to ensure that the api is running
5. Apply all db migrations
    - run `make migrate_up`

### License

GNU GPLv3
