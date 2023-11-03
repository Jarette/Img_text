/*****************************************************************************
*                    
*  Author:           Jarette Greene
*  Email:            jkgreene0406@my.msutexas.edu / jarettegreene09@gmail.com
*  Label:            P03
*  Title:			 Image Ascii Art           
*  Course:           CMPS 4143
*  Semester:         Fall 2023
* 
*  Description:
*	
*		This Package contains functions that create colored text
*
* 
*  Usage:
*    - import to main.go file 
* 
*  Files           
*       N/A 
*****************************************************************************/
package Img_text
//necessary packages 
import (
	//fmt : allows for input and output from and to the console
	"fmt"
	"io/ioutil"
	//os : allows to check if a error occurs when getting a file 
	"os"

	// github.com/fogleman/gg : allows for basic image manipulation 
	"github.com/fogleman/gg"

	// golang.org/x/image/font/gofont/goregular: allows for access to different text fonts in go 
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
	//defers this call to the end of the of the program 
	defer os.Remove(tempFile.Name())

	//checking for error in the writing the font
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
	// final name of the image file created 
	dc.SavePNG("hello.png")

}
