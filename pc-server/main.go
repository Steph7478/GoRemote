package main

import (
	"fmt"
	"os"
	"server/internal/routes"
	"server/internal/utils"

	"github.com/getlantern/systray"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	
	go server()
	
	systray.Run(tray, nil)
}

func server() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.SetTrustedProxies([]string{"127.0.0.1", "192.168.0.0/16"})
	routes.Setup(r)
	r.Run(":8080")
}

func tray() {
	ip := utils.GetLocalIP()
	
	iconData, err := os.ReadFile("assets/icon.png")
	if err == nil {
		systray.SetIcon(iconData)
	}
	
	systray.SetTitle("Remote Control")
	systray.SetTooltip(fmt.Sprintf("Remote Control Server\nIP: %s\nPort: 8080", ip))
	
	ipItem := systray.AddMenuItem(fmt.Sprintf("📡 IP: %s", ip), "")
	ipItem.Disable()
	
	portItem := systray.AddMenuItem("🔌 Port: 8080", "")
	portItem.Disable()
	
	statusItem := systray.AddMenuItem("✅ Running", "")
	statusItem.Disable()
	
	systray.AddSeparator()
	
	quitItem := systray.AddMenuItem("❌ Quit", "Close server")
	
	<-quitItem.ClickedCh
	systray.Quit()
}