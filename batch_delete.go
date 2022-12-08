package omegastorage

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type BatchDeleteInput struct {
	Bucket string
	Keys   []string
}

func (s *Storage) BatchDelete(b *BatchDeleteInput) error {
	sess, err := s.getSession()
	if err != nil {
		return err
	}

	bd := s3manager.NewBatchDelete(sess)
	itr := configureDeleteIter(b)

	if err := bd.Delete(context.Background(), itr); err != nil {
		return err
	}

	return nil
}

func configureDeleteIter(b *BatchDeleteInput) *s3manager.DeleteObjectsIterator {
	var objects []s3manager.BatchDeleteObject
	for _, k := range b.Keys {

		objects = append(objects, s3manager.BatchDeleteObject{
			Object: &s3.DeleteObjectInput{
				Bucket: aws.String(b.Bucket),
				Key:    aws.String(k),
			},
		})
	}

	return &s3manager.DeleteObjectsIterator{Objects: objects}
}
