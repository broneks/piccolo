package backblaze

import (
	"context"
	"log"
	"os"

	"github.com/Backblaze/blazer/b2"
)

type BackblazeClient struct {
	client *b2.Client
	bucket *b2.Bucket
}

func NewClient(ctx context.Context) (*BackblazeClient, error) {
	id := os.Getenv("B2_APP_KEY_ID")
	key := os.Getenv("B2_APP_KEY")
	bucketName := os.Getenv("B2_BUCKET_NAME")

	b2Client, err := b2.NewClient(ctx, id, key)
	if err != nil {
		log.Fatalf("cannot create backblaze connection: %v", err)
	}

	bucket, err := b2Client.Bucket(ctx, bucketName)
	if err != nil {
		log.Fatalf("cannot find backblaze bucket: %v", err)
	}

	client := &BackblazeClient{
		client: b2Client,
		bucket: bucket,
	}

	return client, nil
}
