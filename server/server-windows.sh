#!/bin/bash

echo "🚀 Starting Remote Control Server..."
echo ""

echo "📦 Compiling server..."
go build -ldflags="-H windowsgui -s -w" -o Remote-Server.exe main.go
echo "✅ Server compiled!"
echo ""

./Remote-Server.exe &