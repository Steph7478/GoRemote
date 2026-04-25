package config

import (
	"time"

	"github.com/schollz/peerdiscovery"
	"github.com/sqweek/dialog"
)

var found bool

func startDiscovery() {
	payload := []byte("RemoteControl")

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
			if dialog.Message("A device is trying to connect.").
				Title("Remote Control").
				YesNo() {
				found = true
				openServerPort()
			}
		}

		time.Sleep(2 * time.Second)
	}
}