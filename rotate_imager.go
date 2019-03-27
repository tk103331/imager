package imager

import (
	"image"
	"image/color"
	"math"
)

type RotateImager struct {
	img    image.Image
	radian float64
}

func (ri *RotateImager) ColorModel() color.Model {
	return ri.img.ColorModel()
}

func (ri *RotateImager) Bounds() image.Rectangle {
	rect := ri.img.Bounds()

	x1, y1 := ri.rotatePoint(rect.Min.X, rect.Min.Y, ri.radian)
	x2, y2 := ri.rotatePoint(rect.Min.X, rect.Max.Y, ri.radian)
	x3, y3 := ri.rotatePoint(rect.Max.X, rect.Min.Y, ri.radian)
	x4, y4 := ri.rotatePoint(rect.Max.X, rect.Max.Y, ri.radian)

	return image.Rect(min(x1, x2, x3, x4), min(y1, y2, y3, y4), max(x1, x2, x3, x4), max(y1, y2, y3, y4))
}

func (ri *RotateImager) At(x, y int) color.Color {
	x0, y0 := ri.rotatePoint(x, y, ri.radian)
	return ri.img.At(x0, y0)
}

func (ri *RotateImager) rotatePoint(x, y int, radian float64) (int, int) {
	rect := ri.img.Bounds()

	x0 := rect.Dx() / 2
	y0 := rect.Dy() / 2

	d := distance(x, y, x0, y0)

	r := math.Atan(float64(y-y0) / float64(x-x0))

	r1 := r - radian

	x1 := x0 - int(math.Sin(r1)*d)
	y1 := y0 - int(math.Cos(r1)*d)
	return x1, y1
}
