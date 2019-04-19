package imager

import (
	"image"
	"image/color"
)

// ScaleImager
type ScaleImager struct {
	NopImager
	img   image.Image
	scale float64
}

// ColorModel
func (si *ScaleImager) ColorModel() color.Model {
	return si.img.ColorModel()
}

// Bounds
func (si *ScaleImager) Bounds() image.Rectangle {
	rect := si.img.Bounds()
	return image.Rect(0, 0, int(round(float64(rect.Dx())*si.scale)), int(round(float64(rect.Dy())*si.scale)))
}

// At
func (si *ScaleImager) At(x, y int) color.Color {
	rect := si.img.Bounds()
	x = int(round(float64(x) / si.scale))
	if x >= rect.Dx() {
		x = rect.Dx() - 1
	}
	y = int(round(float64(y) / si.scale))
	if y >= rect.Dy() {
		y = rect.Dy() - 1
	}
	return si.img.At(x, y)
}
