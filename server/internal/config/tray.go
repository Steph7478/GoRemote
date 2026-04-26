package config

import (
	"time"

	"github.com/getlantern/systray"
)

var discoveryRunning, serverRunning = true, false

func Tray() {
	iconData := GetIcon()
	if len(iconData) > 0 {
		systray.SetIcon(iconData)
	}

	systray.SetTitle("Remote Control")
	systray.SetTooltip("Port: CLOSED")

	portItem := systray.AddMenuItem(portText(), "")
	portItem.Disable()

	systray.AddSeparator()

	toggle := systray.AddMenuItem("⏸️ Pause Discovery", "")
	quit := systray.AddMenuItem("❌ Quit", "")

	startServer()
	go startDiscovery()

	go func() {
		for range time.Tick(time.Second) {
			portItem.SetTitle(portText())
			updateTooltip()
		}
	}()

	go func() {
		for {
			select {
			case <-toggle.ClickedCh:
				discoveryRunning = !discoveryRunning

				if discoveryRunning {
					toggle.SetTitle("⏸️ Pause Discovery")
				} else {
					toggle.SetTitle("▶️ Start Discovery")
				}

				updateTooltip()

			case <-quit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func portText() string {
	if serverRunning {
		return "🟢 Port: OPEN"
	}
	return "🔌 Port: CLOSED"
}

func updateTooltip() {
	if serverRunning {
		systray.SetTooltip("✅ Port: OPEN")
	} else {
		systray.SetTooltip("🔒 Port: CLOSED")
	}
}
