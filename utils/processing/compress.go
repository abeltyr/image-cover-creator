package processing

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"image/processing/utils/s3"
	"log"

	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
)

func Compress(
	id string,
	mainImage []byte,
	typeData string,
) {

	mainImageReader := bytes.NewReader(mainImage)

	var first image.Image
	var err error

	if typeData == "jpg" || typeData == "jpeg" {
		first, err = jpeg.Decode(mainImageReader)
	}
	if typeData == "png" {
		first, err = png.Decode(mainImageReader)
	}

	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	bounds := first.Bounds()
	w := bounds.Dx()
	h := bounds.Dy()

	width := 300
	height := width * h / w

	var da []byte
	fourth := bytes.NewBuffer(da)
	final := imaging.Resize(first, width, height, imaging.Lanczos)
	webp.Encode(fourth, final, &webp.Options{Lossless: false, Quality: 75})

	fileBytes := bytes.NewReader(fourth.Bytes())
	data, err := s3.Upload("Art"+"/"+id+"/"+"compressed.webp", fileBytes, "public-read")

	log.Println("data", data, err)
}
