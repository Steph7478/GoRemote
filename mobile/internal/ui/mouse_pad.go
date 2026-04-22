package ui

import (
	"image/color"
	"mobile/internal/client"
	"mobile/internal/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type MousePad struct {
	widget.BaseWidget
	client *client.Client
	lastX  float32
	lastY  float32
}

func NewMousePad(client *client.Client) *MousePad {
	m := &MousePad{client: client}
	m.ExtendBaseWidget(m)
	return m
}

func (m *MousePad) CreateRenderer() fyne.WidgetRenderer {
	rect := canvas.NewRectangle(color.NRGBA{R: 50, G: 50, B: 50, A: 255})
	label := canvas.NewText("TOUCHPAD", color.White)
	label.TextSize = 16
	label.Alignment = fyne.TextAlignCenter

	return &mousePadRenderer{rect: rect, label: label}
}

func (m *MousePad) Dragged(e *fyne.DragEvent) {
	dx := e.Position.X - m.lastX
	dy := e.Position.Y - m.lastY

	if dx != 0 || dy != 0 {
		m.client.Send(models.WSMessage{Event: "move", X: float64(dx), Y: float64(dy)})
	}

	m.lastX = e.Position.X
	m.lastY = e.Position.Y
}

func (m *MousePad) DragEnd() {
	m.lastX = 0
	m.lastY = 0
}

func (m *MousePad) Tapped(e *fyne.PointEvent) {
	m.client.Send(models.WSMessage{Event: "click"})
}

type mousePadRenderer struct {
	rect  *canvas.Rectangle
	label *canvas.Text
}

func (r *mousePadRenderer) Layout(size fyne.Size) {
	r.rect.Resize(size)
	r.label.Resize(fyne.NewSize(size.Width, 30))
	r.label.Move(fyne.NewPos(0, size.Height/2-15))
}

func (r *mousePadRenderer) MinSize() fyne.Size {
	return fyne.Size{Width: 0, Height: 0}
}

func (r *mousePadRenderer) Refresh() {
	r.rect.Refresh()
	r.label.Refresh()
}

func (r *mousePadRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.rect, r.label}
}

func (r *mousePadRenderer) Destroy() {}
