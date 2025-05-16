package sss

import (
	"context"
	"errors"
	"log/slog"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	ErrAccessKeyId     = errors.New("MINIO_ACCESS env is not specified")
	ErrSecretAccessKey = errors.New("MINIO_ACCESS env is not specified")
	ErrAlreadyExist    = errors.New("this data is already exist")
)

// s3 for now keeping only post images -> always use same bucket
const bucketName = "post"

type MinioConfig struct {
	log    *slog.Logger
	client *minio.Client
}

func NewMinio(log *slog.Logger, endpoint string, useSsl bool) (*MinioConfig, error) {
	if err := os.Setenv("MINIO_ACCESS", "JABA_SUPER_USER_MINIO"); err != nil {
		return nil, err
	}
	if err := os.Setenv("MINIO_SECRET", "jaba127!368601NO"); err != nil {
		return nil, err
	}

	accessKeyID, ok := os.LookupEnv("MINIO_ACCESS")
	if !ok {
		return nil, ErrAccessKeyId
	}
	secretAccessKey, ok := os.LookupEnv("MINIO_SECRET")
	if !ok {
		return nil, ErrSecretAccessKey
	}

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSsl,
	})
	if err != nil {
		return nil, err
	}
	return &MinioConfig{
		log:    log,
		client: minioClient,
	}, nil
}

func (m *MinioConfig) initBucket(ctx context.Context, bucketName string) error {
	err := m.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := m.client.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			m.log.Warn("We already own this bucket", "id", bucketName)
			return ErrAlreadyExist
		}
		m.log.Error("fail to create bucket", "err", err)
		return err
	}
	return nil
}

func (m *MinioConfig) UploadPostImage(ctx context.Context, id string, filePath string) error {
	err := m.initBucket(ctx, bucketName)
	if err != nil && !errors.Is(err, ErrAlreadyExist) {
		return err
	}
	info, err := m.client.FPutObject(ctx, bucketName, id, filePath, minio.PutObjectOptions{})
	if err != nil {
		m.log.Error("fail to save file", "bucket", bucketName, "id", id, "err", err)
		return err
	}

	m.log.Info("successfully upload file", "info", info)
	return nil
}

func (m *MinioConfig) DownloadPostImage(ctx context.Context, id, filePath string) error {
	err := m.client.FGetObject(ctx, bucketName, id, filePath, minio.GetObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (m *MinioConfig) DeletePostImage(ctx context.Context, id string) error {
	return m.client.RemoveObject(ctx, bucketName, id, minio.RemoveObjectOptions{})
}
