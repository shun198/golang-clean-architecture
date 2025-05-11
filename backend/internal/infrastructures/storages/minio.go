package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioStorage struct {
	client     *minio.Client
	bucketName string
	accessKey  string
	secretKey  string
}

func NewMinioStorage() (*MinioStorage, error) {
	// コンテナ内からのアクセス用
	endpoint := os.Getenv("MINIO_ENDPOINT")
	// http://localhost:9001/browserにログインするときのパス
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")
	bucketName := os.Getenv("MINIO_BUCKET_NAME")

	// https://min.io/docs/minio/linux/developers/go/minio-go.html#id3
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, fmt.Errorf("minio error: %w", err)
	}

	ctx := context.Background()
	exists, err := client.BucketExists(ctx, bucketName)
	if err != nil {
		return nil, fmt.Errorf("bucket error: %w", err)
	}

	if !exists {
		if err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{}); err != nil {
			return nil, fmt.Errorf("cannot create bucket: %w", err)
		}

		log.Printf("bucket '%s' created", bucketName)
	}

	return &MinioStorage{
		client:     client,
		bucketName: bucketName,
		accessKey:  accessKey,
		secretKey:  secretKey,
	}, nil
}
