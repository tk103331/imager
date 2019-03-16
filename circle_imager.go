package imager

import (
	"color"
	"image"
	"math"
)

type CircleMode int

var CircleOuter CircleMode = 0
var CircleLonger CircleMode = 1
var CircleShorter CircleMode = 2

type CircleImager struct {
	Imager
	Mode CircleMode
}

func (ci *CircleImager) Bounds() image.Rectangle {
	x := ci.img.Bounds().Dx()
	y := ci.img.Bounds().Dy()
	shorter, longer := x, y
	if x > y {
		shorter, longer := y, x
	}
	if ci.Mode == CircleOuter {
		mid := math.Sqrt(x*x + y*y)
		return image.Rect(0, 0, mid, mid)
	} else if ci.Mode == CircleShorter {
		return image.Rect(0, 0, shorter, shorter)
	} else if ci.Mode == CircleLonger {
		return image.Rect(0, 0, longer, longer)
	} else {
		return image.Rect(0, 0, shorter, shorter)
	}
}

func (ci *CircleImager) At(x, y int) color.Color {
	return
}
