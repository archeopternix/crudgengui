package repository

import (
    "bytes"
    "context"
    "crudgengui/internal/model"
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/s3"
)

// ModelToS3 struct represents a model for handling S3 file operations.
type ModelToS3 struct {
    s3Bucket string
    s3Key    string
    s3Client *s3.Client
}

// NewModelToS3 creates a new instance of ModelToS3 with the provided S3 bucket and key.
func NewModelToS3(bucket, key string) (*ModelToS3, error) {
    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
        return nil, err
    }

    client := s3.NewFromConfig(cfg)

    return &ModelToS3{
        s3Bucket: bucket,
        s3Key:    key,
        s3Client: client,
    }, nil
}

// WriteModel writes the given model to S3 in YAML format.
func (ym *ModelToS3) WriteModel(m *model.Model) error {
    var buf bytes.Buffer
    if err := m.WriteYAML(&buf); err != nil {
        return err
    }

    _, err := ym.s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
        Bucket: aws.String(ym.s3Bucket),
        Key:    aws.String(ym.s3Key),
        Body:   bytes.NewReader(buf.Bytes()),
    })
    return err
}

// ReadModel reads the model from S3 and populates the given model instance.
func (ym *ModelToS3) ReadModel(m *model.Model) error {
    resp, err := ym.s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
        Bucket: aws.String(ym.s3Bucket),
        Key:    aws.String(ym.s3Key),
    })
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    return m.ReadYAML(resp.Body)
}
