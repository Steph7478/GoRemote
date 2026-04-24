package ui

import (
	"fmt"
	"mobile/internal/client"
	"mobile/internal/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Sender interface{ Send(models.WSMessage) error }

type MainScreen struct {
	c          *client.Client
	status     *widget.Label
	connect    *widget.Button
	disconnect *widget.Button
	mouse      *MousePad
	keyboard   *Keyboard
	speedValue float64
	speedLabel *widget.Label
}

func NewMainScreen() *MainScreen {
	return &MainScreen{
		mouse:      NewMousePad(nil),
		keyboard:   NewKeyboard(nil),
		speedValue: 1.0,
	}
}

func (s *MainScreen) Build() fyne.CanvasObject {
	s.status = widget.NewLabel("🔴")
	s.connect = widget.NewButton("Connect", s.connectToServer)
	s.disconnect = widget.NewButton("Disconnect", s.disconnectFromServer)
	s.disconnect.Hide()

	s.speedLabel = widget.NewLabelWithStyle(
		fmt.Sprintf("%.1f", s.speedValue),
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)

	btnMinus := widget.NewButton("−", func() {
		if s.speedValue > 0.1 {
			s.speedValue -= 0.1
			s.speedLabel.SetText(fmt.Sprintf("%.1f", s.speedValue))
			if s.c != nil {
				s.c.SetSensitivity(s.speedValue)
			}
		}
	})
	btnMinus.Importance = widget.MediumImportance

	btnPlus := widget.NewButton("+", func() {
		if s.speedValue < 5.0 {
			s.speedValue += 0.1
			s.speedLabel.SetText(fmt.Sprintf("%.1f", s.speedValue))
			if s.c != nil {
				s.c.SetSensitivity(s.speedValue)
			}
		}
	})
	btnPlus.Importance = widget.MediumImportance

	speedContainer := container.NewHBox(
		widget.NewLabelWithStyle("Speed", fyne.TextAlignCenter, fyne.TextStyle{}),
		btnMinus,
		s.speedLabel,
		btnPlus,
	)

	topBar := container.NewHBox(
		s.connect,
		s.disconnect,
		layout.NewSpacer(),
		speedContainer,
		layout.NewSpacer(),
		s.status,
	)

	keyboardRow := container.NewBorder(
		nil, nil,
		container.NewHBox(s.keyboard.BtnDelete, s.keyboard.BtnEnter),
		container.NewHBox(s.keyboard.BtnSend),
		s.keyboard,
	)

	content := container.NewBorder(
		container.NewVBox(
			widget.NewLabelWithStyle("⌨️ KEYBOARD", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
			keyboardRow,
			widget.NewSeparator(),
			widget.NewLabelWithStyle("🖱️ MOUSE PAD", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		),
		nil, nil, nil,
		s.mouse,
	)

	return container.NewBorder(
		container.NewVBox(topBar, widget.NewSeparator()),
		nil, nil, nil,
		content,
	)
}

func (s *MainScreen) connectToServer() {
	s.status.SetText("⏳")
	s.connect.Disable()

	go func() {
		ip, _ := client.DiscoverServer()
		fyne.Do(func() {
			if ip == "" {
				s.status.SetText("❌")
				s.connect.Enable()
				return
			}
			c := client.NewClient(ip)
			if err := c.Connect(); err != nil {
				s.status.SetText("❌")
				s.connect.Enable()
				return
			}
			s.c = c
			s.mouse.sender = c
			s.keyboard.sender = c
			s.keyboard.Enable()
			c.SetSensitivity(s.speedValue)
			s.status.SetText("✅")
			s.connect.Hide()
			s.disconnect.Show()
		})
	}()
}

func (s *MainScreen) disconnectFromServer() {
	if s.c != nil {
		s.c.Close()
	}
	s.mouse.sender = nil
	s.keyboard.sender = nil
	s.keyboard.Disable()
	s.keyboard.SetText("")
	s.status.SetText("🔴")
	s.disconnect.Hide()
	s.connect.Show()
	s.connect.Enable()
}

func (s *MainScreen) GetClient() *client.Client { return s.c }
