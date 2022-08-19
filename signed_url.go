package omegastorage

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (s *Storage) GetSignedUrl(key string, duration time.Duration) (string, error) {
	cfg := s.Config

	sess, err := s.GetSession()
	if err != nil {
		return "", err
	}

	client := s3.New(sess)

	i := &s3.GetObjectInput{
		Bucket: aws.String(cfg.Bucket),
		Key:    aws.String(key),
	}

	req, _ := client.GetObjectRequest(i)

	url, err := req.Presign(duration)

	if err != nil {
		return "", err
	}

	return url, nil
}
