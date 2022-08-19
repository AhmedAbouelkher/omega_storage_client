package omegastorage

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type ObjectSignedUrlInput struct {
	Bucket   string
	Key      string
	Duration time.Duration
}

func (s *Storage) GetSignedUrl(i *ObjectSignedUrlInput) (string, error) {
	sess, err := s.GetSession()
	if err != nil {
		return "", err
	}

	client := s3.New(sess)

	req, _ := client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(i.Bucket),
		Key:    aws.String(i.Key),
	},
	)

	url, err := req.Presign(i.Duration)

	if err != nil {
		return "", err
	}

	return url, nil
}
