package dataImage

import (
	"image/jpeg"
	"log"
	"os"

	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
)

func Compress(
	mainImage string,
	result string,
	compressValue int,
) {

	image, err := os.Open(mainImage)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	first, err := jpeg.Decode(image)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	bounds := first.Bounds()
	w := bounds.Dx()
	h := bounds.Dy()
	defer image.Close()

	fourth, err := os.Create(result)
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}
	final := imaging.Resize(first, w/compressValue, h/compressValue, imaging.Lanczos)
	webp.Encode(fourth, final, &webp.Options{Lossless: false, Quality: 75})
	defer fourth.Close()
}
