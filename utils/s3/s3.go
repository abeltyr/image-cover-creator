package s3

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetS3Client() (*s3.S3, error) {

	key := os.Getenv("AWS_ACCESS_KEY_DATA")
	secret := os.Getenv("AWS_SECRET_ACCESS_KEY_DATA")
	location := os.Getenv("AWS_REGION_DATA")

	session, _ := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(key, secret, ""),
		Region:           aws.String(location),
		S3ForcePathStyle: aws.Bool(false), //
	})

	s3Client := s3.New(session)

	return s3Client, nil

}
