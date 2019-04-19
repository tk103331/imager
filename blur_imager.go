package imager

import (
	"image"
	"image/color"
)

// BlurImager
type BlurImager struct {
	img   image.Image
	level int
}

// ColorModel
func (bi *BlurImager) ColorModel() color.Model {
	return bi.img.ColorModel()
}

// Bounds
func (bi *BlurImager) Bounds() image.Rectangle {
	return bi.img.Bounds()
}

// At
func (bi *BlurImager) At(x, y int) color.Color {
	colors := make([]color.Color, 0)
	rect := bi.Bounds()
	for i := -bi.level; i <= bi.level; i++ {
		sx := i + x
		if sx < 0 || sx >= rect.Dx() {
			sx = x - i
		}
		for j := -bi.level; j <= bi.level; j++ {
			sy := j + y
			if sy < 0 || sy >= rect.Dy() {
				sy = y - j
			}
			colors = append(colors, bi.img.At(sx, sy))
		}
	}
	return avgColor(colors)
}
