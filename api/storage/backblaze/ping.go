package backblaze

import (
	"context"
)

func (b *BackblazeClient) Ping(ctx context.Context) error {
	if _, err := b.client.ListBuckets(ctx); err != nil {
		return err
	}

	return nil
}
