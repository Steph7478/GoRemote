#!/bin/bash

go build -ldflags="-H windowsgui" -o RemoteControl.exe main.go

curl -s -L -o rcedit.exe https://github.com/electron/rcedit/releases/download/v2.0.0/rcedit-x64.exe
./rcedit.exe RemoteControl.exe --set-icon assets/icon.ico

rm -f rcedit.exe

echo "Build complete! Output: RemoteControl.exe"