package main

import (
	"image/process/utils/processing"
	"image/process/utils/s3"
	"image/process/utils/size"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

type MyEvent struct {
	Id    string `json:"id"`
	Image string `json:"image"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	id := event.Id

	background, err := s3.GetFile("Frame/background.jpg")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	title := event.Image

	mainImage, err := s3.GetFile(title)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	log.Println("downloaded")
	processing.ImageLayering(
		id,
		background,
		mainImage,
		size.Portrait("xl"),
	)

	processing.ImageLayering(
		id,
		background,
		mainImage,
		size.Portrait("lg"),
	)

	processing.ImageLayering(
		id,
		background,
		mainImage,
		size.Portrait("md"),
	)

	processing.ImageLayering(
		id,
		background,
		mainImage,
		size.Portrait("sm"),
	)

	processing.Compress(
		id,
		mainImage,
	)

	// TODO: setup http request

	return MyResponse{Message: "Done"}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
