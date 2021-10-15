package main

import (
	"flag"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

var data = []string{"a", "string", "list"}

type doubleTappableLabel struct {
	widget.Label
}

func newDoubleTappableLabel(payload string) *doubleTappableLabel {
	w := &doubleTappableLabel{}
	w.ExtendBaseWidget(w)
	w.SetText(payload)
	return w
}

func (d *doubleTappableLabel) DoubleTapped(_ *fyne.PointEvent) {
	log.Printf("DoubleTapped: %v", d.Text)
}

func main() {
	var useDoubleTappableLabel = flag.Bool("d", false, "use double tappable list items")
	flag.Parse()
	log.Printf("useDoubleTappableLabel: %v", *useDoubleTappableLabel)
	myApp := app.New()
	myWindow := myApp.NewWindow("List Widget")

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			if *useDoubleTappableLabel {
				return newDoubleTappableLabel("template")
			} else {
				return widget.NewLabel("template")
			}
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			if *useDoubleTappableLabel {
				o.(*doubleTappableLabel).SetText(data[i])
			} else {
				o.(*widget.Label).SetText(data[i])
			}
		})
	list.OnSelected = func(i widget.ListItemID) {
		log.Printf("Selected %v", i)
	}
	myWindow.SetContent(list)
	myWindow.Resize(fyne.Size{Height: 320, Width: 480})
	myWindow.ShowAndRun()
}
