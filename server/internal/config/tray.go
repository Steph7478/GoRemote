package config

import (
	"fmt"
	"server/internal/utils"
	"time"

	"github.com/getlantern/systray"
)

var discoveryRunning, serverRunning = true, false

func Tray() {
	ip := utils.GetLocalIP()

	iconData := GetIcon()
	if len(iconData) > 0 {
		systray.SetIcon(iconData)
	}

	systray.SetTitle("Remote Control")
	systray.SetTooltip("Remote Control - Port: CLOSED")

	ipItem := systray.AddMenuItem(fmt.Sprintf("📡 %s:8080", ip), "")
	ipItem.Disable()

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
		systray.SetTooltip("✅ Remote Control - Port: OPEN")
	} else {
		systray.SetTooltip("🔒 Remote Control - Port: CLOSED")
	}
}
