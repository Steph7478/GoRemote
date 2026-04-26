#!/bin/bash

go build -o RemoteControl main.go

cat > RemoteControl.desktop << 'EOF'
[Desktop Entry]
Name=Remote Control
Comment=Remote Control Application
Exec=RemoteControl
Icon=remotecontrol
Terminal=false
Type=Application
Categories=Network;
StartupNotify=true
EOF

cat > install.sh << 'EOF'
#!/bin/bash
if [ "$EUID" -ne 0 ]; then 
    echo "Requesting sudo privileges..."
    exec sudo "$0" "$@"
fi

echo "Installing Remote Control..."

# Check if rule already exists
if ufw status | grep -q "8080/tcp.*ALLOW"; then
    echo "⚠️  Firewall rule for port 8080 already exists!"
else
    # Add firewall rule for port
    ufw allow 8080/tcp
    echo "✅ Firewall rule for port 8080 added!"
fi

# Optional: If using firewalld instead of ufw
if command -v firewall-cmd &> /dev/null; then
    firewall-cmd --permanent --add-port=8080/tcp
    firewall-cmd --reload
    echo "✅ firewalld rule added!"
fi

echo ""
echo "Installation complete!"
echo "You can now run: ./RemoteControl"
EOF

# Uninstall script
cat > uninstall.sh << 'EOF'
#!/bin/bash
if [ "$EUID" -ne 0 ]; then 
    echo "Requesting sudo privileges..."
    exec sudo "$0" "$@"
fi

echo "Removing Remote Control firewall rules..."

# Remove ufw rule if exists
if ufw status | grep -q "8080/tcp.*ALLOW"; then
    ufw delete allow 8080/tcp
    echo "✅ ufw rule removed!"
else
    echo "⚠️  ufw rule not found"
fi

# Remove firewalld rule if exists
if command -v firewall-cmd &> /dev/null; then
    firewall-cmd --permanent --remove-port=8080/tcp
    firewall-cmd --reload
    echo "✅ firewalld rule removed!"
fi

echo ""
echo "Uninstallation complete!"
EOF

chmod +x install.sh uninstall.sh RemoteControl

cat > run.sh << 'EOF'
#!/bin/bash
# Get the directory where this script is located
DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$DIR"
./RemoteControl
EOF
chmod +x run.sh

echo "Build complete!"
echo ""
echo "Files generated:"
echo "  - RemoteControl (executable)"
echo "  - install.sh (run with sudo to configure firewall)"
echo "  - uninstall.sh (run with sudo to remove firewall rules)"
echo "  - run.sh (launcher script)"
echo "  - RemoteControl.desktop (desktop entry)"