package wasabi

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type wasabiConfig struct {
	bucket string
	region string
}

type WasabiClient struct {
	config    *wasabiConfig
	client    *s3.Client
	presigner *s3.PresignClient
	uploader  *manager.Uploader
}

func NewClient(ctx context.Context) (*WasabiClient, error) {
	accessKeyID := os.Getenv("WASABI_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("WASABI_SECRET_ACCESS_KEY")
	bucketName := os.Getenv("WASABI_BUCKET_NAME")
	bucketRegion := os.Getenv("WASABI_BUCKET_REGION")

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(
				func(service, region string, options ...interface{}) (aws.Endpoint, error) {
					if service == s3.ServiceID {
						return aws.Endpoint{
							URL:           "https://s3.wasabisys.com",
							SigningRegion: "us-east-1",
						}, nil
					}
					return aws.Endpoint{}, &aws.EndpointNotFoundError{}
				},
			),
		),
	)
	if err != nil {
		log.Println("Configuration error:", err.Error())
		panic(err)
	}

	s3Client := s3.NewFromConfig(cfg)

	client := &WasabiClient{
		config:    &wasabiConfig{bucket: bucketName, region: bucketRegion},
		client:    s3Client,
		presigner: s3.NewPresignClient(s3Client),
		uploader:  manager.NewUploader(s3Client),
	}

	return client, nil
}
