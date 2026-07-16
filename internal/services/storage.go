package services

import (
	"context"
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

	objectUrl := "http://url/photos/" + objectName

	return objectUrl, nil
}