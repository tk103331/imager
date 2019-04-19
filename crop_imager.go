package imager

import (
	"image"
	"image/color"
)

// CropImager
type CropImager struct {
	img  image.Image
	rect image.Rectangle
}

// ColorModel
func (ci *CropImager) ColorModel() color.Model {
	return ci.img.ColorModel()
}

// Bounds
func (ci *CropImager) Bounds() image.Rectangle {
	return image.Rect(0, 0, ci.rect.Dx()-1, ci.rect.Dy()-1)
}

// At
func (ci *CropImager) At(x, y int) color.Color {
	min := ci.rect.Min
	return ci.img.At(x+min.X, y+min.Y)
}
