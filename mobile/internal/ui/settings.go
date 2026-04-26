package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateSettings(screen *MainScreen) fyne.CanvasObject {
	speedSlider := widget.NewSlider(0.1, 5.0)
	speedSlider.SetValue(screen.speedValue)
	speedSlider.Step = 0.1

	valueLabel := widget.NewLabelWithStyle(
		fmt.Sprintf("%.1f", screen.speedValue),
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)

	speedSlider.OnChanged = func(v float64) {
		screen.speedValue = v
		valueLabel.SetText(fmt.Sprintf("%.1f", v))
		fyne.CurrentApp().Preferences().SetFloat("mouseSpeed", v)
		if c := screen.GetClient(); c != nil {
			c.SetSensitivity(v)
		}
	}

	resetBtn := widget.NewButton("Reset to 1.0", func() {
		speedSlider.SetValue(1.0)
	})

	return container.NewVBox(
		widget.NewLabelWithStyle("⚙️ Mouse Settings", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewSeparator(),
		widget.NewLabel("Sensitivity:"),
		speedSlider,
		container.NewCenter(valueLabel),
		widget.NewSeparator(),
		resetBtn,
		widget.NewLabel("(Connect to apply in real-time)"),
	)
}
