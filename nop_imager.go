package imager

import (
	"image"
	"image/color"
)

// NopImager
type NopImager struct {
	img image.Image
}

// ColorModel
func (ni *NopImager) ColorModel() color.Model {
	return ni.img.ColorModel()
}

// Bounds
func (ni *NopImager) Bounds() image.Rectangle {
	return ni.img.Bounds()
}

// At
func (ni *NopImager) At(x, y int) color.Color {
	return ni.img.At(x, y)
}
