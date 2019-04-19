package imager

import (
	"image"
)
import "image/color"

// ImagerWrapper
type ImagerWrapper struct {
	imager Imager
}

// Imager
type Imager interface {
	ColorModel() color.Model
	Bounds() image.Rectangle
	At(x, y int) color.Color
}

// New create a ImagerWrapper.
func New(img image.Image) *ImagerWrapper {
	return &ImagerWrapper{imager: &NopImager{img}}
}

// Circle
func (iw *ImagerWrapper) Circle(mode CircleMode) *ImagerWrapper {
	return &ImagerWrapper{imager: &CircleImager{img: iw.imager, mode: mode}}
}

// Crop
func (iw *ImagerWrapper) Crop(rect image.Rectangle) *ImagerWrapper {
	return &ImagerWrapper{imager: &CropImager{img: iw.imager, rect: rect}}
}

// Flip
func (iw *ImagerWrapper) Flip(mode FlipMode) *ImagerWrapper {
	return &ImagerWrapper{imager: &FlipImager{img: iw.imager, mode: mode}}
}

// Scale
func (iw *ImagerWrapper) Scale(scale float64) *ImagerWrapper {
	return &ImagerWrapper{imager: &ScaleImager{img: iw.imager, scale: scale}}
}

// Round
func (iw *ImagerWrapper) Round(radius int) *ImagerWrapper {
	return &ImagerWrapper{imager: &RoundImager{img: iw.imager, radius: radius}}
}

// Blur
func (iw *ImagerWrapper) Blur(level int) *ImagerWrapper {
	return &ImagerWrapper{imager: &BlurImager{img: iw.imager, level: level}}
}

// Sharp
func (iw *ImagerWrapper) Sharp() *ImagerWrapper {
	return &ImagerWrapper{imager: &SharpImager{img: iw.imager}}
}

// Rotate
func (iw *ImagerWrapper) Rotate(radian float64) *ImagerWrapper {
	return &ImagerWrapper{imager: &RotateImager{img: iw.imager, radian: modRad(radian)}}
}

// ColorModel
func (iw *ImagerWrapper) ColorModel() color.Model {
	return iw.imager.ColorModel()
}

// Bounds
func (iw *ImagerWrapper) Bounds() image.Rectangle {
	return iw.imager.Bounds()
}

// At
func (iw *ImagerWrapper) At(x, y int) color.Color {
	return iw.imager.At(x, y)
}
