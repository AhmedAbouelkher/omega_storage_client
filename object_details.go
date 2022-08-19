package omegastorage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type ObjectMetadata struct {
	Metadata map[string]string
}

func (s *Storage) GetMetadata(key string) (*ObjectMetadata, error) {
	cfg := s.Config

	sess, err := s.GetSession()
	if err != nil {
		return nil, err
	}

	client := s3.New(sess)

	md, err := client.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(cfg.Bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return nil, err
	}

	return &ObjectMetadata{
		Metadata: aws.StringValueMap(md.Metadata),
	}, nil
}
