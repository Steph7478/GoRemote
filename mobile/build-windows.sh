#!/bin/bash

go build -o RemoteControl.exe main.go

curl -s -L -o rcedit.exe https://github.com/electron/rcedit/releases/download/v2.0.0/rcedit-x64.exe
./rcedit.exe RemoteControl.exe --set-icon assets/icon.ico

echo 'netsh advfirewall firewall add rule name="Remote Control" dir=in action=allow protocol=TCP localport=8080' > install.bat
echo 'netsh advfirewall firewall delete rule name="Remote Control"' > uninstall.bat