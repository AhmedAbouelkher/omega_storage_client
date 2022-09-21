package main

import (
	"fmt"
	"os"

	omegastorage "github.com/AhmedAbouelkher/omega_storage_client"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	accessKeyID     = os.Getenv("ACCESS_KEY_ID")
	secretAccessKey = os.Getenv("SECRET_ACCESS_KEY")
)

func main() {
	stg := omegastorage.Storage{
		Config: &omegastorage.S3Config{
			Region:    "us-east-1",
			AccessKey: accessKeyID,
			Secret:    secretAccessKey,
			EndpointResolver: func(service, region string, opts ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
				return endpoints.ResolvedEndpoint{
					URL:         "http://localhost:9000",
					PartitionID: endpoints.AwsPartitionID,
				}, nil
			},
		},
	}

	svc, err := stg.S3Client()
	if err != nil {
		panic(err)
	}

	out, err := svc.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		panic(err)
	}

	for _, b := range out.Buckets {
		fmt.Println(*b.Name)
	}
}
