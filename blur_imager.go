package imager

import (
	"image"
	"image/color"
)

type BlurImager struct {
	img   image.Image
	level int
}

func (bi *BlurImager) ColorModel() color.Model {
	return bi.img.ColorModel()
}

func (bi *BlurImager) Bounds() image.Rectangle {
	return bi.img.Bounds()
}

func (bi *BlurImager) At(x, y int) color.Color {
	colors := make([]color.Color, 0)
	for i := x - bi.level; i <= x+bi.level; i++ {
		for j := y - bi.level; j <= y+bi.level; j++ {

			colors = append(colors, bi.img.At(i, j))
		}
	}
	return avgColor(colors)
}
