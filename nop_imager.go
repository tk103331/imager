package imager

import (
	"image"
	"image/color"
)

type NopImager struct {
	img image.Image
}

func (ni *NopImager) ColorModel() color.Model {
	return ni.img.ColorModel()
}

func (ni *NopImager) Bounds() image.Rectangle {
	return ni.img.Bounds()
}

func (ni *NopImager) At(x, y int) color.Color {
	return ni.img.At(x, y)
}
