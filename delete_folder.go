package omegastorage

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type DeleteFolderInput struct {
	Bucket string
	Key    string
}

func (s *Storage) DeleteFolder(ctx context.Context, d *DeleteFolderInput) error {
	sess, err := s.getSession()
	if err != nil {
		return err
	}
	client := s3.New(sess)
	i := &s3.ListObjectsInput{
		Bucket: aws.String(d.Bucket),
		Prefix: aws.String(d.Key),
	}
	resp, err := client.ListObjects(i)
	if err != nil {
		return err
	}
	var keys []string
	for _, o := range resp.Contents {
		keys = append(keys, *o.Key)
	}
	if len(keys) == 0 {
		return nil
	}
	return s.BatchDelete(
		ctx,
		&BatchDeleteInput{
			Bucket: d.Bucket,
			Keys:   keys,
		},
	)
}
