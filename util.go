package imager

import (
	"image"
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
