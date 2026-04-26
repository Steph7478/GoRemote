package ui

import (
	"image/color"
	"mobile/internal/models"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type MousePad struct {
	widget.BaseWidget
	sender      Sender
	last        fyne.Position
	screenW     float64
	screenH     float64
	lastTapTime time.Time
	tapCount    int
	isDragging  bool
	isScrolling bool
	isMoving    bool
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

	if !m.isMoving {
		m.isMoving = true
		m.last = e.Position

		if time.Since(m.lastTapTime) < 300*time.Millisecond {
			if m.tapCount == 1 {
				m.isDragging = true
				m.sender.Send(models.WSMessage{Event: "down"})
			} else if m.tapCount >= 2 {
				m.isScrolling = true
			}
		}

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

	if m.isScrolling {
		m.sender.Send(models.WSMessage{
			Event: "scroll",
			X:     0,
			Y:     deltaPadY / 5.0,
		})
	} else {
		m.sender.Send(models.WSMessage{
			Event: "move",
			X:     dx,
			Y:     dy,
		})
	}

	m.last = e.Position
}

func (m *MousePad) DragEnd() {
	if m.isDragging {
		m.sender.Send(models.WSMessage{Event: "up"})
		m.isDragging = false
	}
	m.isScrolling = false
	m.isMoving = false
	m.last = fyne.Position{}
	m.tapCount = 0
}

func (m *MousePad) Tapped(e *fyne.PointEvent) {
	if m.sender == nil {
		return
	}
	now := time.Now()
	if now.Sub(m.lastTapTime) < 300*time.Millisecond {
		m.tapCount++
	} else {
		m.tapCount = 1
	}
	m.lastTapTime = now
	m.sender.Send(models.WSMessage{Event: "click"})
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
