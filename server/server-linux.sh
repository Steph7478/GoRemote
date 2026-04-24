#!/bin/bash

echo "🚀 Starting Remote Control Server..."
echo ""

echo "📦 Compiling server..."
go build -o Remote-Server main.go
echo "✅ Server compiled!"
echo ""

nohup ./Remote-Server > /dev/null 2>&1 &
echo "✅ Server running in background (no terminal)"