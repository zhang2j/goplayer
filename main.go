package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/zhang2j/goplayer/consts"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("GoPlayer")
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(consts.Width, consts.Height))

	// slider := widget.NewSlider(0, 1)
	// playButton := widget.NewButtonWithIcon("play", theme.MediaPlayIcon(), nil)

	// view := container.NewVBox(slider, playButton)

	// window.SetContent(view)
	window.ShowAndRun()
}
