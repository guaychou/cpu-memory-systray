#!/bin/bash
echo "`date` : Build starting . . . "
mkdir -p ~/Applications/CpuMemoryService.app/Contents/MacOS
go mod download
GOMAXPROCS=1 go build -o CpuMemoryService 2>err.log
if [ -s "err.log" ];then
  echo -e "`date` : Build failed, check err.log "
  exit 2
fi
cp CpuMemoryService ~/Applications/CpuMemoryService.app/Contents/MacOS
cat << EOF > ~/Applications/CpuMemoryService.app/Contents/Info.plist
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
        <key>NSHighResolutionCapable</key>
        <string>True</string>
        <!-- avoid showing the app on the Dock -->
        <key>LSUIElement</key>
        <string>1</string>
</dict>
</plist>
EOF
echo -e "`date` : Build Success . . . Find your app, in your home Applications "
