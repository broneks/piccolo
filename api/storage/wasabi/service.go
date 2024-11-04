package wasabi

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const EXPIRATION_DURATION = time.Hour * 24

func GetPresignedUrl(ctx context.Context, key string) (string, error) {
	var err error

	client, err := newClient(context.Background())
	if err != nil {
		return "", err
	}

	presignClient := s3.NewPresignClient(client)
	presignedUrl, err := presignClient.PresignGetObject(
		context.Background(),
		&s3.GetObjectInput{
			Bucket: aws.String(os.Getenv("WASABI_BUCKET_NAME")),
			Key:    aws.String(key),
		},
		s3.WithPresignExpires(EXPIRATION_DURATION),
	)
	if err != nil {
		log.Println("error: %w", err)
		return "", err
	}

	return presignedUrl.URL, nil
}
