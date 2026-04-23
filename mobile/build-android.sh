#!/bin/bash

echo "📱 Building Android APK..."
echo "========================"
echo ""

go mod tidy
go install fyne.io/tools/cmd/fyne@latest
~/go/bin/fyne package -os android/arm64 \
    -app-id com.remotecontrol.app \
    -icon assets/icon.png \
    -name "Remote Control"

echo ""
echo "✅ Done!"
ls -lh *.apk