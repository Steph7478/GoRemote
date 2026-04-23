#!/bin/bash

echo "📱 Building Android APK with fyne-cross..."
echo "=========================================="
echo ""

go run github.com/fyne-io/fyne-cross@latest android \
    -app-id com.remotecontrol.app \
    -icon assets/icon.png \
    -name "Remote Control" \
    -app-version 1.0.0 \
    -app-build 1 \
    -release \
    -arch arm64 \
    -ldflags="-s -w"

echo ""
echo "✅ Android APK ready!"
find fyne-cross/dist -name "*.apk" -type f | head -1 | xargs ls -lh