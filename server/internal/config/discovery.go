package config

import (
	"fmt"
	"server/internal/utils"
	"time"

	"github.com/schollz/peerdiscovery"
	"github.com/sqweek/dialog"
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
			if dialog.Message("Device trying to connect:\n%s", peers[0].Address).
				Title("Remote Control").
				YesNo() {
				found = true
				openServerPort()
			}
		}

		time.Sleep(2 * time.Second)
	}
}
