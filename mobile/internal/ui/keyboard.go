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
}

func NewKeyboard(sender Sender) *Keyboard {
	k := &Keyboard{sender: sender}
	k.ExtendBaseWidget(k)
	k.SetPlaceHolder("Type...")

	if sender == nil {
		k.Disable()
	}

	k.BtnDelete = widget.NewButton("Del", func() {
		k.backspace()
	})

	k.BtnEnter = widget.NewButton("Enter", func() {
		k.sendText()
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
		k.backspace()

	case fyne.KeyReturn, fyne.KeyEnter:
		k.sendText()
	}
}

func (k *Keyboard) backspace() {
	if k.sender == nil || k.Text == "" {
		return
	}

	runes := []rune(k.Text)
	k.SetText(string(runes[:len(runes)-1]))
}

func (k *Keyboard) sendText() {
	if k.sender == nil || k.Text == "" {
		return
	}

	k.sender.Send(models.WSMessage{
		Event: "type",
		Text:  k.Text,
	})

	k.SetText("")
}
