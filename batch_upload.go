package omegastorage

import (
	"context"
	"errors"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type BatchUploadInput struct {
	Bucket  string
	Objects []*UploadObjectInput
	Folder  string
}

func (s *Storage) BatchUpload(ctx context.Context, b *BatchUploadInput) error {
	sess, err := s.getSession()
	if err != nil {
		return err
	}
	if len(b.Objects) == 0 {
		return errors.New("no objects to upload")
	}
	// Create an uploader with the session and default options
	svc := s3manager.NewUploader(sess)
	iter := configureUploadIter(b)
	if err := svc.UploadWithIterator(ctx, iter); err != nil {
		return err
	}

	return nil
}

func configureUploadIter(b *BatchUploadInput) *s3manager.UploadObjectsIterator {
	var objects []s3manager.BatchUploadObject
	for _, o := range b.Objects {
		bkt := o.Bucket
		if bkt == "" {
			bkt = b.Bucket
		}
		objects = append(objects, s3manager.BatchUploadObject{
			Object: &s3manager.UploadInput{
				Bucket:   aws.String(bkt),
				Key:      aws.String(path.Join(b.Folder, o.Key)),
				Body:     o.Body,
				Metadata: aws.StringMap(o.Metadata),
			},
		})
	}
	return &s3manager.UploadObjectsIterator{Objects: objects}
}
