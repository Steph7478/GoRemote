#!/bin/bash

echo "🔨 Compiling server for Windows (Test)..."
GOOS=windows GOARCH=amd64 go build -o Remote-Server.exe main.go

echo "✅ Done! Binary generated: Remote-Server.exe"
echo "Tip: On Windows, simply run the generated .exe file."

./Remote-Server.exe
