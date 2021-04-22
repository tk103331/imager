package main

import (
	"fyne.io/fyne/v2"
	"github.com/tk103331/imager"
	"image"
	"sync"
)

var AllFilters = map[string]func() Filter {
	"Circle": func() Filter {
		return &CircleFilter{Mode: imager.CircleOuter}
	},
	"Flip": func() Filter {
		return &FlipFilter{Mode: imager.FlipHorizontal}
	},
	"Scale": func() Filter {
		return &ScaleFilter{Scale: 1}
	},
	"Round": func() Filter {
		return &RoundFilter{Radius: 0}
	},
	"Blur": func() Filter {
		return &BlurFilter{Level: 0}
	},
	"Rotate": func() Filter {
		return &RotateFilter{Radian: 0}
	},
	"Sharp": func() Filter {
		return &SharpFilter{}
	},
}

type Filter interface {
	Name() string
	Do(image.Image) image.Image
	Object() fyne.CanvasObject
	SetOnUpdate(func())
}

type BaseFilter struct {
	parent *Filters
	OnUpdate func()
}

func (f *BaseFilter) SetOnUpdate(fn func()) {
	f.OnUpdate = fn
}

func (f *BaseFilter) Update() {
	if f.OnUpdate != nil {
		f.OnUpdate()
	}
}

type Filters struct {
	image image.Image
	filters []Filter
	index int
	OnUpdate func()
	mux sync.Mutex
}

func NewFilters(img image.Image, onUpdate func()) *Filters {
	return &Filters{image: img, filters: make([]Filter, 0), index: -1, OnUpdate: onUpdate}
}

func (f *Filters) Do() image.Image {
	f.mux.Lock()
	defer f.mux.Unlock()
	im := f.image
	for i, filter := range f.filters {
		if i <= f.index{
			im = filter.Do(im)
		}
	}
	return im
}

func (f *Filters) Add(filter Filter) {
	f.mux.Lock()
	defer f.mux.Unlock()
	if len(f.filters) > 0 {
		f.filters = append(f.filters[:f.index], filter)
	} else {
		f.filters = append(f.filters, filter)
	}
	f.index = f.index + 1
}

func (f *Filters) Undo() {
	f.mux.Lock()
	defer f.mux.Unlock()
	if f.index >= 0 {
		f.index = f.index - 1
	}
}

func (f *Filters) CanUndo() bool {
	return f.index > -1
}

func (f *Filters) Redo() {
	f.mux.Lock()
	defer f.mux.Unlock()
	if f.index < len(f.filters) {
		f.index = f.index + 1
	}
}

func (f *Filters) CanRedo() bool {
	return f.index < len(f.filters)-1
}

func (f *Filters) Reset() {
	f.mux.Lock()
	defer f.mux.Unlock()
	f.filters = make([]Filter, 0)
	f.index = -1
}


