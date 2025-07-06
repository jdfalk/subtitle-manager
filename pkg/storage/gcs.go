// file: pkg/storage/gcs.go
// version: 1.0.0
// guid: 67890123-45f0-6789-0123-456789012345

package storage

import (
	"context"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// GCSProvider implements StorageProvider for Google Cloud Storage.
type GCSProvider struct {
	client *storage.Client
	bucket string
}

// NewGCSProvider creates a new Google Cloud Storage provider.
func NewGCSProvider(config StorageConfig) (*GCSProvider, error) {
	if config.GCSBucket == "" {
		return nil, fmt.Errorf("%w: GCS bucket name is required", ErrConfigurationMissing)
	}

	ctx := context.Background()
	var client *storage.Client
	var err error
	if config.GCSCredentials != "" {
		client, err = storage.NewClient(ctx, option.WithCredentialsFile(config.GCSCredentials))
	} else {
		client, err = storage.NewClient(ctx)
	}
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrConnectionFailed, err)
	}

	return &GCSProvider{
		client: client,
		bucket: config.GCSBucket,
	}, nil
}

// Store saves content to Google Cloud Storage.
func (gp *GCSProvider) Store(ctx context.Context, key string, content io.Reader, contentType string) error {
	if key == "" {
		return ErrInvalidKey
	}

	obj := gp.client.Bucket(gp.bucket).Object(key)
	w := obj.NewWriter(ctx)
	if contentType != "" {
		w.ContentType = contentType
	}
	if _, err := io.Copy(w, content); err != nil {
		_ = w.CloseWithError(err)
		return fmt.Errorf("failed to upload object: %w", err)
	}
	return w.Close()
}

// Retrieve reads content from Google Cloud Storage.
func (gp *GCSProvider) Retrieve(ctx context.Context, key string) (io.ReadCloser, error) {
	if key == "" {
		return nil, ErrInvalidKey
	}
	r, err := gp.client.Bucket(gp.bucket).Object(key).NewReader(ctx)
	if err != nil {
		if err == storage.ErrObjectNotExist {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to download object: %w", err)
	}
	return r, nil
}

// Delete removes content from Google Cloud Storage.
func (gp *GCSProvider) Delete(ctx context.Context, key string) error {
	if key == "" {
		return ErrInvalidKey
	}
	err := gp.client.Bucket(gp.bucket).Object(key).Delete(ctx)
	if err != nil {
		if err == storage.ErrObjectNotExist {
			return ErrNotFound
		}
		return fmt.Errorf("failed to delete object: %w", err)
	}
	return nil
}

// Exists checks if content exists in Google Cloud Storage.
func (gp *GCSProvider) Exists(ctx context.Context, key string) (bool, error) {
	if key == "" {
		return false, ErrInvalidKey
	}
	_, err := gp.client.Bucket(gp.bucket).Object(key).Attrs(ctx)
	if err != nil {
		if err == storage.ErrObjectNotExist {
			return false, nil
		}
		return false, fmt.Errorf("failed to check object: %w", err)
	}
	return true, nil
}

// List returns keys matching the given prefix in Google Cloud Storage.
func (gp *GCSProvider) List(ctx context.Context, prefix string) ([]string, error) {
	it := gp.client.Bucket(gp.bucket).Objects(ctx, &storage.Query{Prefix: prefix})
	var keys []string
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to list objects: %w", err)
		}
		keys = append(keys, attrs.Name)
	}
	return keys, nil
}

// GetURL returns a signed URL for the object in Google Cloud Storage.
func (gp *GCSProvider) GetURL(ctx context.Context, key string, expiry time.Duration) (string, error) {
	if key == "" {
		return "", ErrInvalidKey
	}
	if expiry <= 0 {
		expiry = time.Hour
	}
	// For now return a public URL. Signed URLs require service account creds.
	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", gp.bucket, key), nil
}

// Close is a no-op for Google Cloud Storage (connections are managed by the SDK).
func (gp *GCSProvider) Close() error {
	if gp.client != nil {
		return gp.client.Close()
	}
	return nil
}
