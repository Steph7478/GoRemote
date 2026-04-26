package main

import (
	"mobile/internal/ui"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Remote Control")

	icon, _ := fyne.LoadResourceFromPath("assets/icon.ico")
	if icon != nil {
		w.SetIcon(icon)
	}

	if runtime.GOOS == "android" || runtime.GOOS == "ios" {
		w.SetFullScreen(true)
	} else {
		w.Resize(fyne.NewSize(400, 650))
	}

	w.SetContent(ui.NewMainScreen().Build())
	w.ShowAndRun()
}
