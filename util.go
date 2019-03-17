package imager

import (
	"image"
)

func inRect(rect image.Rectangle, x, y int) bool {
	if rect.Min.X > x || rect.Min.Y > y || rect.Max.X < x || rect.Max.Y < y {
		return false
	} else {
		return true
	}
}
