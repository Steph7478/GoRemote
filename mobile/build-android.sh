#!/bin/bash

echo "📱 Building Android APK with fyne-cross..."
echo "=========================================="
echo ""

go run github.com/fyne-io/fyne-cross@latest android \
    -app-id com.remotecontrol.app \
    -icon assets/icon.png \
    -name "Remote Control" \
    -app-version 1.0.0 \
    -app-build 1

echo ""
echo "✅ Android APK ready!"
ls -lh fyne-cross/dist/android/*.apk