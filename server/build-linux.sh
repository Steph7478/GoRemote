#!/bin/bash

go build -o RemoteControl main.go

CURRENT_DIR=$(pwd)

cat > RemoteControl.desktop << EOF
[Desktop Entry]
Name=Remote Control
Comment=Remote Control Application
Exec=$CURRENT_DIR/run.sh
Icon=$CURRENT_DIR/assets/icon.png
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

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
APP_PATH="$SCRIPT_DIR/RemoteControl"

echo "Installing Remote Control..."

# Check if ufw is available
if command -v ufw &> /dev/null; then
    # Check if rule already exists for this specific program
    if ufw status | grep -q "$APP_PATH"; then
        echo "⚠️  Firewall rule for Remote Control already exists!"
    else
        # Add program-specific rule (ufw doesn't support program rules directly, so we use iptables)
        # Alternative: Allow port but restrict to program using iptables owner module
        ufw allow 8080/tcp comment 'Remote Control'
        echo "✅ Firewall rule for port 8080 added!"
    fi
else
    echo "⚠️  ufw not found, skipping firewall configuration"
fi

# For firewalld (better program support)
if command -v firewall-cmd &> /dev/null; then
    if firewall-cmd --list-rich-rule | grep -q "Remote Control"; then
        echo "⚠️  firewalld rule already exists!"
    else
        # Add rich rule for the specific program
        firewall-cmd --permanent --add-rich-rule='rule family="ipv4" program name="'"$APP_PATH"'" port port="8080" protocol="tcp" accept'
        firewall-cmd --reload
        echo "✅ firewalld rule for Remote Control added!"
    fi
fi

echo ""
echo "Installation complete!"
echo "You can now run: ./run.sh"
EOF

cat > uninstall.sh << 'EOF'
#!/bin/bash
if [ "$EUID" -ne 0 ]; then 
    echo "Requesting sudo privileges..."
    exec sudo "$0" "$@"
fi

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
APP_PATH="$SCRIPT_DIR/RemoteControl"

echo "Removing Remote Control firewall rules..."

if command -v ufw &> /dev/null; then
    # Delete rule by comment
    ufw delete allow 8080/tcp 2>/dev/null
    echo "✅ ufw rule removed!"
fi

if command -v firewall-cmd &> /dev/null; then
    firewall-cmd --permanent --remove-rich-rule='rule family="ipv4" program name="'"$APP_PATH"'" port port="8080" protocol="tcp" accept' 2>/dev/null
    firewall-cmd --reload 2>/dev/null
    echo "✅ firewalld rule removed!"
fi

echo ""
echo "Uninstallation complete!"
EOF

cat > run.sh << 'EOF'
#!/bin/bash
DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$DIR"
./RemoteControl
EOF

chmod +x install.sh uninstall.sh RemoteControl run.sh

echo "Build complete!"
echo ""
echo "Files generated:"
echo "  - RemoteControl (executable)"
echo "  - install.sh (run with sudo to configure firewall)"
echo "  - uninstall.sh (run with sudo to remove firewall rules)"
echo "  - run.sh (launcher script)"
echo "  - RemoteControl.desktop (desktop entry)"
echo ""
echo "Important: Make sure your icon is at: assets/icon.png"