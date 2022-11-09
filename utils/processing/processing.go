package processing

import (
	"bytes"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"image/processing/model"
	"image/processing/utils/s3"
	"log"

	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
)

func ImageLayering(
	id string,
	background []byte,
	mainImage []byte,
	typeData string,
	selectSize model.SizeDetail,
) {

	backgroundReader := bytes.NewReader(background)
	mainImageReader := bytes.NewReader(mainImage)

	first, err := jpeg.Decode(backgroundReader)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}

	frame, err := s3.GetFile(selectSize.Name)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	frameReader := bytes.NewReader(frame)
	secondData, err := png.Decode(frameReader)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	second := imaging.Resize(secondData, selectSize.FrameSize.X, selectSize.FrameSize.Y, imaging.Lanczos)

	var third image.Image

	if typeData == "jpg" || typeData == "jpeg" {
		third, err = jpeg.Decode(mainImageReader)
	}

	if typeData == "png" {
		third, err = png.Decode(mainImageReader)
	}

	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	dstImage128 := imaging.Resize(third, selectSize.ImageSize.X, selectSize.ImageSize.Y, imaging.Lanczos)

	offset := image.Pt(selectSize.FrameOffset.X, selectSize.FrameOffset.Y)
	offset1 := image.Pt(selectSize.ImageOffset.X, selectSize.ImageOffset.Y)
	b := first.Bounds()
	image4 := image.NewRGBA(b)
	draw.Draw(image4, b, first, image.Point{}, draw.Src)
	draw.Draw(image4, second.Bounds().Add(offset), second, image.Point{}, draw.Over)
	draw.Draw(image4, second.Bounds().Add(offset1), dstImage128, image.Point{}, draw.Over)

	var da []byte
	fourth := bytes.NewBuffer(da)
	final := imaging.Fit(image4, 1400, 1400, imaging.Lanczos)
	webp.Encode(fourth, final, &webp.Options{Lossless: false, Quality: 75})

	fileBytes := bytes.NewReader(fourth.Bytes())
	data, err := s3.Upload("Art"+"/"+id+"/"+selectSize.Title+".webp", fileBytes, "public-read")

	log.Println("data", data, err)

}
