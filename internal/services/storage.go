package services

import (
	"context"

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