package services

import (
	"context"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/storage"
)

type StorageService struct {
	client *storage.Client
	bucket *storage.BucketHandle
}

func NewStorageService(client *storage.Client, bucketName string) *StorageService {
	return &StorageService{
		client: client,
		bucket: client.Bucket(bucketName),
	}
}

func NewStorageClient(ctx context.Context) (*storage.Client, error) {
	return storage.NewClient(ctx)
}

func (s *StorageService) HealthCheck(ctx context.Context) error {
	_, err := s.bucket.Attrs(ctx)
	return err
}

func (s *StorageService) Upload(
    ctx context.Context,
    reader io.Reader,
    objectName string,
) (string, error) {
	objectH := s.bucket.Object(objectName)
	writer := objectH.NewWriter(ctx)

	_, err := io.Copy(writer, reader)
	if err != nil {
		return "", err
	}

	err = writer.Close()
	if err != nil {
		return "", err
	}

	return objectName, nil
}

func (s *StorageService) Get(
	ctx context.Context,
	objectName string,
) (string, error) {
	 signedURL, err := s.bucket.SignedURL(
        objectName,
        &storage.SignedURLOptions{
            Scheme:  storage.SigningSchemeV4,
            Method:  "GET",
            Expires: time.Now().Add(15 * time.Minute),
        },
    )
    if err != nil {
        return "", fmt.Errorf(
            "generate signed URL for object %q: %w",
            objectName,
            err,
        )
    }

	return signedURL, nil
}