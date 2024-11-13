package wasabi

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (w *WasabiClient) Ping(ctx context.Context) error {
	if _, err := w.client.ListBuckets(ctx, &s3.ListBucketsInput{}); err != nil {
		return err
	}

	return nil
}
