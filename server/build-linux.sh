#!/bin/bash

go build -o RemoteControl main.go

cat > RemoteControl.desktop << 'EOF'
[Desktop Entry]
Name=Remote Control
Comment=Remote Control Application
Exec=$PWD/RemoteControl
Icon=$PWD/assets/icon.png
Terminal=false
Type=Application
Categories=Network;
EOF

cat > install.sh << 'EOF'
#!/bin/bash
if [ "$EUID" -ne 0 ]; then 
    exec sudo "$0" "$@"
fi
ufw allow 8080/tcp
echo "✅ Firewall rule added!"
EOF

chmod +x install.sh RemoteControl

echo "Build complete!"