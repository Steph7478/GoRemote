package main

import (
	"mobile/internal/ui"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type resizeListener struct {
	window fyne.Window
}

func (r *resizeListener) Resize(size fyne.Size) {
	r.window.Canvas().Refresh(r.window.Content())
}

func main() {
	a := app.New()
	w := a.NewWindow("Remote Control")

	icon, err := fyne.LoadResourceFromPath("assets/icon.png")
	if err == nil {
		w.SetIcon(icon)
	}

	if runtime.GOOS == "android" || runtime.GOOS == "ios" {
		w.SetFullScreen(true)
	} else {
		w.Resize(fyne.NewSize(400, 650))
	}

	screen := ui.NewMainScreen(w)
	w.SetContent(screen.Build())
	w.ShowAndRun()
}
