#!/bin/bash

echo "🚀 Starting Remote Control Server..."
echo ""

echo "📦 Compiling server..."
go build -o server main.go
echo "✅ Server compiled!"
echo ""

./server