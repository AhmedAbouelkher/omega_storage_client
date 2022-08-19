package omegastorage

import "github.com/aws/aws-sdk-go/aws/session"

type Storage struct {
	Config *S3Config
}

func (s *Storage) GetSession() (*session.Session, error) {
	return s.Config.GetSession()
}
