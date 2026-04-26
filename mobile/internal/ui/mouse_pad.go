package ui

import (
	"image/color"
	"time"
	"mobile/internal/models"
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
	tapTimer    *time.Timer
	tapCount    int
	isDragging  bool
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
	m.Refresh()
}

func (m *MousePad) Dragged(e *fyne.DragEvent) {
	if m.sender == nil {
		return
	}

	if m.last.X == 0 && m.last.Y == 0 {
		m.last = e.Position
		
		if m.tapCount == 2 {
			m.isDragging = true
			m.sender.Send(models.WSMessage{Event: "down"})
		}
		return
	}

	padW := float64(m.Size().Width)
	padH := float64(m.Size().Height)
	
	if padW == 0 || padH == 0 {
		return
	}

	deltaX := float64(e.Position.X - m.last.X)
	deltaY := float64(e.Position.Y - m.last.Y)
	
	scale := (m.screenW/padW + m.screenH/padH) / 2
	dx := deltaX * scale
	dy := deltaY * scale

	if m.tapCount >= 3 {
		m.sender.Send(models.WSMessage{
			Event: "scroll",
			X:     0,
			Y:     dy / 5,
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
	m.last = fyne.Position{X: 0, Y: 0}
	m.tapCount = 0
	if m.tapTimer != nil {
		m.tapTimer.Stop()
	}
}

func (m *MousePad) Tapped(e *fyne.PointEvent) {
	if m.sender == nil {
		return
	}

	m.tapCount++
	
	if m.tapTimer != nil {
		m.tapTimer.Stop()
	}
	
	m.tapTimer = time.AfterFunc(300*time.Millisecond, func() {
		if m.tapCount == 1 {
			m.sender.Send(models.WSMessage{Event: "click"})
		}
		m.tapCount = 0
	})
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
	return fyne.NewSize(200, 200)
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