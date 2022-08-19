package omegastorage

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type ObjectInput struct {
	Reader   io.Reader
	Key      string
	Metadata map[string]string
}

type ObjectUpload struct {
	Location string
	ETag     string
}

func (s *Storage) Upload(obj *ObjectInput) (*ObjectUpload, error) {
	cfg := s.Config

	sess, err := s.GetSession()
	if err != nil {
		return nil, err
	}
	// Create an svc with the session and default options
	svc := s3manager.NewUploader(sess)

	// Upload the file to S3.
	result, err := svc.Upload(&s3manager.UploadInput{
		Bucket:   aws.String(cfg.Bucket),
		Key:      aws.String(obj.Key),
		Body:     obj.Reader,
		Metadata: aws.StringMap(obj.Metadata),
	})
	if err != nil {
		return nil, err
	}

	u := &ObjectUpload{
		Location: result.Location,
		ETag:     *result.ETag,
	}
	return u, nil
}
