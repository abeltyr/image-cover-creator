package main

import (
	"image/processing/utils/processing"
	"image/processing/utils/s3"
	"image/processing/utils/size"
	"log"
	"strings"

	"github.com/joho/godotenv"
)

// "image/processing/dataImage"
// "image/processing/size"

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	id := "asas-1212"

	background, err := s3.GetFile("Frame/background.jpg")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	title := "Frame/image.jpg"
	res1 := strings.SplitN(title, ".", -1)

	mainImage, err := s3.GetFile(title)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	log.Println("downloaded")
	processing.ImageLayering(
		id,
		background,
		mainImage,
		res1[len(res1)-1],
		size.Portrait("xl"),
	)

	processing.ImageLayering(
		id,
		background,
		mainImage,
		res1[len(res1)-1],
		size.Portrait("lg"),
	)

	processing.ImageLayering(
		id,
		background,
		mainImage,
		res1[len(res1)-1],
		size.Portrait("md"),
	)

	processing.ImageLayering(
		id,
		background,
		mainImage,
		res1[len(res1)-1],
		size.Portrait("sm"),
	)

	processing.Compress(
		id,
		mainImage,
		res1[len(res1)-1],
	)
}
