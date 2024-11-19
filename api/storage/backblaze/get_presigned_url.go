package backblaze

import (
	"context"
	"fmt"
	"log/slog"
	"time"
)

const EXPIRATION_DURATION = time.Hour * 24
const FILE_NOT_FOUND_IMAGE_URL = "/image-not-found.png"

func (bc *BackblazeClient) GetPresignedUrl(ctx context.Context, filename, userId string) (string, time.Duration) {
	if !bc.doesObjectExist(ctx, filename, userId) {
		slog.Error(fmt.Sprintf("file \"%s\" does not exist in backblaze for user id \"%s\".", filename, userId))
		return FILE_NOT_FOUND_IMAGE_URL, 0
	}

	name := newObjectName(filename, userId)

	obj := bc.bucket.Object(name)

	url, err := obj.AuthURL(ctx, EXPIRATION_DURATION, "")
	if err != nil {
		slog.Error(fmt.Sprintf("failed to get presigned url: %v", err))
		return "", 0
	}

	presignedUrl := url.String()

	return presignedUrl, EXPIRATION_DURATION - (time.Minute * 5)
}
