# 🖱️ Remote Control

Remote Control is a cross-platform app that lets you control your PC from your phone over the same WiFi network.

The project has two parts:

- **PC Server**: runs on your computer, stays in the system tray, exposes a WebSocket server and receives mouse/keyboard commands.
- **Mobile/Desktop Client**: built with Fyne, discovers the server automatically on the local network and sends remote control commands.

---

## 📱 Overview

Remote Control turns your phone into a wireless mouse and keyboard for your PC.

The app uses **automatic server discovery**, so you do not need to manually type the PC IP address. Just run the server on your computer, open the app on your phone, and tap **Connect**.

Communication happens through:

- **WebSocket** on port `8080`
- **Local network discovery** using peer discovery
- **JSON messages** for mouse, keyboard and text commands

---

## ✨ Features

- 🖱️ **Wireless Mouse Pad**  
  Use your phone screen as a touchpad.

- 👆 **Tap to Click**  
  Tap on the mouse pad area to click on the PC.

- ⌨️ **Keyboard Input**  
  Type text in the mobile input before sending it to the PC.

- 🔙 **Remote Delete / Backspace Button**  
  The app has a `Del` button that sends a backspace command directly to the PC.

- ↵ **Remote Enter Button**  
  The app has an `Enter` button.  
  If the input has text, it sends the text.  
  If the input is empty, it sends an Enter key command to the PC.

- 📲 **Normal Mobile Input Behavior**  
  The phone keyboard works normally inside the input field.  
  You can type, delete, edit text, and use the cursor before sending.

- ⚡ **Low Latency**  
  Real-time control using WebSocket.

- 🔎 **Automatic Server Discovery**  
  The client searches for the PC server automatically on the local network.

- 🎚️ **Adjustable Mouse Speed**  
  Mouse sensitivity can be changed inside the app from `0.1` to `5.0`.

- 🖥️ **System Tray Server**  
  The PC server runs in the tray and shows the current IP and port.

- 📱 **Mobile Support**  
  Works on Android and iOS through Fyne.

- 💻 **Desktop Client Support**  
  The client can also run on Windows, Linux and macOS.

---

## 📋 Requirements

### Development

- Go `1.26` or higher
- Fyne `v2`
- Gin
- Gorilla WebSocket
- peerdiscovery
- systray

### Android Build

To build for Android, you need:

- Android SDK
- Android NDK
- Java/JDK
- Fyne CLI

### iOS Build

To build for iOS, you need:

- macOS
- Xcode
- Fyne CLI

---

## 🚀 How It Works

### PC Server

The PC server:

1. Starts a WebSocket server on `0.0.0.0:8080`

2. Registers routes, including the WebSocket endpoint `/ws`

3. Broadcasts its local IP on the network using peer discovery with this payload format:

`RemoteControl:<LOCAL_IP>`

4. Opens a system tray icon showing:

`<LOCAL_IP>:8080`

`Port: 8080`

`Running`

5. Receives commands from the client and controls the mouse/keyboard on the PC.

---

### Client App

The client app:

1. Starts disconnected.
2. When the user taps **Connect**, it searches the local network for the server.
3. If a server is found, it connects to `ws://<SERVER_IP>:8080/ws`
4. After connecting, it enables the mouse pad and keyboard.
5. Mouse movement is multiplied by the selected sensitivity value.
6. Keyboard and mouse events are sent as JSON messages.

---

## 🔌 Firewall Notice

The PC server uses port `8080`.

Your firewall must allow incoming TCP connections on this port.

If the mobile app cannot connect, the most common reason is that the PC firewall is blocking port `8080`.

---

## 🪟 Windows Firewall

On Windows, allow port `8080` with this command in an Administrator terminal:

`netsh advfirewall firewall add rule name="RemoteControl" dir=in action=allow protocol=tcp localport=8080`

The server also shows a **Firewall** warning in the tray menu when it detects that the port may not be reachable.

---

## 🐧 Linux Firewall

If you use `ufw`, allow the port with:

`sudo ufw allow 8080/tcp`

If you use another firewall manager, allow incoming TCP traffic on port `8080`.

---

## 🍎 macOS Firewall

On macOS, allow the server application to accept incoming connections when prompted.

You can also check it manually in:

`System Settings > Network > Firewall`

Make sure the server app is allowed to receive incoming connections.

---

## 📡 Network Requirements

Both devices must be on the same local network.

Example:

`PC: 192.168.1.20`

`Phone: 192.168.1.35`

The app may not work if:

- the phone is using mobile data instead of WiFi;
- the PC and phone are on different networks;
- the router blocks device-to-device communication;
- the firewall blocks port `8080`;
- VPN software changes the local network route.

Some routers have a setting called **AP Isolation**, **Client Isolation**, or **Wireless Isolation**.

This must be disabled, otherwise devices on the same WiFi cannot see each other.

---

## 🎮 Usage

### 1. Start the PC Server

Run the server on your computer.

The server starts in the system tray and shows the local IP and port, for example:

`192.168.x.x:8080`

---

### 2. Open the Client App

Open the app on your phone or desktop.

The app starts with the status `🔴`.

Tap **Connect**.

The app will automatically search for the PC server.

---

### 3. Connection Status

The status icon means:

- `🔴` Disconnected
- `⏳` Connecting / searching server
- `✅` Connected
- `❌` Failed to connect

When connected, the **Connect** button is hidden and the **Disconnect** button is shown.

---

### 4. Mouse Control

Use the mouse pad area to control the PC mouse.

- Drag: move mouse
- Tap: click
- Sensitivity: adjust with `-` and `+`

The sensitivity range is `0.1` to `5.0`.

Default value: `1.0`

---

### 5. Keyboard Control

The keyboard area has:

- a normal text input
- a `Del` button
- an `Enter` button

The phone keyboard behaves normally inside the input field.

You can:

- type text;
- delete text with the phone keyboard;
- edit before sending;
- move the cursor normally;
- send the text when ready.

---

## 🛠️ Building from Source

### Build Mobile Client

Android:

`cd mobile`

`./build-android.sh`

iOS:

`cd mobile`

`./build-ios.sh`

Android builds require Android SDK/NDK configured locally.

iOS builds require macOS and Xcode.

---

### Build Desktop Client

Windows:

`cd mobile`

`./build-windows.sh`

Linux:

`cd mobile`

`./build-linux.sh`

---

### Build PC Server

Windows:

`cd server`

`./server-windows.sh`

Linux:

`cd server`

`./server-linux.sh`

---

## ⚙️ Configuration

Default configuration:

- Server host: `0.0.0.0`
- Server port: `8080`
- WebSocket path: `/ws`
- Protocol: `ws://`
- Discovery payload: `RemoteControl:<ip>`
- Mouse sensitivity: `0.1` to `5.0`
- Default sensitivity: `1.0`

---

## 🧪 Troubleshooting

### App does not connect

Check:

1. PC server is running.
2. Phone and PC are on the same WiFi.
3. Firewall allows TCP port `8080`.
4. Router does not block device-to-device communication.
5. VPN is disabled or not interfering with local network traffic.

---

### Server appears to run, but phone cannot find it

Try:

- disabling VPN;
- checking if both devices are on the same subnet;
- disabling AP/client isolation on the router;
- allowing port `8080` in the firewall;
- restarting the server after changing firewall settings.

---

### Windows blocks the connection

Run this command as Administrator:

`netsh advfirewall firewall add rule name="RemoteControl" dir=in action=allow protocol=tcp localport=8080`

---

## 🧩 Main Components

### PC Server

Responsible for:

- starting the HTTP/WebSocket server;
- advertising the server on the local network;
- showing tray status;
- warning about firewall issues;
- receiving remote control commands.

---

### Client

Responsible for:

- discovering the server automatically;
- connecting to the WebSocket endpoint;
- sending mouse movement;
- sending keyboard text;
- sending remote key commands;
- managing sensitivity;
- updating UI connection status.

---

## 🔐 Security Notice

Remote Control is intended for trusted local networks only.

It does not provide authentication or encryption by default.

Use it only on networks you trust, such as your home WiFi.

Avoid using it on public WiFi networks unless you add authentication and transport security.

---