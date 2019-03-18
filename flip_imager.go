package imager

import (
	"image"
	"image/color"
)

type FlipMode int

var FlipHorizontal FlipMode = 1
var FlipVertical FlipMode = 2

type FlipImager struct {
	img  image.Image
	mode FlipMode
}

func (ri *FlipImager) ColorModel() color.Model {
	return ri.img.ColorModel()
}

func (ri *FlipImager) Bounds() image.Rectangle {
	return ri.img.Bounds()
}

func (ri *FlipImager) At(x, y int) color.Color {
	rect := ri.Bounds()
	if ri.mode == FlipVertical {
		return ri.img.At(x, rect.Dy()-y-1)
	} else {
		// FlipHorizontal
		return ri.img.At(rect.Dx()-x-1, y)
	}
}
