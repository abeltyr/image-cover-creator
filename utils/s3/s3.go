package s3

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetS3Client() (*s3.S3, error) {

	key := os.Getenv("SPACE_KEY")
	secret := os.Getenv("SPACE_SECRET")
	endPoint := os.Getenv("SPACE_ENDPOINT")
	location := os.Getenv("SPACE_LOCATION")

	session, _ := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:         aws.String(endPoint),
		Region:           aws.String(location),
		S3ForcePathStyle: aws.Bool(false), //
	})

	s3Client := s3.New(session)

	return s3Client, nil

}
