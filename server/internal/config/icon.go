package config

import (
	_ "embed"
	"runtime"
)

//go:embed assets/icon.ico
var iconICO []byte

//go:embed assets/icon.png
var iconPNG []byte

func GetIcon() []byte {
	if runtime.GOOS == "windows" {
		return iconICO
	}
	return iconPNG
}
