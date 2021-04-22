package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/tk103331/imager"
	"image"
)

type CircleFilter struct{
	BaseFilter
	Mode imager.CircleMode
	parent *Filters
}

func (f *CircleFilter) Name() string {
	return "Circle"
}

func (f *CircleFilter) Do(i image.Image) image.Image {
	return imager.New(i).Circle(f.Mode)
}

func (f *CircleFilter) Object() fyne.CanvasObject {

	label := widget.NewLabel("Mode: ")
	radio := widget.NewRadioGroup([]string{"Outer", "Longer", "Shorter"}, func(s string) {
		if s == "Outer" {
			f.Mode = imager.CircleOuter
		} else if s == "Longer" {
			f.Mode = imager.CircleLonger
		} else if s == "Shorter" {
			f.Mode = imager.CircleShorter
		}
		f.Update()
	})
	radio.Selected = "Outer"
	radio.Horizontal = true
	return container.NewHBox(label, radio)
}

