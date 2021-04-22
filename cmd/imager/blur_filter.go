package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/tk103331/imager"
	"image"
)

type BlurFilter struct{
	BaseFilter
	Level int
}

func (f *BlurFilter) Name() string {
	return "Blur"
}

func (f *BlurFilter) Do(i image.Image) image.Image {
	return imager.New(i).Blur(f.Level)
}

func (f *BlurFilter) Object() fyne.CanvasObject {
	label := widget.NewLabel("Level: ")
	slider := widget.NewSlider(0, 5)
	slider.Value = float64(f.Level)
	slider.OnChanged = func(s float64) {
		f.Level = int(s)
		f.Update()
	}
	return container.NewHBox(label, slider)
}
