#!/bin/bash

echo "🚀 Starting Remote Control Server..."
echo ""

echo "📦 Compiling server..."
go build -o Remote-Server.exe main.go
echo "✅ Server compiled!"
echo ""