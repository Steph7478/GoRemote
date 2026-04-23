package ui

import (
	"mobile/internal/client"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type MainScreen struct {
	window fyne.Window
	client *client.Client
	ip     *widget.Entry
	status *widget.Label
	btns   struct {
		connect    *widget.Button
		disconnect *widget.Button
	}
	tabs *container.AppTabs
}

func NewMainScreen(window fyne.Window) *MainScreen {
	return &MainScreen{window: window}
}

func (s *MainScreen) Build() fyne.CanvasObject {
	s.ip = widget.NewEntry()
	s.ip.SetPlaceHolder("PC IP (ex: 192.168.18.32)")

	if data := loadSettings(); data.LastIP != "" {
		s.ip.SetText(data.LastIP)
	}

	s.ip.OnChanged = func(ip string) {
		if ip != "" {
			settings := loadSettings()
			settings.LastIP = ip
			saveSettings(settings)
		}
	}

	s.status = widget.NewLabel("🔴 Disconnected")

	s.btns.connect = widget.NewButton("Connect", s.onConnect)
	s.btns.disconnect = widget.NewButton("Disconnect", s.onDisconnect)
	s.btns.disconnect.Hide()

	s.tabs = container.NewAppTabs(
		container.NewTabItem("🖱️ Mouse", s.buildMouse()),
		container.NewTabItem("⌨️ Keyboard", s.buildKeyboard()),
		container.NewTabItem("⚙️ Settings", CreateSettings(nil)),
	)

	topBar := container.NewVBox(
		container.NewBorder(nil, nil, nil,
			container.NewHBox(s.btns.connect, s.btns.disconnect),
			s.ip),
		s.status,
		widget.NewSeparator(),
	)

	return container.NewBorder(topBar, nil, nil, nil, s.tabs)
}

func (s *MainScreen) buildMouse() fyne.CanvasObject {
	if s.client == nil {
		return container.NewCenter(widget.NewLabel("Connect first"))
	}
	return container.NewMax(NewMousePad(s.client))
}

func (s *MainScreen) buildKeyboard() fyne.CanvasObject {
	if s.client == nil {
		return container.NewCenter(widget.NewLabel("Connect first"))
	}
	return CreateKeyboard(s.client)
}

func (s *MainScreen) onConnect() {
	s.closeClient()

	ip := s.ip.Text
	s.client = client.NewClient(ip)

	if err := s.client.Connect(); err != nil {
		s.status.SetText("❌ Failed")
		return
	}

	if settings := loadSettings(); settings.MouseSpeed != 1.0 {
		s.client.SetSensitivity(settings.MouseSpeed)
	}

	s.status.SetText("✅ Connected")
	s.btns.connect.Hide()
	s.btns.disconnect.Show()
	s.updateTabs()
}

func (s *MainScreen) onDisconnect() {
	s.closeClient()
	s.status.SetText("🔴 Disconnected")
	s.btns.disconnect.Hide()
	s.btns.connect.Show()
	s.updateTabs()
}

func (s *MainScreen) closeClient() {
	if s.client != nil {
		s.client.Close()
		s.client = nil
	}
}

func (s *MainScreen) updateTabs() {
	s.tabs.Items[0].Content = s.buildMouse()
	s.tabs.Items[1].Content = s.buildKeyboard()
	s.tabs.Refresh()
}
