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

:: Get the full path where the EXE will be
set "APP_PATH=%ProgramFiles%\RemoteControl\RemoteControl.exe"

:: Create rules for both port AND program
echo Adding firewall rule for port 8080...
netsh advfirewall firewall add rule name="Remote Control - Port" dir=in action=allow protocol=TCP localport=8080

echo Adding firewall rule for the application...
netsh advfirewall firewall add rule name="Remote Control - App" dir=in action=allow program="%APP_PATH%" enable=yes

echo.
echo ✅ Firewall rules added successfully!
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

echo Removing Remote Control firewall rules...

netsh advfirewall firewall delete rule name="Remote Control - Port"
netsh advfirewall firewall delete rule name="Remote Control - App"

echo ✅ Firewall rules removed!
echo.
echo Uninstallation complete!
timeout /t 3
EOF

rm -f rcedit.exe