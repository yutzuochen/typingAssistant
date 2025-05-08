package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func main() {
	fmt.Println("=== Listening for keyboard and mouse ====")
	keyChan := hook.Start()
	defer hook.End()
	for key := range keyChan {
		if key.Kind == hook.KeyHold && key.Rawcode == 20 && key.Keychar == 65535 { // listen CapsLock key press
			go generateVills()
		} else if key.Kind == hook.KeyDown && key.Rawcode == 67 { // listen for "c" or "C"
			go cavDancing()
		}
	}
}

func generateVills() {
	time.After(83 * time.Millisecond)
	robotgo.KeyTap("d")
	time.After(75 * time.Millisecond)
	robotgo.KeyTap("d")
}

func cavDancing() {
	robotgo.KeyTap("x")
	time.After(296 * time.Millisecond)
	robotgo.KeyTap("z")
	time.After(373 * time.Millisecond)
	robotgo.KeyTap("x")
	time.After(178 * time.Millisecond)
	robotgo.KeyTap("z")
	time.After(178 * time.Millisecond)
	robotgo.KeyTap("x")
	time.After(178 * time.Millisecond)
	robotgo.KeyTap("z")
}
