#!/bin/bash

echo "🏗️ Building Remote Control for all platforms"
echo "==========================================="
echo ""

./build-android.sh
echo ""
./build-ios.sh
echo ""
./build-desktop.sh

echo ""
echo "🎉 All builds complete!"
ls -lh *.apk *.exe *.app 2>/dev/null