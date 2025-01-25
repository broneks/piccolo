package types

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type ServerMailer interface {
	SendResetPassword(ctx context.Context, email, token string) error
}

// TODO: remove pg references - use database/sql?
type ServerDB interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Ping(ctx context.Context) error
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
}

type ServerCache interface {
	Get(ctx context.Context, key string) (string, error)
	Ping(ctx context.Context) error
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	AddListItems(ctx context.Context, key string, value ...any) error
	IsListItem(ctx context.Context, key string, value any) (bool, error)
}

type ServerObjectStorage interface {
	GetPresignedUrl(ctx context.Context, filename, userId string) (string, time.Duration)
	Ping(ctx context.Context) error
	UploadFile(ctx context.Context, fileUpload FileUpload) (string, error)
	DeleteFile(ctx context.Context, fileName, userId string) error
}

type Server struct {
	Mailer        ServerMailer
	DB            ServerDB
	Cache         ServerCache
	ObjectStorage ServerObjectStorage
}
