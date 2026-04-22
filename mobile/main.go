package main

import (
	"mobile/internal/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Remote Control")
	w.Resize(fyne.NewSize(400, 650))

	screen := ui.NewMainScreen(w)
	w.SetContent(screen.Build())

	w.ShowAndRun()
}