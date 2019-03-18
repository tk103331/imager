package imager

import (
	"image"
	"image/color"
	"math"
)

type ScaleImager struct {
	NopImager
	img   image.Image
	scale float64
}

func (si *ScaleImager) ColorModel() color.Model {
	return si.img.ColorModel()
}

func (si *ScaleImager) Bounds() image.Rectangle {
	rect := si.img.Bounds()
	return image.Rect(0, 0, int(math.Round(float64(rect.Dx())*si.scale)), int(math.Round(float64(rect.Dy())*si.scale)))
}

func (si *ScaleImager) At(x, y int) color.Color {
	return si.img.At(int(math.Round(float64(x)/si.scale)), int(math.Round(float64(y)/si.scale)))
}
