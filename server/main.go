package main

import (
	"fmt"
	"os"
	"server/internal/routes"
	"server/internal/utils"
	"time"

	"github.com/getlantern/systray"
	"github.com/gin-gonic/gin"
	"github.com/schollz/peerdiscovery"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	go startServer()
	systray.Run(tray, nil)
}

func startServer() {
	go startDiscovery()
	r := gin.New()
	r.Use(gin.Recovery())

	r.SetTrustedProxies([]string{"127.0.0.1", "192.168.0.0/16"})
	routes.Setup(r)

	fmt.Println("✅ Server running on port 8080")

	if err := r.Run("0.0.0.0:8080"); err != nil {
		fmt.Printf("❌ Server error: %v\n", err)
	}
}

func startDiscovery() {
	ip := utils.GetLocalIP()
	payload := []byte(fmt.Sprintf("RemoteControl:%s", ip))

	for {
		_, err := peerdiscovery.Discover(peerdiscovery.Settings{
			Payload:   payload,
			Limit:     1,
			Delay:     time.Second * 2,
			TimeLimit: 0,
			AllowSelf: false,
		})

		if err != nil {
			fmt.Printf("❌ Discovery error: %v\n", err)
			time.Sleep(time.Second * 5)
			continue
		}

		time.Sleep(time.Second * 2)
	}
}

func tray() {
	ip := utils.GetLocalIP()

	iconData, err := os.ReadFile("assets/icon.ico")
	if err == nil {
		systray.SetIcon(iconData)
	}

	systray.SetTitle("Remote Control")
	systray.SetTooltip(fmt.Sprintf("Remote Control Server\nIP: %s\nPort: 8080", ip))

	ipItem := systray.AddMenuItem(fmt.Sprintf("📡 IP: %s", ip), "Server IP address")
	ipItem.Disable()

	portItem := systray.AddMenuItem("🔌 Port: 8080", "Server port")
	portItem.Disable()

	statusItem := systray.AddMenuItem("🟢 Running", "Server status")
	statusItem.Disable()

	systray.AddSeparator()

	quitItem := systray.AddMenuItem("❌ Quit", "Close server")

	<-quitItem.ClickedCh
	fmt.Println("🛑 Shutting down...")
	systray.Quit()
}
