package imager

import "image"
import "image/color"

type Imager struct {
	img image.Image
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
