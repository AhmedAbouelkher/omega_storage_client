package omegastorage

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Storage struct {
	Config *S3Config
}

func (s *Storage) getSession() (*session.Session, error) {
	return s.Config.getSession()
}

func (s *Storage) S3Client() (*s3.S3, error) {
	sess, err := s.Config.getSession()
	if err != nil {
		return nil, err
	}
	return s3.New(sess), nil
}
