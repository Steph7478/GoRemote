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
	moveBtn    *widget.Button
	dragBtn    *widget.Button
	scrollBtn  *widget.Button
	leftBtn    *widget.Button
	rightBtn   *widget.Button
}

func NewMainScreen() *MainScreen {
	savedSpeed := fyne.CurrentApp().Preferences().FloatWithFallback("mouseSpeed", 1.0)
	return &MainScreen{
		mouse:      NewMousePad(nil),
		keyboard:   NewKeyboard(nil),
		speedValue: savedSpeed,
	}
}

func (s *MainScreen) Build() fyne.CanvasObject {
	s.status = widget.NewLabel("🔴")
	s.connect = widget.NewButton("Connect", s.connectToServer)
	s.disconnect = widget.NewButton("Disconnect", s.disconnectFromServer)
	s.disconnect.Hide()

	s.moveBtn = widget.NewButton("🖱️ Move", func() {
		s.mouse.SetMode("move")
		s.updateButtons("move")
	})
	s.moveBtn.Importance = widget.HighImportance
	s.moveBtn.Disable()

	s.dragBtn = widget.NewButton("👆 Select", func() {
		s.mouse.SetMode("drag")
		s.updateButtons("drag")
	})
	s.dragBtn.Disable()

	s.scrollBtn = widget.NewButton("📜 Scroll", func() {
		s.mouse.SetMode("scroll")
		s.updateButtons("scroll")
	})
	s.scrollBtn.Disable()

	s.leftBtn = widget.NewButton("Left Click", func() {
		if s.c != nil {
			s.c.Send(models.WSMessage{Event: "left_click"})
		}
	})
	s.leftBtn.Disable()

	s.rightBtn = widget.NewButton("Right Click", func() {
		if s.c != nil {
			s.c.Send(models.WSMessage{Event: "right_click"})
		}
	})
	s.rightBtn.Disable()

	s.updateButtons("move")

	modeButtons := container.NewGridWithColumns(3, s.moveBtn, s.dragBtn, s.scrollBtn)
	clickButtons := container.NewGridWithColumns(2, s.leftBtn, s.rightBtn)

	s.speedLabel = widget.NewLabelWithStyle(fmt.Sprintf("%.1f", s.speedValue), fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	btnMinus := widget.NewButton("−", func() {
		if s.speedValue > 0.1 {
			s.speedValue -= 0.1
			s.speedLabel.SetText(fmt.Sprintf("%.1f", s.speedValue))
			fyne.CurrentApp().Preferences().SetFloat("mouseSpeed", s.speedValue)
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
			fyne.CurrentApp().Preferences().SetFloat("mouseSpeed", s.speedValue)
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

	topBar := container.NewHBox(s.connect, s.disconnect, layout.NewSpacer(), speedContainer, layout.NewSpacer(), s.status)

	keyboardRow := container.NewBorder(nil, nil, nil, container.NewHBox(s.keyboard.BtnDelete, s.keyboard.BtnEnter), s.keyboard)

	content := container.NewBorder(
		container.NewVBox(keyboardRow, widget.NewSeparator(), modeButtons, clickButtons),
		nil, nil, nil,
		s.mouse,
	)

	return container.NewBorder(container.NewVBox(topBar, widget.NewSeparator()), nil, nil, nil, content)
}

func (s *MainScreen) updateButtons(mode string) {
	if s.moveBtn == nil || s.dragBtn == nil || s.scrollBtn == nil {
		return
	}
	switch mode {
	case "move":
		s.moveBtn.Importance = widget.HighImportance
		s.dragBtn.Importance = widget.MediumImportance
		s.scrollBtn.Importance = widget.MediumImportance
	case "drag":
		s.moveBtn.Importance = widget.MediumImportance
		s.dragBtn.Importance = widget.HighImportance
		s.scrollBtn.Importance = widget.MediumImportance
	case "scroll":
		s.moveBtn.Importance = widget.MediumImportance
		s.dragBtn.Importance = widget.MediumImportance
		s.scrollBtn.Importance = widget.HighImportance
	}
	s.moveBtn.Refresh()
	s.dragBtn.Refresh()
	s.scrollBtn.Refresh()
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
			s.leftBtn.Enable()
			s.rightBtn.Enable()
			s.moveBtn.Enable()
			s.dragBtn.Enable()
			s.scrollBtn.Enable()
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
		s.c = nil
	}
	s.mouse.sender = nil
	s.keyboard.sender = nil
	s.keyboard.Disable()
	s.leftBtn.Disable()
	s.rightBtn.Disable()
	s.moveBtn.Disable()
	s.dragBtn.Disable()
	s.scrollBtn.Disable()
	s.keyboard.SetText("")
	s.status.SetText("🔴")
	s.disconnect.Hide()
	s.connect.Show()
	s.connect.Enable()
}

func (s *MainScreen) GetClient() *client.Client { return s.c }