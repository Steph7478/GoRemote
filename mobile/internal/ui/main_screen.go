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

	modeButtons  map[string]*widget.Button
	clickButtons map[string]*widget.Button
}

func NewMainScreen() *MainScreen {
	savedSpeed := fyne.CurrentApp().Preferences().FloatWithFallback("mouseSpeed", 1.0)

	s := &MainScreen{
		mouse:        NewMousePad(nil),
		keyboard:     NewKeyboard(nil),
		speedValue:   savedSpeed,
		modeButtons:  make(map[string]*widget.Button),
		clickButtons: make(map[string]*widget.Button),
	}

	s.initButtons()
	return s
}

func (s *MainScreen) initButtons() {
	s.modeButtons["move"] = s.createModeButton("Move", "move")
	s.modeButtons["drag"] = s.createModeButton("Select", "drag")
	s.modeButtons["scroll"] = s.createModeButton("Scroll", "scroll")

	s.clickButtons["left"] = s.createClickButton("Left Click", "left_click")
	s.clickButtons["right"] = s.createClickButton("Right Click", "right_click")

	s.connect = widget.NewButton("Connect", s.connectToServer)
	s.disconnect = widget.NewButton("Disconnect", s.disconnectFromServer)
	s.disconnect.Hide()
}

func (s *MainScreen) createModeButton(label, mode string) *widget.Button {
	btn := widget.NewButton(label, func() {
		s.mouse.SetMode(mode)
		s.highlightModeButton(mode)
	})
	btn.Disable()
	return btn
}

func (s *MainScreen) createClickButton(label, event string) *widget.Button {
	btn := widget.NewButton(label, func() {
		if s.c != nil {
			s.c.Send(models.WSMessage{Event: event})
		}
	})
	btn.Disable()
	return btn
}

func (s *MainScreen) highlightModeButton(activeMode string) {
	for mode, btn := range s.modeButtons {
		if mode == activeMode {
			btn.Importance = widget.HighImportance
		} else {
			btn.Importance = widget.MediumImportance
		}
		btn.Refresh()
	}
}

func (s *MainScreen) setButtonsEnabled(enabled bool) {
	enableFunc := func(btn *widget.Button) {
		if enabled {
			btn.Enable()
		} else {
			btn.Disable()
		}
	}

	for _, btn := range s.modeButtons {
		enableFunc(btn)
	}
	for _, btn := range s.clickButtons {
		enableFunc(btn)
	}
}

func (s *MainScreen) Build() fyne.CanvasObject {
	s.status = widget.NewLabel("🔴")

	s.speedLabel = widget.NewLabelWithStyle(fmt.Sprintf("%.1f", s.speedValue), fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	speedControls := s.buildSpeedControls()

	modeGrid := container.NewGridWithColumns(3,
		s.modeButtons["move"], s.modeButtons["drag"], s.modeButtons["scroll"])
	clickGrid := container.NewGridWithColumns(2,
		s.clickButtons["left"], s.clickButtons["right"])

	topBar := container.NewHBox(
		s.connect, s.disconnect, layout.NewSpacer(),
		speedControls, layout.NewSpacer(),
		s.status,
	)

	keyboardRow := container.NewBorder(nil, nil, nil,
		container.NewHBox(s.keyboard.BtnDelete, s.keyboard.BtnEnter),
		s.keyboard)

	content := container.NewBorder(
		container.NewVBox(keyboardRow, widget.NewSeparator(), layout.NewSpacer(), modeGrid, clickGrid),
		nil, nil, nil,
		s.mouse,
	)

	return container.NewBorder(container.NewVBox(topBar, widget.NewSeparator()), nil, nil, nil, content)
}

func (s *MainScreen) buildSpeedControls() *fyne.Container {
	adjustSpeed := func(delta float64, min, max float64) {
		newSpeed := s.speedValue + delta
		if newSpeed >= min && newSpeed <= max {
			s.speedValue = newSpeed
			s.speedLabel.SetText(fmt.Sprintf("%.1f", s.speedValue))
			fyne.CurrentApp().Preferences().SetFloat("mouseSpeed", s.speedValue)
			if s.c != nil {
				s.c.SetSensitivity(s.speedValue)
			}
		}
	}

	btnMinus := widget.NewButton("−", func() { adjustSpeed(-0.1, 0.1, 5.0) })
	btnPlus := widget.NewButton("+", func() { adjustSpeed(0.1, 0.1, 5.0) })
	btnMinus.Importance = widget.MediumImportance
	btnPlus.Importance = widget.MediumImportance

	return container.NewHBox(
		widget.NewLabel("Speed"),
		btnMinus,
		s.speedLabel,
		btnPlus,
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
			s.setButtonsEnabled(true)
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
	s.keyboard.SetText("")
	s.setButtonsEnabled(false)

	s.status.SetText("🔴")
	s.disconnect.Hide()
	s.connect.Show()
	s.connect.Enable()
}

func (s *MainScreen) GetClient() *client.Client { return s.c }
