package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type changeable interface {
	Set(x, y int, c color.Color)
	image.Image
}

func main() {
	imgfile, err := os.Open("gopher.png")
	if err != nil {
		panic(err.Error())
	}
	defer imgfile.Close()

	img, err := png.Decode(imgfile)
	if err != nil {
		panic(err.Error())
	}

	var cimg changeable
	var ok bool
	if cimg, ok = img.(changeable); ok {
		addLabel(cimg, 100, 140, "สวัสดีจ้า")
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
		Face: fontface(),
		Dot:  point,
	}
	d.DrawString(label)
}

func fontface() font.Face {
	ttf, _ := ioutil.ReadFile("EkkamaiStandard-Light.ttf")
	f, _ := truetype.Parse(ttf)
	return truetype.NewFace(f, nil)
}
