#!/bin/bash

echo "🏗️ Building Remote Control for all platforms"
echo "==========================================="
echo ""

./build-android.sh
echo ""
./build-ios.sh
echo ""
./build-windows.sh
echo ""
./build-linux.sh

echo ""
echo "🎉 All builds complete!"
ls -lh *.apk *.exe *.app 2>/dev/null