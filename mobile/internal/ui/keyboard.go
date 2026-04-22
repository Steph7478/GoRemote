package ui

import (
	"mobile/internal/client"
	"mobile/internal/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateKeyboard(client *client.Client) *fyne.Container {
	input := widget.NewEntry()
	input.SetPlaceHolder("Type and press Enter")
	
	input.OnSubmitted = func(t string) {
		if t != "" {
			client.Send(models.WSMessage{Event: "type", Text: t + "\n"})
			input.SetText("")
		}
	}
	
	sendBtn := widget.NewButton("Send", func() {
		if input.Text != "" {
			client.Send(models.WSMessage{Event: "type", Text: input.Text + "\n"})
			input.SetText("")
		}
	})
	
	return container.NewVBox(
		widget.NewLabel("⌨️ Keyboard"),
		input,
		sendBtn,
	)
}