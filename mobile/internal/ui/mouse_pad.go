package ui

import (
	"image/color"
	"mobile/internal/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type MousePad struct {
	widget.BaseWidget
	sender  Sender
	last    fyne.Position
	screenW float64
	screenH float64
}

func NewMousePad(sender Sender) *MousePad {
	m := &MousePad{sender: sender, screenW: 1920, screenH: 1080}
	m.ExtendBaseWidget(m)
	return m
}

func (m *MousePad) SetScreenSize(w, h float64) {
	m.screenW = w
	m.screenH = h
}

func (m *MousePad) SetSender(sender Sender) {
	m.sender = sender
	m.last = fyne.Position{}
	m.Refresh()
}

func (m *MousePad) Dragged(e *fyne.DragEvent) {
	if m.sender == nil {
		return
	}
	if m.last.X == 0 && m.last.Y == 0 {
		m.last = e.Position
		return
	}

	padW := float64(m.Size().Width)
	padH := float64(m.Size().Height)

	if padW == 0 || padH == 0 {
		return
	}

	deltaPadX := float64(e.Position.X - m.last.X)
	deltaPadY := float64(e.Position.Y - m.last.Y)

	scaleX := m.screenW / padW
	scaleY := m.screenH / padH

	scale := (scaleX + scaleY) / 2

	dx := deltaPadX * scale
	dy := deltaPadY * scale

	m.sender.Send(models.WSMessage{
		Event: "move",
		X:     dx,
		Y:     dy,
	})

	m.last = e.Position
}

func (m *MousePad) DragEnd() {
	m.last = fyne.Position{}
}

func (m *MousePad) Tapped(e *fyne.PointEvent) {
	if m.sender != nil {
		m.sender.Send(models.WSMessage{Event: "click"})
	}
}

func (m *MousePad) Scrolled(e *fyne.ScrollEvent) {
	if m.sender == nil {
		return
	}
	m.sender.Send(models.WSMessage{
		Event: "scroll",
		X:     float64(e.Scrolled.DX),
		Y:     float64(e.Scrolled.DY),
	})
}

func (m *MousePad) CreateRenderer() fyne.WidgetRenderer {
	rect := canvas.NewRectangle(color.NRGBA{50, 50, 50, 255})
	return &mousePadRenderer{rect, m}
}

type mousePadRenderer struct {
	rect *canvas.Rectangle
	m    *MousePad
}

func (r *mousePadRenderer) Layout(s fyne.Size) {
	r.rect.Resize(s)
}

func (r *mousePadRenderer) MinSize() fyne.Size {
	return fyne.NewSize(300, 300)
}

func (r *mousePadRenderer) Refresh() {
	if r.m.sender == nil {
		r.rect.FillColor = color.NRGBA{80, 80, 80, 255}
	} else {
		r.rect.FillColor = color.NRGBA{50, 50, 50, 255}
	}
	r.rect.Refresh()
}

func (r *mousePadRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.rect}
}

func (r *mousePadRenderer) Destroy() {}
