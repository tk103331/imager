package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
)

type Window struct {
	app fyne.App
	win fyne.Window
	viewer *fyne.Container
	radio *widget.RadioGroup
	undo *widget.Button
	redo *widget.Button

	origin image.Image
	filters *Filters
	filter Filter
}

func Run() {
	myApp := app.NewWithID("com.github.tk103331.imager")
	(&Window{app: myApp}).Run()
}

func (w *Window) Run() {

	w.win = w.app.NewWindow("Imager")
	w.initUI()
	w.win.Resize(fyne.NewSize(800,600))
	w.win.ShowAndRun()
}

func (w *Window) initUI() {
	buttonBox := container.NewHBox(widget.NewLabel("Filter:"))
	toolbox := container.NewVBox(buttonBox, container.NewHBox(), container.NewHBox())

	for n := range AllFilters {
		func(name string){
			btn := widget.NewButton(name, func() {
				if w.origin == nil {
					return
				}
				f := AllFilters[name]()
				f.SetOnUpdate(w.refreshImg)
				w.filter = f
				w.refreshImg()



				btns := container.NewHBox(
					widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
						toolbox.Objects[1] = container.NewHBox()
						w.filter = nil
						w.refreshImg()
					}),
					widget.NewButtonWithIcon("", theme.ConfirmIcon(), func() {
						w.filters.Add(f)
						w.filter = nil
						toolbox.Objects[1] = container.NewHBox()
						w.refreshImg()
					}))
				toolbox.Objects[1] = container.NewBorder(nil, nil, nil, btns, f.Object())
			})
			buttonBox.Add(btn)
		}(n)
	}
	buttonBox.Add(layout.NewSpacer())
	w.undo = widget.NewButtonWithIcon("", theme.ContentUndoIcon(), func() {
		w.filters.Undo()
		w.refreshImg()
	})
	w.redo = widget.NewButtonWithIcon("", theme.ContentRedoIcon(), func() {
		w.filters.Redo()
		w.refreshImg()
	})
	buttonBox.Add(w.undo)
	buttonBox.Add(w.redo)
	buttonBox.Add(widget.NewButtonWithIcon("", theme.FileImageIcon(), func() {
		dlg := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.NewError(err, w.win)
				return
			}
			if reader != nil {
				defer reader.Close()
				w.loadImg(reader)
			}

		}, w.win)
		dlg.SetFilter(storage.NewExtensionFileFilter([]string{".png",".jpg",".jpeg"}))
		dlg.Show()
	}))
	buttonBox.Add(widget.NewButtonWithIcon("", theme.DocumentSaveIcon(), func() {
		dlg := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.NewError(err, w.win)
				return
			}
			if writer != nil {
				defer writer.Close()
				w.saveImg(writer)
			}

		}, w.win)
		dlg.Show()
	}))

	w.viewer = container.NewCenter(widget.NewLabel(""))
	scroll := container.NewScroll(container.NewMax(canvas.NewRectangle(color.Black), w.viewer))

	content := container.NewBorder(toolbox, nil, nil, nil, scroll)

	w.win.SetContent(content)
	w.refreshImg()
}

func (w *Window) refreshImg() {
	if w.origin != nil {
		img := w.filters.Do()
		if w.filter != nil {
			img = w.filter.Do(img)
		}
		imgObj := canvas.NewImageFromImage(img)
		imgObj.FillMode = canvas.ImageFillOriginal
		imgObj.Resize(fyne.NewSize(float32(img.Bounds().Dx()), float32(img.Bounds().Dy())))
		w.viewer.Objects[0] = imgObj
		w.viewer.Refresh()

		if w.filters.CanUndo() {
			w.undo.Enable()
		} else {
			w.undo.Disable()
		}
		if w.filters.CanRedo() {
			w.redo.Enable()
		} else {
			w.redo.Disable()
		}

	} else {
		img := canvas.NewImageFromResource(theme.FyneLogo())
		img.FillMode = canvas.ImageFillOriginal
		w.viewer.Objects[0] = img
		w.viewer.Refresh()

		w.undo.Disable()
		w.redo.Disable()
	}
}

func (w *Window) loadImg(reader fyne.URIReadCloser) {
	ext := reader.URI().Extension()
	if ext == ".png" {
		img, _ := png.Decode(reader)
		w.origin = img
		w.filters = NewFilters(w.origin, w.viewer.Refresh)
		w.filters.OnUpdate = w.refreshImg
	} else if ext == ".jpg" || ext == ".jpeg" {
		img, _ := jpeg.Decode(reader)
		w.origin = img
	}
	w.refreshImg()
}

func (w *Window) saveImg(writer fyne.URIWriteCloser) {
	img := w.filters.Do()
	ext := writer.URI().Extension()
	if ext == ".png" {
		png.Encode(writer, img)
	} else if ext == ".jpg" || ext == ".jpeg" {
		jpeg.Encode(writer, img, nil)
	}
}