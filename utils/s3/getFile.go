package s3

import (
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetFile(
	key string,
) ([]byte, error) {

	bucket := os.Getenv("AWS_BUCKET")

	s3Client, err := GetS3Client()
	if err != nil {
		return nil, err
	}

	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	result, err := s3Client.GetObject(input)
	if err != nil {
		return nil, err
	}

	fileByte, err := io.ReadAll(result.Body)

	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	return fileByte, nil
}
