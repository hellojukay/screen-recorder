package main

import (
	"github.com/kbinani/screenshot"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"os"
	"time"
)

func main() {
	var n = 10
	palette := append(palette.WebSafe, color.Transparent)
	outGif := &gif.GIF{}

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(0)
		time.Sleep(time.Duration(100) * time.Millisecond)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		palettedImage := image.NewPaletted(bounds, palette)
		draw.Draw(palettedImage, bounds, img, image.ZP, draw.Src)
		outGif.Image = append(outGif.Image, palettedImage)
		outGif.Delay = append(outGif.Delay, 100)
	}
	fileName := "test.gif"

	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	gif.EncodeAll(f, outGif)
}
