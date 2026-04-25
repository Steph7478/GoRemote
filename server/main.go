package main

import (
	"server/internal/config"

	"github.com/getlantern/systray"
)

func main() {
	systray.Run(config.Tray, nil)
}
