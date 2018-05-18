package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

type Changeable interface {
	Set(x, y int, c color.Color)
	image.Image
}

func main() {
	imgfile, err := os.Open("odds_logo.png")
	if err != nil {
		panic(err.Error())
	}
	defer imgfile.Close()

	img, err := png.Decode(imgfile)
	if err != nil {
		panic(err.Error())
	}

	var cimg Changeable
	var ok bool
	if cimg, ok = img.(Changeable); ok {
		addLabel(cimg, 20, 30, "Hello")
	} else {
		fmt.Println("no luck")
	}

	f, err := os.Create("hello-go.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := png.Encode(f, cimg); err != nil {
		panic(err)
	}

}

func addLabel(img draw.Image, x, y int, label string) {
	col := color.RGBA{200, 100, 0, 255}
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}
