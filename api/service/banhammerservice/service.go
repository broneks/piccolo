package banhammerservice

import (
	"log/slog"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type BanHammerService struct {
	rdb *redis.Client
}

const (
	maxAttempts      = 10
	banDuration      = 1 * time.Hour
	trackingTime     = 10 * time.Minute
	banKeyPrefix     = "ban-hammer:ban:"
	attemptKeyPrefix = "ban-hammer:attempt:"
)

func New() *BanHammerService {
	opts, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		slog.Error("cannot create redis connection", "err", err)
	}

	rdb := redis.NewClient(opts)

	return &BanHammerService{rdb}
}
