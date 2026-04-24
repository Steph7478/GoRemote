package ui

import (
	"mobile/internal/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Keyboard struct {
	widget.Entry
	sender    Sender
	BtnDelete *widget.Button
	BtnEnter  *widget.Button
	BtnSend   *widget.Button
}

func NewKeyboard(sender Sender) *Keyboard {
	k := &Keyboard{sender: sender}
	k.ExtendBaseWidget(k)
	k.SetPlaceHolder("Type...")
	if sender == nil {
		k.Disable()
	}

	k.BtnDelete = widget.NewButton("Del", func() {
		if k.sender != nil {
			if len(k.Text) > 0 {
				k.SetText(k.Text[:len(k.Text)-1])
			}
			k.sender.Send(models.WSMessage{Event: "type", Text: "{backspace}"})
		}
	})

	k.BtnEnter = widget.NewButton("Enter", func() {
		if k.sender != nil {
			k.sender.Send(models.WSMessage{Event: "type", Text: "\n"})
		}
	})

	k.BtnSend = widget.NewButton("Send", func() {
		if k.sender != nil && k.Text != "" {
			k.sender.Send(models.WSMessage{Event: "type", Text: k.Text})
			k.SetText("")
		}
	})

	return k
}

func (k *Keyboard) TypedRune(r rune) {
	if k.sender == nil {
		return
	}
	k.SetText(k.Text + string(r))
}

func (k *Keyboard) TypedKey(ev *fyne.KeyEvent) {
	if k.sender == nil {
		return
	}
	switch ev.Name {
	case fyne.KeyBackspace:
		if len(k.Text) > 0 {
			k.SetText(k.Text[:len(k.Text)-1])
		}
		k.sender.Send(models.WSMessage{Event: "type", Text: "{backspace}"})
	case fyne.KeyReturn, fyne.KeyEnter:
		if k.Text != "" {
			k.sender.Send(models.WSMessage{Event: "type", Text: k.Text})
			k.SetText("")
		} else {
			k.sender.Send(models.WSMessage{Event: "type", Text: "\n"})
		}
	}
}
