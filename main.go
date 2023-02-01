package main

import (
	"image/process/utils/processing"
	"image/process/utils/s3"
	"image/process/utils/size"
	"image/process/utils/webhook"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Id                string `json:"id"`
	ImageProcessingId string `json:"imageProcessingId"`
	Image             string `json:"image"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {

	// err := godotenv.Load(".env")

	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

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

	webhook.Call(id, event.ImageProcessingId)

	return MyResponse{Message: "Done"}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
	// HandleLambdaEvent(MyEvent{Id: "281ff00d-7530-48d9-9070-273cd99aeb02", ImageProcessingId: "a6e9726e-c7b8-4780-ae33-930ed1259dac", Image: "ArtWork/leon-lee-spkbavO3a5A.jpg"})
}
