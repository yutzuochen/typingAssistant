$Dir = "C:\Users\ASUS\Documents\typingAssistant\"
$ExecutableFile = "auto_cav.exe"
$LinkPath = "C:\Users\ASUS\Desktop\auto_cav.lnk"
$startIn = "C:\Users\ASUS\Documents\typingAssistant"

go build -o $ExecutableFile main.go
$WShell = New-Object -ComObject WScript.Shell
$Shortcut = $WShell.CreateShortcut($LinkPath)
$Shortcut.WorkingDirectory = $startIn
$Shortcut.TargetPath = $Dir + $ExecutableFile
$Shortcut.Save()

