package main

import (
	"image/processing/dataImage"
	"image/processing/size"
)

func main() {

	background := "./images/background.jpg"
	image := "image.jpg"

	dataImage.Processing(
		background,
		"./images/60x90P.png",
		image,
		"result/xlResult.webp",
		size.Portrait("xl"),
	)

	dataImage.Processing(
		background,
		"./images/50x80.png",
		image,
		"result/lgResult.webp",
		size.Portrait("lg"),
	)
	dataImage.Processing(
		background,
		"./images/40x60.png",
		image,
		"result/mdResult.webp",
		size.Portrait("md"),
	)
	dataImage.Processing(
		background,
		"./images/30x40.png",
		image,
		"result/smResult.webp",
		size.Portrait("sm"),
	)

	dataImage.Compress(
		image,
		"result/compressed.webp",
		10,
	)
}
