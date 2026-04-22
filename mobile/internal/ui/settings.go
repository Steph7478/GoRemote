package ui

import (
	"mobile/internal/client"
	"strconv"

	"encoding/json"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type SettingsData struct {
	MouseSpeed float64 `json:"mouse_speed"`
	LastIP     string  `json:"last_ip"`
}

func loadSettings() SettingsData {
	var data SettingsData
	file, err := os.Open("settings.json")
	if err != nil {
		return SettingsData{MouseSpeed: 1.0, LastIP: ""}
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
	data := loadSettings()

	slider := widget.NewSlider(0.5, 3.0)
	slider.SetValue(data.MouseSpeed)
	slider.Step = 0.1
	label := widget.NewLabel("Speed: " + strconv.FormatFloat(data.MouseSpeed, 'f', 2, 64))

	slider.OnChanged = func(v float64) {
		label.SetText("Speed: " + strconv.FormatFloat(v, 'f', 2, 64))
		settings := loadSettings()
		settings.MouseSpeed = v
		saveSettings(settings)
		if c != nil {
			c.SetSensitivity(v)
		}
	}

	resetBtn := widget.NewButton("Reset", func() {
		slider.SetValue(1.0)
	})

	return container.NewVBox(
		widget.NewLabel("⚙️ Settings"),
		widget.NewSeparator(),
		label, slider, resetBtn,
	)
}
