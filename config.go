package omegastorage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
)

type S3Config struct {
	Region           string
	AccessKey        string
	Secret           string
	Endpoint         string
	EndpointResolver endpoints.ResolverFunc
}

func (c *S3Config) GetSession() (*session.Session, error) {
	cfg := &aws.Config{
		Region:      aws.String(c.Region),
		Credentials: credentials.NewStaticCredentials(c.AccessKey, c.Secret, ""),
	}

	if c.Endpoint != "" {
		cfg.Endpoint = aws.String(c.Endpoint)
	}

	if c.EndpointResolver != nil {
		cfg.EndpointResolver = c.EndpointResolver
	}

	s, err := session.NewSession(cfg)
	if err != nil {
		return nil, err
	}

	return s, nil
}
