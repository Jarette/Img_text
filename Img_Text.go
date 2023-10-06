package Img_text

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fogleman/gg"
	"golang.org/x/image/font/gofont/goregular"
)

// Screen_text displays colored text to the screen
func Screen_text(text string) {
	fmt.Printf("\033[31m %s\033[0m",text)
	fmt.Printf("\033[32m %s\033[0m",text)
}

// Color_imgtext prints a colored text image to the screen
func Color_imgtext(text string) {
	const W = 500
	const H = 300

	// Create a temporary file and write the byte slice to it
	tempFile, err := ioutil.TempFile("", "font-*.ttf")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.Write(goregular.TTF); err != nil {
		panic(err)
	}

	dc := gg.NewContext(W, H)

	if err := dc.LoadFontFace(tempFile.Name(), 72); err != nil {
		panic(err)
	}

	dc.SetRGB(1, 1, 1)
	dc.Clear()

	dc.SetRGB(.5, 0, 0)
	dc.DrawStringAnchored(text, W/2, H/2, 0.5, 0.5)
	dc.Stroke()

	dc.SavePNG("hello.png")

}
