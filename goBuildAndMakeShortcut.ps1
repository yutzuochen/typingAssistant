$Dir = "C:\Users\ASUS\Documents\autoPressing\"
$ExecutableFile = "auto_cav.exe"
$LinkPath = "C:\Users\ASUS\Desktop\auto_cav.exe.lnk"

go build -o $ExecutableFile main.go
$WShell = New-Object -ComObject WScript.Shell
$Shortcut = $WShell.CreateShortcut($LinkPath)
$Shortcut.TargetPath = $Dir + $ExecutableFile
$Shortcut.Save()

