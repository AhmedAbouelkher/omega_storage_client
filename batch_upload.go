package omegastorage

import (
	"context"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type BatchUploadInput struct {
	Objects []*ObjectInput
	Folder  string
}

func (s *Storage) BatchUpload(b *BatchUploadInput) (any, error) {
	cfg := s.Config

	sess, err := s.GetSession()
	if err != nil {
		return nil, err
	}

	// Create an uploader with the session and default options
	svc := s3manager.NewUploader(sess)
	iter := configureIter(b, cfg)

	if err := svc.UploadWithIterator(context.Background(), iter); err != nil {
		return nil, err
	}

	return nil, nil
}

func configureIter(b *BatchUploadInput, cfg *S3Config) *s3manager.UploadObjectsIterator {
	var objects []s3manager.BatchUploadObject
	for _, o := range b.Objects {

		objects = append(objects, s3manager.BatchUploadObject{
			Object: &s3manager.UploadInput{
				Bucket:   aws.String(cfg.Bucket),
				Key:      aws.String(path.Join(b.Folder, o.Key)),
				Body:     o.Reader,
				Metadata: aws.StringMap(o.Metadata),
			},
		})
	}

	return &s3manager.UploadObjectsIterator{Objects: objects}
}
