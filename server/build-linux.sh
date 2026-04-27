#!/bin/bash

APP_NAME="RemoteControl"
DISPLAY_NAME="Remote Control"
EXE_NAME="RemoteControl"
FORNECEDOR="Stéphanie Gurgel"
VERSAO="1.0.0"
ICON_PATH="internal/config/assets/icon.png"

go build -o "$EXE_NAME" .

CURRENT_DIR="$(pwd)"

cat > "$APP_NAME.desktop" << EOF
[Desktop Entry]
Version=$VERSAO
Name=$DISPLAY_NAME
GenericName=$APP_NAME
Comment=Remote Control Application
Exec=$CURRENT_DIR/$EXE_NAME
Icon=$CURRENT_DIR/$ICON_PATH
Terminal=false
Type=Application
Categories=Network;
StartupNotify=true
X-App-Version=$VERSAO
X-App-Vendor=$FORNECEDOR
X-App-Company=$FORNECEDOR
X-App-OriginalFilename=$EXE_NAME
X-App-Copyright=Copyright © 2026 $FORNECEDOR
EOF

cat > run.sh << 'EOF'
#!/bin/bash
DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$DIR"
./RemoteControl
EOF

chmod +x run.sh "$EXE_NAME" "$APP_NAME.desktop"

mkdir -p "$HOME/.local/share/applications"
cp "$APP_NAME.desktop" "$HOME/.local/share/applications/"
update-desktop-database "$HOME/.local/share/applications" 2>/dev/null || true

echo "Build complete! Output: $EXE_NAME"
echo "Desktop entry installed: $HOME/.local/share/applications/$APP_NAME.desktop"