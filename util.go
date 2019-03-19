package imager

import (
	"image"
	"image/color"
	"math"
)

func inRect(rect image.Rectangle, x, y int) bool {
	if rect.Min.X > x || rect.Min.Y > y || rect.Max.X < x || rect.Max.Y < y {
		return false
	} else {
		return true
	}
}

func distance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64(pow2(x1-x2) + pow2(y1-y2)))
}

func pow2(v int) float64 {
	return math.Pow(float64(v), 2)
}

func avgColor(colors []color.Color) color.Color {
	var sumR float64 = 0
	var sumG float64 = 0
	var sumB float64 = 0
	var sumA float64 = 0
	for _, c := range colors {
		r, g, b, a := c.RGBA()
		sumR += float64(r >> 8)
		sumG += float64(g >> 8)
		sumB += float64(b >> 8)
		sumA += float64(a >> 8)
	}
	count := float64(len(colors))
	return color.RGBA{uint8(sumR / count), uint8(sumG / count), uint8(sumB / count), uint8(sumA / count)}
}

func round(value float64) float64 {
	return math.Ceil(value + 0.5)
}
