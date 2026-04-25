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
	k.SetPlaceHolder("Type here, send to PC")

	if sender == nil {
		k.Disable()
	}

	k.BtnDelete = widget.NewButton("Del", func() {
		k.key("backspace")
	})

	k.BtnEnter = widget.NewButton("Enter", func() {
		k.enter()
	})

	return k
}

func (k *Keyboard) TypedKey(ev *fyne.KeyEvent) {
	if k.sender == nil {
		return
	}

	switch ev.Name {
	case fyne.KeyReturn, fyne.KeyEnter:
		k.enter()

	default:
		k.Entry.TypedKey(ev)
	}
}

func (k *Keyboard) enter() {
	if k.sender == nil {
		return
	}

	if k.Text != "" {
		k.sender.Send(models.WSMessage{
			Event: "type",
			Text:  k.Text,
		})
		k.SetText("")
		return
	}

	k.key("enter")
}

func (k *Keyboard) key(key string) {
	if k.sender == nil {
		return
	}

	k.sender.Send(models.WSMessage{
		Event: "key",
		Key:   key,
	})
}
