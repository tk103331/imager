package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/tk103331/imager"
	"image"
)

type ScaleFilter struct{
	BaseFilter
	Scale float64
}

func (f *ScaleFilter) Name() string {
	return "Scale"
}

func (f *ScaleFilter) Do(i image.Image) image.Image {
	return imager.New(i).Scale(f.Scale)
}

func (f *ScaleFilter) Object() fyne.CanvasObject {
	label := widget.NewLabel("Scale: ")
	slider := widget.NewSlider(0, 5)
	slider.Value = f.Scale
	slider.OnChanged = func(s float64) {
		f.Scale = s
		f.Update()
	}
	return container.NewBorder(nil, nil, label, nil, slider)
}
