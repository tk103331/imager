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
	return ri.img.Bounds()
}

func (ri *RotateImager) At(x, y int) color.Color {

	rect := ri.img.Bounds()

	x0 := rect.Dx() / 2
	y0 := rect.Dy() / 2

	d := distance(x, y, x0, y0)

	r := math.Atan(math.Abs(float64(x-x0)) / math.Abs(float64(y-y0)))

	r1 := r - ri.radian

	x1 := x0 - int(math.Sin(r1)*d)
	y1 := y0 - int(math.Cos(r1)*d)

	if r1 > 0 && r1 <= math.Pi/4 {
		x1 = -1
		y1 = -1
	} else if r1 <= math.Pi/2 {
		x1 = -1
		y1 = -1
	} else if r1 <= math.Pi*3/4 {
		// x1 = -1
		// y1 = -1
	} else if r1 <= math.Pi {
		x1 = -1
		y1 = -1
	}
	return ri.img.At(x1, y1)
}
