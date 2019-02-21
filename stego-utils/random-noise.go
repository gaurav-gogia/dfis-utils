package main

import (
	"image"
	"image/png"
	"math/rand"
	"os"
)

func noise() {
	img := image.NewRGBA(image.Rect(0, 0, 500, 500))

	for pix := 0; pix < 500*500; pix++ {
		offset := 4 * pix
		img.Pix[0+offset] = uint8(rand.Intn(256))
		img.Pix[1+offset] = uint8(rand.Intn(256))
		img.Pix[2+offset] = uint8(rand.Intn(256))
		img.Pix[3+offset] = 255
	}

	out, err := os.Create("noise.png")
	handerr(err)

	handerr(png.Encode(out, img))
}
