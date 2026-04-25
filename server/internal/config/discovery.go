package config

import (
	"fmt"
	"server/internal/utils"
	"time"

	"github.com/schollz/peerdiscovery"
)

var found bool

func startDiscovery() {
	payload := []byte(fmt.Sprintf("RemoteControl:%s", utils.GetLocalIP()))

	for {
		if !discoveryRunning {
			time.Sleep(time.Second)
			continue
		}

		peers, _ := peerdiscovery.Discover(peerdiscovery.Settings{
			Payload:   payload,
			Limit:     1,
			Delay:     2 * time.Second,
			AllowSelf: false,
		})

		found = found && serverRunning

		if !found && len(peers) > 0 {
			fmt.Println("📡 Device connected:", peers[0].Address)
			found = true
			openServerPort()
		}

		time.Sleep(2 * time.Second)
	}
}
