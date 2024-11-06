package pg

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresClient struct {
	*pgxpool.Pool
}

func NewClient(ctx context.Context) (*PostgresClient, error) {
	db, err := pgxpool.New(ctx, os.Getenv("DB_DOCKER_URL"))
	if err != nil {
		return nil, err
	}

	client := &PostgresClient{db}

	return client, nil
}
