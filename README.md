# 🖱️ Remote Control

A cross-platform application to control your PC remotely from your mobile device over WiFi.

## 📱 Overview

Remote Control allows you to use your smartphone as a wireless mouse and keyboard for your computer. The project consists of two main components:

- **Mobile App** - Touchpad interface for Android/iOS built with Fyne
- **PC Server** - Lightweight server that receives commands and controls mouse/keyboard

## ✨ Features

- 🖱️ **Touchpad Mode** - Use your phone as a wireless mouse
- 🖱️ **Click Support** - Tap to click on the touchpad
- ⌨️ **Keyboard Input** - Send text directly from your phone
- ⚡ **Low Latency** - Real-time communication via WebSocket
- 🔌 **Easy Connection** - Simple IP-based connection
- 🎚️ **Adjustable Sensitivity** - Customize mouse speed
- 📱 **Cross-Platform Mobile** - Works on Android and iOS
- 💻 **Desktop Support** - Client also available for Windows, Linux, macOS

## 📋 Prerequisites

### For Development
- Go 1.26 or higher
- Fyne v2.7.3
- WebSocket library

### For Mobile Build (Android)
- Docker (for fyne-cross) OR Android NDK + SDK

## 🚀 Installation

### PC Server Setup

1. Navigate to server directory: `cd pc-server`

2. Run the server:
   - Windows: `./server-windows.sh`
   - Linux: `./server-linux.sh`

The server will display your local IP address. Note this for connecting from your mobile device.

### Mobile Client Setup

#### On Android
`cd mobile && ./build-android.sh`

#### On iOS
`cd mobile && ./build-ios.sh`

#### Desktop Client
- Windows: `./build-windows.sh`
- Linux: `./build-linux.sh`

## 🎮 Usage

1. **Start the PC Server** - Run the server executable on your computer and note the IP address displayed

2. **Connect Mobile App** - Open the app on your phone, enter the PC's IP address, tap "Connect"

3. **Control Your PC**
   - Move mouse: Drag finger on the touchpad area
   - Click: Tap on the touchpad
   - Scroll: Two-finger scroll on supported devices
   - Type: Use the keyboard tab to send text

## 🛠️ Building from Source

### Build Client
- Windows: `./build-windows.sh` generates `Remote-Control.exe`
- Linux: `./build-linux.sh` generates `Remote-Control`
- Android: `./build-android.sh` generates `Remote-Control.apk`
- iOS: `./build-ios.sh` generates `Remote-Control.app`

### Build PC Server
- Windows: `./server-windows.sh` generates `Remote-Server.exe`
- Linux: `./server-linux.sh` generates `Remote-Server`

## 🔧 Configuration

- **Server Port**: 8080 (default)
- **Protocol**: WebSocket (ws://)
- **Mouse Sensitivity**: Adjustable from 0.5 to 3.0
- **Auto-save**: Last used IP is saved locally