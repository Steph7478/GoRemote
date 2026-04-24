package main

import (
	"fmt"
	"net"
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
	r.SetTrustedProxies([]string{"127.0.0.1", "192.168.0.0/16", "10.0.0.0/8"})
	routes.Setup(r)
	r.Run("0.0.0.0:8080")
}

func startDiscovery() {
	ip := utils.GetLocalIP()
	payload := []byte(fmt.Sprintf("RemoteControl:%s", ip))
	for {
		peerdiscovery.Discover(peerdiscovery.Settings{Payload: payload, Limit: 1, Delay: time.Second * 2, AllowSelf: false})
		time.Sleep(time.Second * 2)
	}
}

func tray() {
	ip := utils.GetLocalIP()
	if icon, _ := os.ReadFile("assets/icon.ico"); icon != nil {
		systray.SetIcon(icon)
	}
	systray.SetTitle("Remote Control")
	systray.SetTooltip(fmt.Sprintf("%s:8080", ip))

	makeItem := func(t string, d bool) *systray.MenuItem {
		item := systray.AddMenuItem(t, "")
		if d {
			item.Disable()
		}
		return item
	}

	makeItem(fmt.Sprintf("📡 %s:8080", ip), true)
	makeItem("🔌 Port: 8080", true)
	makeItem("🟢 Running", true)
	systray.AddSeparator()

	if ip != "" && ip != "127.0.0.1" {
		if _, err := net.DialTimeout("tcp", fmt.Sprintf("%s:8080", ip), 500*time.Millisecond); err != nil {
			warn := makeItem("⚠️ Firewall", false)
			go func() {
				<-warn.ClickedCh
				fmt.Println("\nRun: netsh advfirewall firewall add rule name=\"RemoteControl\" dir=in action=allow protocol=tcp localport=8080")
			}()
			systray.AddSeparator()
		}
	}

	quit := makeItem("❌ Quit", false)
	<-quit.ClickedCh
	systray.Quit()
}
