package wasabi

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const EXPIRATION_DURATION = time.Hour * 24

func (wc *WasabiClient) GetPresignedUrl(ctx context.Context, key string) (string, time.Duration) {
	presignedUrl, err := wc.presigner.PresignGetObject(
		ctx,
		&s3.GetObjectInput{
			Bucket: aws.String(wc.config.bucket),
			Key:    aws.String(key),
		},
		s3.WithPresignExpires(EXPIRATION_DURATION),
	)
	if err != nil {
		log.Println("failed to get presigned url: %w", err)
		return "", 0
	}

	return presignedUrl.URL, EXPIRATION_DURATION
}
