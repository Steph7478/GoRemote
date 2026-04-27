#!/bin/bash

APP_NAME="RemoteControl"
EXE_NAME="RemoteControl.exe"
FORNECEDOR="Stéphanie Gurgel"
VERSAO="1.0.0"

go build -ldflags="-H windowsgui" -o "$EXE_NAME" .

curl -s -L -o rcedit.exe https://github.com/electron/rcedit/releases/download/v2.0.0/rcedit-x64.exe

./rcedit.exe "$EXE_NAME" \
  --set-icon "internal/config/assets/icon.ico" \
  --set-version-string "CompanyName" "$FORNECEDOR" \
  --set-version-string "FileDescription" "$APP_NAME" \
  --set-version-string "ProductName" "$APP_NAME" \
  --set-version-string "OriginalFilename" "$EXE_NAME" \
  --set-version-string "LegalCopyright" "Copyright © 2026 $FORNECEDOR" \
  --set-file-version "$VERSAO" \
  --set-product-version "$VERSAO"

rm -f rcedit.exe

echo "Build complete! Output: $EXE_NAME"