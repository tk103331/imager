package imager

import (
	"image"
	"image/color"
)

type CropImager struct {
	*Imager
	rect image.Rectangle
}

func (ci *CropImager) ColorModel() color.Model {
	return ci.img.ColorModel()
}

func (ci *CropImager) Bounds() image.Rectangle {
	return ci.rect
}

func (ci *CropImager) At(x, y int) color.Color {
	return ci.img.At(x, y)
}
