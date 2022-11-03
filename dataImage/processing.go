package dataImage

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"image/processing/model"
	"log"
	"os"

	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
)

func Processing(
	background string,
	frame string,
	mainImage string,
	result string,
	selectSize model.SizeDetail,
) {

	image1, err := os.Open(background)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	first, err := jpeg.Decode(image1)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer image1.Close()

	image2, err := os.Open(frame)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	secondData, err := png.Decode(image2)
	second := imaging.Resize(secondData, selectSize.FrameSize.X, selectSize.FrameSize.Y, imaging.Lanczos)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer image2.Close()

	image3, err := os.Open(mainImage)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	third, err := jpeg.Decode(image3)

	dstImage128 := imaging.Resize(third, selectSize.ImageSize.X, selectSize.ImageSize.Y, imaging.Lanczos)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer image3.Close()

	offset := image.Pt(selectSize.FrameOffset.X, selectSize.FrameOffset.Y)
	offset1 := image.Pt(selectSize.ImageOffset.X, selectSize.ImageOffset.Y)
	b := first.Bounds()
	image4 := image.NewRGBA(b)
	draw.Draw(image4, b, first, image.Point{}, draw.Src)
	draw.Draw(image4, second.Bounds().Add(offset), second, image.Point{}, draw.Over)
	draw.Draw(image4, second.Bounds().Add(offset1), dstImage128, image.Point{}, draw.Over)

	fourth, err := os.Create(result)
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}
	final := imaging.Fit(image4, 1200, 1200, imaging.Lanczos)
	webp.Encode(fourth, final, &webp.Options{Lossless: false, Quality: 75})
	defer fourth.Close()
}
