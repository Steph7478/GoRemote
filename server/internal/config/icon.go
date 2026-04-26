package config

import (
    "embed"
    "runtime"
)

var iconICO []byte
var iconPNG []byte

func GetIcon() []byte {
    if runtime.GOOS == "windows" {
        return iconICO
    }
    return iconPNG
}