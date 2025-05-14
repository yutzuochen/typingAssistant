Remove-Item -Path "resource.syso" -ErrorAction SilentlyContinue
magick cav27.jpg -define icon:auto-resize=256,128,64,48,32,16 cav27.ico
windres resource.rc -O coff -o resource.syso