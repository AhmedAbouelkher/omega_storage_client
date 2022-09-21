package main

import (
	"log"
	"os"

	omegastorage "github.com/AhmedAbouelkher/omega_storage_client"
)

func main() {
	s := omegastorage.Storage{
		Config: &omegastorage.S3Config{
			Region:    "us-east-1",
			AccessKey: "AKIAJXQZQQ7QQQQQQQQ",
			Secret:    "123456789",
		},
	}

	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	out, err := s.Upload(&omegastorage.UploadObjectInput{
		Bucket: "my-bucket",
		Key:    "my-key",
		Body:   f,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(out.ETag)
}
