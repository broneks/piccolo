# Piccolo




## Steps to setup

1. make a duplicate of `.env.sample` and fill in values.
1. install go version `1.23.2` and ensure your GOPATH is setup
1. run `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
1. run `docker compose up -d` and let the thing work.
1. you can then view logs for the running containers with `docker compose logs -f`
1. App should run on localhost:8001
1. Test your db works by 
   1. first make sure that if you postgres running locally on your mac that you stop it (brew services stop postgresql)
   1. running `docker-compose exec postgres bash`
   1. enter `psql <paste in the DB_LOCALHOST_URL from .env`.
1. run `make migrate_up`

