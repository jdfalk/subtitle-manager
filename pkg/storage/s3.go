// file: pkg/storage/s3.go
// version: 1.0.0
// guid: 45678901-23de-f456-7890-123456789012

package storage

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3Provider implements StorageProvider for Amazon S3 storage.
type S3Provider struct {
	s3Client *s3.S3
	bucket   string
}

// NewS3Provider creates a new S3 storage provider.
func NewS3Provider(config StorageConfig) (*S3Provider, error) {
	if config.S3Bucket == "" {
		return nil, fmt.Errorf("%w: S3 bucket name is required", ErrConfigurationMissing)
	}
	
	// Configure AWS session
	awsConfig := &aws.Config{
		Region: aws.String(config.S3Region),
	}
	
	// Set custom endpoint if provided (for S3-compatible services)
	if config.S3Endpoint != "" {
		awsConfig.Endpoint = aws.String(config.S3Endpoint)
		awsConfig.S3ForcePathStyle = aws.Bool(true)
	}
	
	// Set credentials if provided
	if config.S3AccessKey != "" && config.S3SecretKey != "" {
		awsConfig.Credentials = credentials.NewStaticCredentials(
			config.S3AccessKey,
			config.S3SecretKey,
			"", // session token
		)
	}
	
	sess, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrConnectionFailed, err)
	}
	
	s3Client := s3.New(sess)
	
	return &S3Provider{
		s3Client: s3Client,
		bucket:   config.S3Bucket,
	}, nil
}

// Store saves content to S3.
func (sp *S3Provider) Store(ctx context.Context, key string, content io.Reader, contentType string) error {
	if key == "" {
		return ErrInvalidKey
	}
	
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	
	input := &s3.PutObjectInput{
		Bucket:      aws.String(sp.bucket),
		Key:         aws.String(key),
		Body:        aws.ReadSeekCloser(content),
		ContentType: aws.String(contentType),
	}
	
	_, err := sp.s3Client.PutObjectWithContext(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to store object in S3: %w", err)
	}
	
	return nil
}

// Retrieve reads content from S3.
func (sp *S3Provider) Retrieve(ctx context.Context, key string) (io.ReadCloser, error) {
	if key == "" {
		return nil, ErrInvalidKey
	}
	
	input := &s3.GetObjectInput{
		Bucket: aws.String(sp.bucket),
		Key:    aws.String(key),
	}
	
	result, err := sp.s3Client.GetObjectWithContext(ctx, input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchKey:
				return nil, ErrNotFound
			}
		}
		return nil, fmt.Errorf("failed to retrieve object from S3: %w", err)
	}
	
	return result.Body, nil
}

// Delete removes content from S3.
func (sp *S3Provider) Delete(ctx context.Context, key string) error {
	if key == "" {
		return ErrInvalidKey
	}
	
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(sp.bucket),
		Key:    aws.String(key),
	}
	
	_, err := sp.s3Client.DeleteObjectWithContext(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to delete object from S3: %w", err)
	}
	
	return nil
}

// Exists checks if content exists in S3.
func (sp *S3Provider) Exists(ctx context.Context, key string) (bool, error) {
	if key == "" {
		return false, ErrInvalidKey
	}
	
	input := &s3.HeadObjectInput{
		Bucket: aws.String(sp.bucket),
		Key:    aws.String(key),
	}
	
	_, err := sp.s3Client.HeadObjectWithContext(ctx, input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchKey, "NotFound":
				return false, nil
			}
		}
		return false, fmt.Errorf("failed to check object existence in S3: %w", err)
	}
	
	return true, nil
}

// List returns keys matching the given prefix in S3.
func (sp *S3Provider) List(ctx context.Context, prefix string) ([]string, error) {
	var keys []string
	
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(sp.bucket),
		Prefix: aws.String(prefix),
	}
	
	err := sp.s3Client.ListObjectsV2PagesWithContext(ctx, input, func(page *s3.ListObjectsV2Output, lastPage bool) bool {
		for _, obj := range page.Contents {
			if obj.Key != nil {
				keys = append(keys, *obj.Key)
			}
		}
		return !lastPage
	})
	
	if err != nil {
		return nil, fmt.Errorf("failed to list objects in S3: %w", err)
	}
	
	return keys, nil
}

// GetURL returns a pre-signed URL for the object in S3.
func (sp *S3Provider) GetURL(ctx context.Context, key string, expiry time.Duration) (string, error) {
	if key == "" {
		return "", ErrInvalidKey
	}
	
	if expiry <= 0 {
		expiry = 1 * time.Hour // Default expiry
	}
	
	req, _ := sp.s3Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(sp.bucket),
		Key:    aws.String(key),
	})
	
	url, err := req.Presign(expiry)
	if err != nil {
		return "", fmt.Errorf("failed to generate pre-signed URL: %w", err)
	}
	
	return url, nil
}

// Close is a no-op for S3 (connections are managed by the SDK).
func (sp *S3Provider) Close() error {
	return nil
}