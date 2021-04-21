package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/tk103331/imager"
	"image"
)

type FlipFilter struct{
	BaseFilter
	Mode imager.FlipMode
}

func (f *FlipFilter) Name() string {
	return "Flip"
}

func (f *FlipFilter) Do(i image.Image) image.Image {
	return imager.New(i).Flip(f.Mode)
}

func (f *FlipFilter) Object() fyne.CanvasObject {

	label := widget.NewLabel("Mode: ")
	radio := widget.NewRadioGroup([]string{"Horizontal", "Vertical"}, func(s string) {
		if s == "Horizontal" {
			f.Mode = imager.FlipHorizontal
		} else if s == "Vertical" {
			f.Mode = imager.FlipVertical
		}
		f.Update()
	})
	radio.Horizontal = true
	return container.NewHBox(label, radio)
}


