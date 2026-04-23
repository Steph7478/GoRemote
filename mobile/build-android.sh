#!/bin/bash

echo "📱 Building Android APK..."
echo "========================"
echo ""

go mod tidy
go install fyne.io/fyne/v2/cmd/fyne@latest
~/go/bin/fyne package -os android \
    -appID com.remotecontrol.app \
    -icon assets/icon.png \
    -name "Remote Control" \
    -appVersion 1.0.0 \
    -arch arm64

echo ""
echo "✅ Done!"
ls -lh *.apk