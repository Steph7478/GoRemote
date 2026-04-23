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
	client   *client.Client
	last     fyne.Position
	dragging bool
}

func NewMousePad(client *client.Client) *MousePad {
	m := &MousePad{client: client}
	m.ExtendBaseWidget(m)
	return m
}

func (m *MousePad) CreateRenderer() fyne.WidgetRenderer {
	return &mousePadRenderer{
		rect:  canvas.NewRectangle(color.NRGBA{R: 50, G: 50, B: 50, A: 255}),
		label: canvas.NewText("TOUCHPAD", color.White),
	}
}

func (m *MousePad) Dragged(e *fyne.DragEvent) {
	if !m.dragging {
		m.dragging = true
		m.last = e.Position
		return
	}

	dx := e.Position.X - m.last.X
	dy := e.Position.Y - m.last.Y

	if dx != 0 || dy != 0 {
		m.client.Send(models.WSMessage{
			Event: "move",
			X:     float64(dx),
			Y:     float64(dy),
		})
	}

	m.last = e.Position
}

func (m *MousePad) DragEnd() {
	m.dragging = false
	m.last = fyne.Position{}
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

	labelSize := r.label.MinSize()
	r.label.Resize(labelSize)
	r.label.Move(fyne.NewPos(
		(size.Width-labelSize.Width)/2,
		(size.Height-labelSize.Height)/2,
	))
}

func (r *mousePadRenderer) MinSize() fyne.Size {
	return fyne.NewSize(100, 100)
}

func (r *mousePadRenderer) Refresh() {
	r.rect.Refresh()
	r.label.Refresh()
}

func (r *mousePadRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.rect, r.label}
}

func (r *mousePadRenderer) Destroy() {}