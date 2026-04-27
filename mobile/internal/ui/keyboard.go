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

	k.BtnDelete = widget.NewButton("Del", func() {
		k.key("backspace")
	})
	k.BtnDelete.Disable()

	k.BtnEnter = widget.NewButton("Enter", func() {
		k.enter()
	})
	k.BtnEnter.Disable()

	if sender != nil {
		k.Enable()
	} else {
		k.Disable()
	}

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
	if k.sender == nil || k.Disabled() {
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
	if k.sender == nil || k.Disabled() {
		return
	}
	k.sender.Send(models.WSMessage{
		Event: "key",
		Key:   key,
	})
}
func (k *Keyboard) Disable() {
	k.Entry.Disable()
	if k.BtnDelete != nil {
		k.BtnDelete.Disable()
	}
	if k.BtnEnter != nil {
		k.BtnEnter.Disable()
	}
}

func (k *Keyboard) Enable() {
	k.Entry.Enable()
	if k.sender != nil && k.BtnDelete != nil && k.BtnEnter != nil {
		k.BtnDelete.Enable()
		k.BtnEnter.Enable()
	}
}