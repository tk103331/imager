package imager

import (
	"image"
	"image/color"
)

// RoundImager
type RoundImager struct {
	img    image.Image
	radius int
}

// ColorModel
func (ri *RoundImager) ColorModel() color.Model {
	return ri.img.ColorModel()
}

// Bounds
func (ri *RoundImager) Bounds() image.Rectangle {
	return ri.img.Bounds()
}

// At
func (ri *RoundImager) At(x, y int) color.Color {
	rect := ri.Bounds()
	w := rect.Dx()
	h := rect.Dy()
	if x < ri.radius && y < ri.radius {
		distance := distance(x, y, ri.radius, ri.radius)
		if distance >= float64(ri.radius) {
			return OuterColor
		}
	}
	if x < ri.radius && y > (h-ri.radius) {
		distance := distance(x, y, ri.radius, h-ri.radius)
		if distance >= float64(ri.radius) {
			return OuterColor
		}
	}
	if x > (w-ri.radius) && y > (h-ri.radius) {
		distance := distance(x, y, w-ri.radius, h-ri.radius)
		if distance >= float64(ri.radius) {
			return OuterColor
		}
	}
	if x > (w-ri.radius) && y < ri.radius {
		distance := distance(x, y, w-ri.radius, ri.radius)
		if distance >= float64(ri.radius) {
			return OuterColor
		}
	}
	return ri.img.At(x, y)
}
