package dataImage

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
)

func Sm() {
	image1, err := os.Open("./images/background.jpg")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	first, err := jpeg.Decode(image1)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer image1.Close()

	image2, err := os.Open("./images/30x40.png")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	secondData, err := png.Decode(image2)
	second := imaging.Resize(secondData, 655, 964, imaging.Lanczos)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer image2.Close()

	image3, err := os.Open("./image.jpg")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	third, err := jpeg.Decode(image3)

	resize := imaging.Resize(third, 625, 928, imaging.Lanczos)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer image3.Close()

	offset := image.Pt(2150, 800)
	offset1 := image.Pt(2165, 818)
	b := first.Bounds()
	image4 := image.NewRGBA(b)
	draw.Draw(image4, b, first, image.ZP, draw.Src)
	draw.Draw(image4, second.Bounds().Add(offset), second, image.ZP, draw.Over)
	draw.Draw(image4, second.Bounds().Add(offset1), resize, image.ZP, draw.Over)

	fourth, err := os.Create("result/smResult.webp")
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}
	final := imaging.Fit(image4, 1200, 1200, imaging.Lanczos)
	webp.Encode(fourth, final, &webp.Options{Lossless: false, Quality: 75})
	defer fourth.Close()
}
