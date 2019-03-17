package imager

import "image"
import "image/color"

type Imager struct {
	img image.Image
}

func New(image image.Image) *Imager {
	return &Imager{image}
}

func (i *Imager) Circle(mode CircleMode) *CircleImager {
	return &CircleImager{Imager: i, mode: mode}
}

func (i *Imager) Crop(rect image.Rectangle) *CropImager {
	return &CropImager{Imager: i, rect: rect}
}

func (i *Imager) Flip(mode FlipMode) *FlipImager {
	return &FlipImager{Imager: i, mode: mode}
}

func (i *Imager) ColorModel() color.Model {
	return i.img.ColorModel()
}

func (i *Imager) Bounds() image.Rectangle {
	return i.img.Bounds()
}

func (i *Imager) At(x, y int) color.Color {
	return i.img.At(x, y)
}
