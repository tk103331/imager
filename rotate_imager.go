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

	x0 := rect.Dx() / 2
	y0 := rect.Dy() / 2

	x1, y1 := ri.rotatePoint(x0, y0, rect.Min.X, rect.Min.Y, ri.radian)
	x2, y2 := ri.rotatePoint(x0, y0, rect.Max.X, rect.Min.Y, ri.radian)
	x3, y3 := ri.rotatePoint(x0, y0, rect.Max.X, rect.Max.Y, ri.radian)
	x4, y4 := ri.rotatePoint(x0, y0, rect.Min.X, rect.Max.Y, ri.radian)

	return image.Rect(min(x1, x2, x3, x4), min(y1, y2, y3, y4), max(x1, x2, x3, x4), max(y1, y2, y3, y4))
}

func (ri *RotateImager) At(x, y int) color.Color {
	rect := ri.img.Bounds()
	x0 := rect.Dx() / 2
	y0 := rect.Dy() / 2
	x1, y1 := ri.rotatePoint(x0, y0, x, y, -ri.radian)
	return ri.img.At(x1, y1)
}

func (ri *RotateImager) rotatePoint(x0, y0 int, x, y int, r float64) (int, int) {

	var r0 float64
	if x == x0 && y == y0 {
		return x, y
	} else if x == x0 {
		if y < y0 {
			r0 = math.Pi / 2
		} else {
			r0 = math.Pi * 3 / 2
		}
	} else if y == y0 {
		if x < x0 {
			r0 = 0
		} else {
			r0 = math.Pi
		}
	} else {
		r0 = math.Abs(math.Atan(float64(y-y0) / float64(x-x0)))
		if x < x0 && y < y0 {
			r0 = r0
		} else if x > x0 && y < y0 {
			r0 = math.Pi - r0
		} else if x > x0 && y > y0 {
			r0 = math.Pi + r0
		} else if x < x0 && y > y0 {
			r0 = math.Pi*2 - r0
		}
	}

	r1 := modRad(r0 + r)

	for r1 > math.Pi*2 {
		r1 = r1 - math.Pi*2
	}

	d := distance(x, y, x0, y0)
	x1, y1 := 0, 0
	if 0 < r1 && r1 < math.Pi/2 {
		r1 = r1
		x1 = x0 - int(math.Cos(r1)*d)
		y1 = y0 - int(math.Sin(r1)*d)
	} else if math.Pi/2 < r1 && r1 < math.Pi {
		r1 = math.Pi - r1
		x1 = x0 + int(math.Cos(r1)*d)
		y1 = y0 - int(math.Sin(r1)*d)
	} else if math.Pi < r1 && r1 < math.Pi*3/2 {
		r1 = r1 - math.Pi
		x1 = x0 + int(math.Cos(r1)*d)
		y1 = y0 + int(math.Sin(r1)*d)
	} else if math.Pi*3/2 < r1 && r1 < math.Pi*2 {
		r1 = math.Pi*2 - r1
		x1 = x0 - int(math.Cos(r1)*d)
		y1 = y0 + int(math.Sin(r1)*d)
	}

	return x1, y1
}
