#!/bin/bash

echo "📱 Starting Mobile App (test mode on PC)..."
echo ""

echo "📦 Compiling app..."
go build -o app main.go
echo "✅ App compiled!"
echo ""

./app