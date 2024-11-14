package backblaze

import (
	"context"
)

func (b *BackblazeClient) Ping(ctx context.Context) error {
	if _, err := b.client.Bucket(ctx, b.bucket.Name()); err != nil {
		return err
	}

	return nil
}
