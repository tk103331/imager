package imager

import (
	"image"
	"image/color"
	"math"
)

type CircleMode int

var CircleOuter CircleMode = 0
var CircleLonger CircleMode = 1
var CircleShorter CircleMode = 2

var InnerColor color.Color = color.White
var OuterColor color.Color = color.RGBA{0, 0, 0, 0}

type CircleImager struct {
	img  image.Image
	mode CircleMode
}

func (ci *CircleImager) diameter() int {
	rect := ci.img.Bounds()
	w := rect.Dx()
	h := rect.Dy()
	shorter, longer := w, h
	if w > h {
		shorter, longer = h, w
	}
	diameter := shorter
	switch ci.mode {
	case CircleOuter:
		diameter = int(math.Sqrt(float64(w*w + h*h)))
	case CircleShorter:
		diameter = shorter
	case CircleLonger:
		diameter = longer
	}
	return diameter
}

func (ci *CircleImager) ColorModel() color.Model {
	return ci.img.ColorModel()
}

func (ci *CircleImager) Bounds() image.Rectangle {
	diameter := ci.diameter()
	return image.Rect(0, 0, diameter, diameter)
}

func (ci *CircleImager) At(x, y int) color.Color {
	rect := ci.img.Bounds()
	w := rect.Dx()
	h := rect.Dy()
	diameter := ci.diameter()

	x = x + (w-diameter)/2
	y = y + (h-diameter)/2

	distance := distance(x, y, w/2, h/2)
	if distance <= float64(diameter)/2 {
		if inRect(rect, x, y) {
			return ci.img.At(x, y)
		} else {
			return InnerColor
		}
	} else {
		return OuterColor
	}
}
