package imgMod

import (
	"os"

	"github.com/fogleman/gg"
	"golang.org/x/image/font/gofont/goregular"
)

func CreatePicText() {

	// Create a temporary file and write the byte slice to it
	tempFile, err := os.CreateTemp("", "font-*.ttf")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.Write(goregular.TTF); err != nil {
		panic(err)
	}

	dc, err := gg.LoadPNG("downloaded_image.png")
	if err != nil {
		panic(err)
	}
	// Get image bounds
	bounds := dc.Bounds()
	// Assign bounds to W and H
	W, H := bounds.Max.X, bounds.Max.Y

	if err := dc.LoadFontFace(tempFile.Name(), 72); err != nil {
		panic(err)
	}

	dc.SetRGB(1, 1, 1)
	dc.Clear()

	dc.SetRGB(.5, 0, 0)
	dc.DrawStringAnchored("Hello, world!", W/2, H/2, 0.5, 0.5)
	dc.Stroke()

	dc.SavePNG("hello.png")
}
