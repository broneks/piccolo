package backblaze

import (
	"context"
	"fmt"
	"log/slog"
	"time"
)

const EXPIRATION_DURATION = time.Hour * 24

func (b *BackblazeClient) GetPresignedUrl(ctx context.Context, filename, userId string) (string, time.Duration) {
	name := newObjectName(filename, userId)

	url, err := b.bucket.Object(name).AuthURL(ctx, EXPIRATION_DURATION, "")
	if err != nil {
		slog.Error(fmt.Sprintf("failed to get presigned url: %v", err))
		return "", 0
	}

	presignedUrl := url.String()

	return presignedUrl, EXPIRATION_DURATION - (time.Minute * 5)
}
