#!/bin/bash

go build -o RemoteControl main.go

CURRENT_DIR=$(pwd)

cat > RemoteControl.desktop << EOF
[Desktop Entry]
Name=Remote Control
Comment=Remote Control Application
Exec=$CURRENT_DIR/RemoteControl
Icon=$CURRENT_DIR/assets/icon.png
Terminal=false
Type=Application
Categories=Network;
EOF

cat > run.sh << 'EOF'
#!/bin/bash
DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$DIR"
./RemoteControl
EOF

chmod +x run.sh RemoteControl

echo "Build complete! Output: RemoteControl"