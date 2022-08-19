package omegastorage

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type UploadObjectInput struct {
	Body     io.Reader
	Bucket   string
	Key      string
	Metadata map[string]string
}

type ObjectUploadOutput struct {
	Location string
	ETag     string
}

func (s *Storage) Upload(u *UploadObjectInput) (*ObjectUploadOutput, error) {
	sess, err := s.GetSession()
	if err != nil {
		return nil, err
	}
	// Create an svc with the session and default options
	svc := s3manager.NewUploader(sess)

	// Upload the file to S3.
	result, err := svc.Upload(&s3manager.UploadInput{
		Bucket:   aws.String(u.Bucket),
		Key:      aws.String(u.Key),
		Body:     u.Body,
		Metadata: aws.StringMap(u.Metadata),
	})
	if err != nil {
		return nil, err
	}

	ur := &ObjectUploadOutput{
		Location: result.Location,
		ETag:     *result.ETag,
	}
	return ur, nil
}
