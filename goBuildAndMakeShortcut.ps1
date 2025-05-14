$Dir = "C:\Users\ASUS\Documents\autoPressing"
$ExecutableFile = "cavCivi.exe"
$LinkPath = "C:\Users\ASUS\Desktop\cavCivi.lnk"

go build -o $ExecutableFile main.go
$WShell = New-Object -ComObject WScript.Shell
$Shortcut = $WShell.CreateShortcut($LinkPath)
$Shortcut.TargetPath = $Dir + $ExecutableFile
$Shortcut.Save()

