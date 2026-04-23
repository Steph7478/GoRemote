package ui

import (
	"encoding/json"
	"mobile/internal/client"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type SettingsData struct {
	MouseSpeed float64 `json:"mouse_speed"`
}

func loadSettings() SettingsData {
	var data SettingsData
	file, err := os.Open("settings.json")
	if err != nil {
		return SettingsData{MouseSpeed: 1.0}
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&data)
	return data
}

func saveSettings(data SettingsData) {
	file, _ := os.Create("settings.json")
	defer file.Close()
	json.NewEncoder(file).Encode(data)
}

func CreateSettings(c *client.Client) *fyne.Container {
	settings := loadSettings()

	slider := widget.NewSlider(0.5, 3.0)
	slider.SetValue(settings.MouseSpeed)
	slider.Step = 0.1

	speedLabel := widget.NewLabel("Speed: " + formatFloat(settings.MouseSpeed))

	slider.OnChanged = func(v float64) {
		speedLabel.SetText("Speed: " + formatFloat(v))
		saveSettings(SettingsData{MouseSpeed: v})
		if c != nil {
			c.SetSensitivity(v)
		}
	}

	return container.NewVBox(
		widget.NewLabel("⚙️ Settings"),
		widget.NewSeparator(),
		speedLabel,
		slider,
		widget.NewButton("Reset", func() { slider.SetValue(1.0) }),
	)
}

func formatFloat(v float64) string {
	return strconv.FormatFloat(v, 'f', 2, 64)
}
