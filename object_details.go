package omegastorage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type ObjectMetadataInput struct {
	Bucket string
	Key    string
}

type ObjectMetadataOutput struct {
	Metadata map[string]string
}

func (s *Storage) GetMetadata(i *ObjectMetadataInput) (*ObjectMetadataOutput, error) {
	sess, err := s.getSession()
	if err != nil {
		return nil, err
	}

	client := s3.New(sess)

	md, err := client.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(i.Bucket),
		Key:    aws.String(i.Key),
	})

	if err != nil {
		return nil, err
	}

	return &ObjectMetadataOutput{
		Metadata: aws.StringValueMap(md.Metadata),
	}, nil
}
