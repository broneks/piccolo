package pg

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
	db *pgxpool.Pool
}

func (pg *postgres) Ping(ctx context.Context) error {
	return pg.db.Ping(ctx)
}

func (pg *postgres) Close() {
	pg.db.Close()
}

var (
	pgInstance *postgres
	pgOnce     sync.Once
)

func Client(ctx context.Context) *postgres {
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, os.Getenv("DB_DOCKER_URL"))
		if err != nil {
			log.Fatalf("unable to create postgres connection pool: %v", err)
			return
		}

		pgInstance = &postgres{db}
	})

	return pgInstance
}
