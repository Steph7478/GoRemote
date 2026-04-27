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
  Use your phone screen as a touchpad with three modes:
  - **Move Mode** (🖱️ Move) - normal mouse movement
  - **Select Mode** (👆 Select) - drag selection with automatic down/up events
  - **Scroll Mode** (📜 Scroll) - drag for scrolling

- 👆 **Tap to Click**  
  Tap on the mouse pad area for left click.

- 🖱️ **Right Click**
  Two-finger tap or long-press for right click.

- 📜 **Scrolling**
  Drag in Scroll mode.

- ⌨️ **Keyboard Input**  
  Type text in the input field before sending to PC.

- 🔙 **Delete Button**  
  `Del` button sends backspace command to PC.

- ↵ **Smart Enter Button**  
  - With text in input: sends typed text
  - Empty input: sends Enter key command

- 📲 **Native Keyboard Behavior**  
  Phone keyboard works normally inside input field.

- ⚡ **Low Latency**  
  Real-time control using WebSocket.

- 🔎 **Automatic Server Discovery**  
  Client searches for PC server automatically.

- 🎚️ **Adjustable Mouse Speed**  
  Sensitivity range: `0.1` to `5.0` (default: `1.0`)

- 🖥️ **System Tray Server**  
  PC server runs in tray showing current IP and port.

- 🔍 **Discovery Control**
  Pause/resume server discovery from system tray.

---

## 🎮 Usage

### 1. Start the PC Server

Run the server on your computer. It starts in system tray showing local IP and port.

### 2. Open Client App

Open app on your phone. Status starts 🔴 (disconnected).

Tap **Connect**. App searches for server automatically.

### 3. Connection Status

- 🔴 Disconnected
- ⏳ Connecting / searching
- ✅ Connected
- ❌ Connection failed

### 4. Mouse Control Modes

**Move Mode** (default):
- Drag for normal mouse movement
- Sensitivity affects movement speed

**Select Mode**:
- Touch down: sends down event
- Drag: moves while holding
- Release: sends up event

**Scroll Mode**:
- Drag for scrolling

### 5. Mouse Actions

- Left click: one-finger tap anywhere on pad
- Right click: two-finger tap or long press
- Scroll: drag in Scroll mode

### 6. Adjust Mouse Speed

Use - and + buttons next to "Speed" label.

Range: 0.1 (slow) to 5.0 (fast)

Speed persists across app restarts.

### 7. Keyboard Control

- Text input: type normally
- Del button: sends backspace to PC
- Enter button: sends text if present, otherwise Enter key

---

## ⚙️ Configuration

- Server port: 8080
- WebSocket path: /ws
- Discovery payload: RemoteControl:<ip>
- Mouse sensitivity: 0.1 to 5.0

---

## 🔌 Firewall Notice

PC server uses port 8080. Firewall must allow incoming TCP connections.

**Windows (Admin):**
netsh advfirewall firewall add rule name="RemoteControl" dir=in action=allow protocol=tcp localport=8080

**Linux (UFW):**
sudo ufw allow 8080/tcp

**macOS:** Allow server app when prompted.

---

## 📡 Network Requirements

Both devices on same local network.

May not work if:
- Phone on mobile data
- Different networks
- Router blocks device-to-device communication (AP Isolation)
- Firewall blocks port 8080
- VPN active

---

## 🧪 Troubleshooting

**App doesn't connect:**
1. PC server running?
2. Same WiFi network?
3. Firewall allows port 8080?
4. AP Isolation disabled?
5. VPN disabled?

**Server runs but phone can't find it:**
- Disable VPN
- Check same subnet
- Disable AP/client isolation
- Allow port 8080 in firewall
- Restart server

---

## 🔐 Security Notice

Remote Control is for trusted local networks only.

No authentication or encryption by default.

Use only on networks you trust (e.g., home WiFi).