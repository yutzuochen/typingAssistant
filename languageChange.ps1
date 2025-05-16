# The function would work only if your English switching is "ctrl+shift+0" 

if (-not ([System.Management.Automation.PSTypeName]'InputLang').Type) {
    Add-Type @"
    using System;
    using System.Runtime.InteropServices;

    public class InputLang {
        [DllImport("user32.dll")]
        public static extern IntPtr GetForegroundWindow();

        [DllImport("user32.dll")]
        public static extern uint GetWindowThreadProcessId(IntPtr hWnd, out uint processId);

        [DllImport("user32.dll")]
        public static extern IntPtr GetKeyboardLayout(uint threadId);
    }
"@
}


function Get-CurrentInputLangID {
    $hWnd = [InputLang]::GetForegroundWindow()
    $threadId = [InputLang]::GetWindowThreadProcessId($hWnd, [ref]0)
    $hkl = [InputLang]::GetKeyboardLayout($threadId)
    return ($hkl.ToInt64() -band 0xFFFF)
}

# Get current input language ID
$langID = Get-CurrentInputLangID
Write-Host "Current input language ID: 0x$langID" -ForegroundColor Cyan

# 0x0404 = Traditional Chinese (Taiwan)
if ($langID -eq 0x0404) {
    Write-Host "Language is Traditional Chinese. Sending Ctrl + Shift + 0 to switch into Englisjj" -ForegroundColor Yellow

    $wshell = New-Object -ComObject wscript.shell
    $wshell.SendKeys("^+0")  # Ctrl + Shift + 0
}
else {
    Write-Host "Input language is not Traditional Chinese. No switch needed." -ForegroundColor Green
}
