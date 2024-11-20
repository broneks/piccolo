package types

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type ServerLogger interface {
	Debug(msg string, args ...any)
	Error(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
}

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
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	SetForever(ctx context.Context, key string, value interface{}) error
}

type ServerObjectStorage interface {
	GetPresignedUrl(ctx context.Context, filename, userId string) (string, time.Duration)
	Ping(ctx context.Context) error
	UploadFile(ctx context.Context, fileUpload FileUpload) (string, error)
}

type Server struct {
	Logger        ServerLogger
	Mailer        ServerMailer
	DB            ServerDB
	Cache         ServerCache
	ObjectStorage ServerObjectStorage
}
