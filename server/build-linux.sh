#!/bin/bash

go build -o RemoteControl main.go

cat > install.sh << 'EOF'
#!/bin/bash
echo "Configuring firewall for Remote Control..."
sudo ufw allow 8080/tcp
echo "Port 8080 opened in firewall!"
echo ""
echo "To run: ./RemoteControl"
EOF
chmod +x install.sh

cat > uninstall.sh << 'EOF'
#!/bin/bash
echo "Removing firewall rule..."
sudo ufw delete allow 8080/tcp
echo "Firewall rule removed!"
echo ""
echo "You can now delete the RemoteControl file"
EOF
chmod +x uninstall.sh

chmod +x RemoteControl

echo "Build complete! Output: RemoteControl"