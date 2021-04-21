package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/tk103331/imager"
	"image"
)

type SharpFilter struct{
	BaseFilter

}

func (f *SharpFilter) Name() string {
	return "Sharp"
}

func (f *SharpFilter) Do(i image.Image) image.Image {
	return imager.New(i).Sharp()
}

func (f *SharpFilter) Object() fyne.CanvasObject {
	return container.NewHBox()
}
