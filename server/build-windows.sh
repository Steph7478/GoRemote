#!/bin/bash

go build -ldflags="-H windowsgui" -o RemoteControl.exe main.go

curl -s -L -o rcedit.exe https://github.com/electron/rcedit/releases/download/v2.0.0/rcedit-x64.exe
./rcedit.exe RemoteControl.exe --set-icon assets/icon.ico

cat > install.bat << 'EOF'
@echo off
:: Request administrator privileges
net session >nul 2>&1
if %errorLevel% neq 0 (
    echo Requesting administrator privileges...
    powershell start -verb runas '%0'
    exit /b
)

echo Installing Remote Control...
netsh advfirewall firewall add rule name="Remote Control" dir=in action=allow protocol=TCP localport=8080
echo Firewall rule added for port 8080!
echo.
echo Installation complete!
timeout /t 3
EOF

cat > uninstall.bat << 'EOF'
@echo off
:: Request administrator privileges
net session >nul 2>&1
if %errorLevel% neq 0 (
    echo Requesting administrator privileges...
    powershell start -verb runas '%0'
    exit /b
)

echo Removing Remote Control firewall rule...
netsh advfirewall firewall delete rule name="Remote Control"
echo Firewall rule removed!
echo.
echo Uninstallation complete!
timeout /t 3
EOF