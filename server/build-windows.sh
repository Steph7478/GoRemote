#!/bin/bash

go build -ldflags="-H windowsgui" -o RemoteControl.exe main.go

curl -s -L -o rcedit.exe https://github.com/electron/rcedit/releases/download/v2.0.0/rcedit-x64.exe
./rcedit.exe RemoteControl.exe --set-icon config/assets/icon.ico

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

:: Get the CURRENT directory (where the user extracted the files)
set "APP_PATH=%CD%\RemoteControl.exe"

:: Check if rule already exists for this specific app
netsh advfirewall firewall show rule name="Remote Control" >nul 2>&1
if %errorLevel% equ 0 (
    echo Firewall rule already exists for Remote Control!
    echo No changes made.
    timeout /t 3
    exit /b
)

:: Add firewall rule for the application (this is enough!)
echo Adding firewall rule for Remote Control...
netsh advfirewall firewall add rule name="Remote Control" dir=in action=allow program="%APP_PATH%" enable=yes

if %errorLevel% equ 0 (
    echo.
    echo ✅ Firewall rule added successfully!
    echo    Program: %APP_PATH%
) else (
    echo.
    echo ❌ Failed to add firewall rule
)

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

:: Check if rule exists before removing
netsh advfirewall firewall show rule name="Remote Control" >nul 2>&1
if %errorLevel% neq 0 (
    echo Firewall rule not found!
    echo Nothing to remove.
    timeout /t 3
    exit /b
)

:: Remove firewall rule
netsh advfirewall firewall delete rule name="Remote Control"

echo ✅ Firewall rule removed successfully!
timeout /t 3
EOF

rm -f rcedit.exe
echo "Build complete!"