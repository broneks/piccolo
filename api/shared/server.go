package shared

import (
	"log/slog"
	"piccolo/api/storage/pg"
	"piccolo/api/storage/redis"
	"piccolo/api/storage/wasabi"
)

// TODO: Make these interfaces?
type Server struct {
	Logger *slog.Logger
	DB     *pg.PostgresClient
	Redis  *redis.RedisClient
	Wasabi *wasabi.WasabiClient
}
