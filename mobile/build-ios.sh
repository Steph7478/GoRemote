#!/bin/bash

echo "🍎 Building iOS app with fyne-cross..."
echo "======================================"
echo ""

go run github.com/fyne-io/fyne-cross@latest ios \
    -app-id com.remotecontrol.app \
    -icon assets/icon.png \
    -name "Remote Control" \
    -app-version 1.0.0 \
    -app-build 1 \
    -release \
    -arch arm64 \
    -ldflags="-s -w"

echo ""
echo "✅ iOS app ready!"
find fyne-cross/dist -name "*.app" -type d | head -1 | xargs ls -ld