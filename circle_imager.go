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

type CircleImager struct {
	*Imager
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
	rect := ci.img.Bounds()
	w := rect.Dx()
	h := rect.Dy()
	diameter := ci.diameter()

	point := image.Point{int((w - diameter) / 2), int((h - diameter) / 2)}
	return image.Rect(point.X, point.Y, point.X+diameter, point.Y+diameter)
}

func (ci *CircleImager) At(x, y int) color.Color {
	rect := ci.img.Bounds()
	w := rect.Dx()
	h := rect.Dy()
	diameter := ci.diameter()

	distance := math.Sqrt(float64((x-w/2)*(x-w/2) + (y-h/2)*(y-h/2)))
	if distance <= float64(diameter)/2 {
		if inRect(rect, x, y) {
			return ci.img.At(x, y)
		} else {
			return color.White
		}
	} else {
		return ci.img.At(-1, -1)
	}
}
