// file: pkg/storage/azure.go
// version: 1.0.0
// guid: 56789012-34ef-5678-9012-345678901234

package storage

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/bloberror"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
)

// AzureProvider implements StorageProvider for Azure Blob Storage.
type AzureProvider struct {
	client    *azblob.Client
	container string
}

// NewAzureProvider creates a new Azure Blob Storage provider.
func NewAzureProvider(config StorageConfig) (*AzureProvider, error) {
	if config.AzureAccount == "" || config.AzureKey == "" || config.AzureContainer == "" {
		return nil, fmt.Errorf("%w: Azure account, key, and container are required", ErrConfigurationMissing)
	}

	cred, err := azblob.NewSharedKeyCredential(config.AzureAccount, config.AzureKey)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrConnectionFailed, err)
	}

	serviceURL := fmt.Sprintf("https://%s.blob.core.windows.net/", config.AzureAccount)
	client, err := azblob.NewClientWithSharedKeyCredential(serviceURL, cred, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrConnectionFailed, err)
	}

	return &AzureProvider{
		client:    client,
		container: config.AzureContainer,
	}, nil
}

// Store saves content to Azure Blob Storage.
func (ap *AzureProvider) Store(ctx context.Context, key string, content io.Reader, contentType string) error {
	if key == "" {
		return ErrInvalidKey
	}

	if contentType == "" {
		contentType = "application/octet-stream"
	}

	c := ap.client.ServiceClient().NewContainerClient(ap.container).NewBlockBlobClient(key)
	_, err := c.UploadStream(ctx, content, &azblob.UploadStreamOptions{
		HTTPHeaders: &blob.HTTPHeaders{BlobContentType: &contentType},
	})
	if err != nil {
		return fmt.Errorf("failed to upload blob: %w", err)
	}
	return nil
}

// Retrieve reads content from Azure Blob Storage.
func (ap *AzureProvider) Retrieve(ctx context.Context, key string) (io.ReadCloser, error) {
	if key == "" {
		return nil, ErrInvalidKey
	}
	c := ap.client.ServiceClient().NewContainerClient(ap.container).NewBlobClient(key)
	resp, err := c.DownloadStream(ctx, nil)
	if err != nil {
		if bloberror.HasCode(err, bloberror.BlobNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to download blob: %w", err)
	}
	return resp.Body, nil
}

// Delete removes content from Azure Blob Storage.
func (ap *AzureProvider) Delete(ctx context.Context, key string) error {
	if key == "" {
		return ErrInvalidKey
	}
	_, err := ap.client.DeleteBlob(ctx, ap.container, key, nil)
	if err != nil {
		return fmt.Errorf("failed to delete blob: %w", err)
	}
	return nil
}

// Exists checks if content exists in Azure Blob Storage.
func (ap *AzureProvider) Exists(ctx context.Context, key string) (bool, error) {
	if key == "" {
		return false, ErrInvalidKey
	}
	c := ap.client.ServiceClient().NewContainerClient(ap.container).NewBlobClient(key)
	_, err := c.GetProperties(ctx, nil)
	if err != nil {
		if bloberror.HasCode(err, bloberror.BlobNotFound) {
			return false, nil
		}
		return false, fmt.Errorf("failed to check blob: %w", err)
	}
	return true, nil
}

// List returns keys matching the given prefix in Azure Blob Storage.
func (ap *AzureProvider) List(ctx context.Context, prefix string) ([]string, error) {
	pager := ap.client.NewListBlobsFlatPager(ap.container, &azblob.ListBlobsFlatOptions{Prefix: &prefix})
	var keys []string
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to list blobs: %w", err)
		}
		for _, item := range page.Segment.BlobItems {
			if item.Name != nil {
				keys = append(keys, *item.Name)
			}
		}
	}
	return keys, nil
}

// GetURL returns a SAS URL for the blob in Azure.
func (ap *AzureProvider) GetURL(ctx context.Context, key string, expiry time.Duration) (string, error) {
	if key == "" {
		return "", ErrInvalidKey
	}
	if expiry <= 0 {
		expiry = time.Hour
	}
	c := ap.client.ServiceClient().NewContainerClient(ap.container).NewBlobClient(key)
	url, err := c.GetSASURL(sas.BlobPermissions{Read: true}, time.Now().Add(expiry), nil)
	if err != nil {
		return "", fmt.Errorf("failed to generate SAS URL: %w", err)
	}
	return url, nil
}

// Close is a no-op for Azure (connections are managed by the SDK).
func (ap *AzureProvider) Close() error {
	return nil
}
