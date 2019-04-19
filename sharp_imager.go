package imager

import (
	"image"
	"image/color"
)

var filter []int = []int{-1, -1, -1, -1, 9, -1, -1, -1, -1}

type SharpImager struct {
	img image.Image
}

func (si *SharpImager) ColorModel() color.Model {
	return si.img.ColorModel()
}

func (si *SharpImager) Bounds() image.Rectangle {
	return si.img.Bounds()
}

func (si *SharpImager) At(x, y int) color.Color {
	var sumR, sumG, sumB, sumA int = 0, 0, 0, 0

	for i, f := range filter {
		m := i%3 - 1
		n := i/3 - 1
		c := si.img.At(x+m, y+n)
		r, g, b, a := c.RGBA()
		sumR += int(r>>8) * f
		sumG += int(g>>8) * f
		sumB += int(b>>8) * f
		sumA += int(a>>8) * f

	}

	return color.RGBA{cc(sumR), cc(sumG), cc(sumB), cc(sumA)}
}

func cc(c int) uint8 {
	if c < 0 {
		c = 0
	} else if c > 255 {
		c = 255
	}
	return uint8(c)
}
