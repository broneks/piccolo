package pg

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresClient struct {
	*pgxpool.Pool
}

func NewClient(ctx context.Context) (*PostgresClient, error) {
	var err error

	db, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("cannot create database connection: %v", err)
	}

	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	client := &PostgresClient{db}

	return client, nil
}
