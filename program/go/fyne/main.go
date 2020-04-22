package main

import (
	"fmt"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {

	// Fyneを起動
	app := app.New()

	// windowを作成
	w := app.NewWindow("Timer App")
	w.Resize(fyne.NewSize(320, 480))
	w.CenterOnScreen()

	tabs := makeWelcome()

	w.SetContent(tabs)

	w.ShowAndRun()
}

func makeWelcome() *widget.Box {
	var stopChan chan struct{}

	label := widget.NewLabelWithStyle("0 : 00 : 00", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	//Buttons
	btnStart := widget.NewButton("Start", func() {
		if stopChan != nil {
			close(stopChan)
		}
		stopChan = make(chan struct{})
		go timer(label, stopChan)
	})
	btnStop := widget.NewButton("Stop", func() {
		if stopChan != nil {
			close(stopChan)
		}
		stopChan = make(chan struct{})
	})
	btn := fyne.NewContainerWithLayout(layout.NewGridLayout(2), btnStart, btnStop)

	return widget.NewVBox(label, btn)
}

func timer(label *widget.Label, stopCh chan struct{}) {
	var s int
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-stopCh:
			return
		case <-ticker.C:
			s++
			label.SetText(fmt.Sprintf("%d : %02d : %02d", s/3600, (s%3600)/60, s%60))
		}
	}
}
