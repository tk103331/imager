package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/tk103331/imager"
	"image"
	"math"
)

type RotateFilter struct{
	BaseFilter
	Radian float64
}

func (f *RotateFilter) Name() string {
	return "Rotate"
}

func (f *RotateFilter) Do(i image.Image) image.Image {
	return imager.New(i).Rotate(f.Radian)
}

func (f *RotateFilter) Object() fyne.CanvasObject {
	label := widget.NewLabel("Radian: ")
	slider := widget.NewSlider(0, math.Pi)
	slider.Value = f.Radian
	slider.Resize(fyne.NewSize(200, 50))
	slider.OnChanged = func(s float64) {
		f.Radian = s
		f.Update()
	}
	return container.NewHBox(label, slider)
}
