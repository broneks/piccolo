package wasabi

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func newClient(ctx context.Context) (*s3.Client, error) {
	accessKeyID := os.Getenv("WASABI_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("WASABI_SECRET_ACCESS_KEY")

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
		log.Println("Configuration error:", err)
		return nil, err
	}

	return s3.NewFromConfig(cfg), nil
}
