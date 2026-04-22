package ui

import (
	"mobile/internal/client"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type MainScreen struct {
	window        fyne.Window
	client        *client.Client
	status        *widget.Label
	ipEntry       *widget.Entry
	connectBtn    *widget.Button
	disconnectBtn *widget.Button
	tabs          *container.AppTabs
}

func NewMainScreen(window fyne.Window) *MainScreen {
	return &MainScreen{window: window}
}

func (s *MainScreen) Build() fyne.CanvasObject {
	s.ipEntry = widget.NewEntry()
	s.ipEntry.SetPlaceHolder("PC IP (ex: 192.168.18.32)")
	s.status = widget.NewLabel("🔴 Disconnected")

	s.connectBtn = widget.NewButton("Connect", s.onConnect)
	s.disconnectBtn = widget.NewButton("Disconnect", s.onDisconnect)
	s.disconnectBtn.Hide()

	s.tabs = container.NewAppTabs(
		container.NewTabItem("🖱️ Mouse", s.buildMouse()),
		container.NewTabItem("⌨️ Keyboard", s.buildKeyboard()),
		container.NewTabItem("⚙️ Settings", CreateSettings(s.client)),
	)

	ipRow := container.NewBorder(nil, nil, nil, container.NewHBox(s.connectBtn, s.disconnectBtn), s.ipEntry)
	topBar := container.NewVBox(ipRow, s.status, widget.NewSeparator())

	return container.NewBorder(topBar, nil, nil, nil, s.tabs)
}

func (s *MainScreen) buildMouse() fyne.CanvasObject {
	if s.client == nil {
		return container.NewCenter(widget.NewLabel("Connect first"))
	}
	return container.NewStack(NewMousePad(s.client))
}

func (s *MainScreen) buildKeyboard() fyne.CanvasObject {
	if s.client == nil {
		return container.NewCenter(widget.NewLabel("Connect first"))
	}
	return CreateKeyboard(s.client)
}

func (s *MainScreen) refreshTabs() {
	s.tabs.Items[0].Content = s.buildMouse()
	s.tabs.Items[1].Content = s.buildKeyboard()
	s.tabs.Refresh()
}

func (s *MainScreen) onConnect() {
	if s.client != nil {
		s.client.Close()
	}
	s.client = client.NewClient(s.ipEntry.Text)
	if err := s.client.Connect(); err != nil {
		s.status.SetText("❌ Failed")
		return
	}

	if data := loadSettings(); data.MouseSpeed != 1.0 {
		s.client.SetSensitivity(data.MouseSpeed)
	}

	s.status.SetText("✅ Connected")
	s.connectBtn.Hide()
	s.disconnectBtn.Show()
	s.refreshTabs()
}

func (s *MainScreen) onDisconnect() {
	if s.client != nil {
		s.client.Close()
		s.client = nil
	}
	s.status.SetText("🔴 Disconnected")
	s.disconnectBtn.Hide()
	s.connectBtn.Show()
	s.refreshTabs()
}
