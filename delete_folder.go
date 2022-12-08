package omegastorage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type DeleteFolderInput struct {
	Bucket string
	Key    string
}

func (s *Storage) DeleteFolder(d *DeleteFolderInput) error {
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

	return s.BatchDelete(
		&BatchDeleteInput{
			Bucket: d.Bucket,
			Keys:   keys,
		},
	)
}
