package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateSettings(screen *MainScreen) fyne.CanvasObject {
	speed := widget.NewSlider(0.1, 3.0)
	speed.SetValue(1.0)
	speed.OnChanged = func(v float64) {
		if c := screen.GetClient(); c != nil {
			c.SetSensitivity(v)
		}
	}
	return container.NewVBox(
		widget.NewLabel("⚙️ Mouse Speed"), speed,
		widget.NewLabel("(connect to adjust)"),
	)
}
