package backblaze

import (
	"context"
)

func (bc *BackblazeClient) Ping(ctx context.Context) error {
	if _, err := bc.client.Bucket(ctx, bc.bucket.Name()); err != nil {
		return err
	}

	return nil
}
