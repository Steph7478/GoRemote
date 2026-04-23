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

	go server()

	systray.Run(tray, nil)
}

func server() {
	go func() {
		ip := utils.GetLocalIP()
		payload := []byte(fmt.Sprintf("RemoteControl:%s", ip))

		_, err := peerdiscovery.Discover(peerdiscovery.Settings{
			Payload:   payload,
			Limit:     9999,
			Delay:     time.Second * 2,
			AllowSelf: true,
		})
		if err != nil {
			fmt.Printf("Discovery error: %v\n", err)
		}
	}()

	r := gin.New()
	r.Use(gin.Recovery())
	r.SetTrustedProxies([]string{"127.0.0.1", "192.168.0.0/16"})
	routes.Setup(r)

	fmt.Println("Server running on port 8080")
	r.Run(":8080")
}

func tray() {
	ip := utils.GetLocalIP()

	iconData, err := os.ReadFile("assets/icon.ico")
	if err == nil {
		systray.SetIcon(iconData)
	}

	systray.SetTitle("Remote Control")
	systray.SetTooltip(fmt.Sprintf("Remote Control Server\nIP: %s\nPort: 8080", ip))

	ipItem := systray.AddMenuItem(fmt.Sprintf("%s", ip), "")
	ipItem.Disable()

	portItem := systray.AddMenuItem("Port: 8080", "")
	portItem.Disable()

	statusItem := systray.AddMenuItem("Running", "")
	statusItem.Disable()

	systray.AddSeparator()

	quitItem := systray.AddMenuItem("Quit", "Close server")

	<-quitItem.ClickedCh
	systray.Quit()
}
