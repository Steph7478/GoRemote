#!/bin/bash

echo "📦 Building Linux executable..."
go build -o "Remote-Control" main.go
echo "🎯 Done!"
ls -lh Remote-Control