#!/bin/bash

echo "🔨 Compiling mobile client (Windows Desktop version) for testing..."

GOOS=windows GOARCH=amd64 go build -o Remote-Control.exe main.go

echo "✅ Done! Binary generated: Remote-Control.exe"
echo "Tip: This will open the app interface in a window on Windows."

./Remote-Control.exe
