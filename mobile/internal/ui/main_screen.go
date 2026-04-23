package ui

import (
	"mobile/internal/client"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type MainScreen struct {
	client     *client.Client
	status     *widget.Label
	connect    *widget.Button
	disconnect *widget.Button
	search     *widget.Button
	tabs       *container.AppTabs
}

func NewMainScreen(_ fyne.Window) *MainScreen { return &MainScreen{} }

func (s *MainScreen) Build() fyne.CanvasObject {
	s.status = widget.NewLabel("🔴 Disconnected")
	s.connect = widget.NewButton("Connect", s.onConnect)
	s.disconnect = widget.NewButton("Disconnect", s.onDisconnect)
	s.search = widget.NewButton("🔍 Search PC", s.onSearch)
	s.connect.Hide()
	s.disconnect.Hide()

	s.tabs = container.NewAppTabs(
		container.NewTabItem("🖱️ Mouse", s.buildMouse()),
		container.NewTabItem("⌨️ Keyboard", s.buildKeyboard()),
		container.NewTabItem("⚙️ Settings", CreateSettings(nil)),
	)

	return container.NewBorder(
		container.NewVBox(
			container.NewHBox(s.search, s.connect, s.disconnect, layout.NewSpacer(), s.status),
			widget.NewSeparator(),
		), nil, nil, nil, s.tabs)
}

func (s *MainScreen) onSearch() {
	s.status.SetText("⏳ Searching...")
	s.search.Disable()
	go func() {
		ip, _ := client.DiscoverServer()
		fyne.Do(func() {
			if ip == "" {
				s.status.SetText("❌ Not found")
				s.search.Enable()
				return
			}
			s.status.SetText("✅ Connecting...")
			s.search.Enable()
			s.connect.Show()
			s.client = client.NewClient(ip)
			if s.client.Connect() != nil {
				s.status.SetText("❌ Failed")
				s.connect.Hide()
				return
			}
			settings := loadSettings()
			if settings.MouseSpeed != 1.0 {
				s.client.SetSensitivity(settings.MouseSpeed)
			}
			s.status.SetText("✅ Connected")
			s.connect.Hide()
			s.disconnect.Show()
			s.search.Hide()
			s.updateTabs()
		})
	}()
}

func (s *MainScreen) buildMouse() fyne.CanvasObject {
	if s.client == nil {
		return container.NewCenter(widget.NewLabel("Click Search PC"))
	}
	return container.NewStack(NewMousePad(s.client))
}

func (s *MainScreen) buildKeyboard() fyne.CanvasObject {
	if s.client == nil {
		return container.NewCenter(widget.NewLabel("Click Search PC"))
	}
	return CreateKeyboard(s.client)
}

func (s *MainScreen) onConnect() { s.onSearch() }
func (s *MainScreen) onDisconnect() {
	if s.client != nil {
		s.client.Close()
		s.client = nil
	}
	s.status.SetText("🔴 Disconnected")
	s.disconnect.Hide()
	s.connect.Hide()
	s.search.Show()
	s.updateTabs()
}

func (s *MainScreen) updateTabs() {
	s.tabs.Items[0].Content = s.buildMouse()
	s.tabs.Items[1].Content = s.buildKeyboard()
	s.tabs.Refresh()
}
