#!/bin/bash

echo "🍎 Building iOS app..."
echo "======================"

go mod tidy
go install fyne.io/tools/cmd/fyne@latest

~/go/bin/fyne package -os ios \
    -app-id com.remotecontrol.app \
    -icon assets/icon.ico \
    -name "Remote Control"

mkdir Payload
cp -r RemoteControl.app Payload/
zip -r RemoteControl.ipa Payload

echo "✅ Done! Output: RemoteControl.ipa"