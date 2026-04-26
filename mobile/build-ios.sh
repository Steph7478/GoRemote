#!/bin/bash

echo "🍎 Building iOS app..."
echo "====================="
echo ""

go mod tidy
go install fyne.io/tools/cmd/fyne@latest
~/go/bin/fyne package -os ios/arm64 \
    -app-id com.remotecontrol.app \
    -icon assets/icon.ico \
    -name "Remote Control"

echo ""
echo "✅ Done!"
ls -lh *.app