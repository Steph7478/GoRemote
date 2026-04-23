#!/bin/bash

echo "📦 Building Windows executable..."
go build -o "Remote-Control.exe" main.go
echo "🎯 Done!"
ls -lh Remote-Control.exe