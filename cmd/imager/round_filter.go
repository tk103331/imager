package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/tk103331/imager"
	"image"
)

type RoundFilter struct{
	BaseFilter
	Radius float64
}

func (f *RoundFilter) Name() string {
	return "Flip"
}

func (f *RoundFilter) Do(i image.Image) image.Image {
	return imager.New(i).Round(int(f.Radius))
}

func (f *RoundFilter) Object() fyne.CanvasObject {
	label := widget.NewLabel("Radius: ")
	slider := widget.NewSlider(0, 50)
	slider.Value = f.Radius
	slider.OnChanged = func(v float64) {
		f.Radius = v
		f.Update()
	}
	return container.NewHBox(label, slider)
}
