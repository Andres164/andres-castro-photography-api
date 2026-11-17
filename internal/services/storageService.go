package services

import (
	"context"
	"fmt"
	"io"
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
) (error) {
	objectH := s.bucket.Object(objectName)
	writer := objectH.NewWriter(ctx)

	_, err := io.Copy(writer, reader)
	if err != nil {
		return err
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	return nil
}

func (s *StorageService) Get(
	ctx context.Context,
	objectName string,
) (io.ReadCloser, error) {
	objectH := s.bucket.Object(objectName)

	reader, err := objectH.NewReader(ctx) // Change to WithJSONReads
	if err != nil {
		return nil, fmt.Errorf(
            "create reader for object %q: %w",
            objectName,
            err,
        )
	}

	return reader, nil
}