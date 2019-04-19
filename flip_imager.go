package imager

import (
	"image"
	"image/color"
)

// FlipMode
type FlipMode int

// FlipHorizontal
var FlipHorizontal FlipMode = 1

// FlipVertical
var FlipVertical FlipMode = 2

// FlipImager
type FlipImager struct {
	img  image.Image
	mode FlipMode
}

// ColorModel
func (ri *FlipImager) ColorModel() color.Model {
	return ri.img.ColorModel()
}

// Bounds
func (ri *FlipImager) Bounds() image.Rectangle {
	return ri.img.Bounds()
}

// At
func (ri *FlipImager) At(x, y int) color.Color {
	rect := ri.Bounds()
	if ri.mode == FlipVertical {
		return ri.img.At(x, rect.Dy()-y-1)
	} else {
		// FlipHorizontal
		return ri.img.At(rect.Dx()-x-1, y)
	}
}
