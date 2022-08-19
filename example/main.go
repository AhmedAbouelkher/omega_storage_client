package main

import (
	"log"
	"os"

	omegastorage "github.com/AhmedAbouelkher/omega_storage_client"
)

func main() {
	s := omegastorage.Storage{
		Config: &omegastorage.S3Config{
			Bucket: "my-bucket",
		},
	}

	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	out, err := s.Upload(&omegastorage.ObjectInput{
		Key:    "my-key",
		Reader: f,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(out.ETag)
}
