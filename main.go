package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var (
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
)
var (
	CloseProgramPath = "sound\\closeProgram.mp3"
	StartProgramPath = "sound\\startProgram.mp3"
	PendingPath      = "sound\\pending.mp3"
	EndPendingPath   = "sound\\endPending.mp3"
)

func main() {

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}
	fmt.Println("Current directory is: ", cwd)
	// Detect language input. if current input is Chinese, then reminds user.

	go playSound(StartProgramPath)
	fmt.Println("======= Listening for pressing by keyboard & mouse =======")
	keyChan := hook.Start()
	defer hook.End()
	ctx, cancel = context.WithCancel(context.Background())
	// for {
	// keyChan := hook.Start()
	mainLoop(keyChan)
	// hook.End()
	// }
	fmt.Println("end of main")
}

func generateVills() {
	// robotgo.KeyTap("caps_lock")
	time.Sleep(83 * time.Millisecond)
	robotgo.KeyTap("d")
	time.Sleep(75 * time.Millisecond)
	robotgo.KeyTap("d")
}

func mainLoop(keyChan chan hook.Event) {
	for key := range keyChan {
		// test{
		if key.Kind == hook.KeyHold || key.Kind == hook.KeyDown {
			fmt.Printf("[main loop start] %+v\n", key)
		}
		// }test
		if key.Kind == hook.KeyHold && key.Rawcode == 187 { // listen "=" key press (my new 'Capslock')
			// cause we might be dancing our cav (zxzx), so we need to stop function - cavDancin first
			fmt.Printf("pressing chaged-CapsLock: %+v\n", key)
			cancel()
			wg.Wait()
			ctx, cancel = context.WithCancel(context.Background())
			// start to generate villagers from TC
			go generateVills()
		} else if key.Kind == hook.KeyDown && key.Rawcode == 67 { // listen for "c" or "C"
			wg.Add(1)
			go cavDancing(ctx)
		} else if key.Kind == hook.KeyDown && key.Rawcode == 13 && key.Mask == 2 { // pressing Enter + Ctrl
			go playSound(CloseProgramPath)
			fmt.Println("Enter + Ctrl pressed. Exiting program.")

			// robotgo.KeyTap("*")
			cancel()
			time.Sleep(103 * time.Millisecond)
			wg.Wait()
			os.Exit(0)
		} else if key.Kind == hook.KeyDown && key.Rawcode == 13 { // pressing Enter
			cancel()
			ctx, cancel = context.WithCancel(context.Background())
			fmt.Println("Enter pressed. Pending program.")
			go playSound(PendingPath)
			time.Sleep(53 * time.Millisecond)
			robotgo.KeyTap(";")
			signal := make(chan bool)
			go pending(keyChan, signal)
			<-signal
			fmt.Println("===== open the pending channel =======")

			// keyChan = hook.Start()
			// hook.End()
			// break
			// pendingCh := make(chan string)
			// hook.End()
			// time.Sleep(100 * time.Millisecond) // let goroutines respond to cancel
			// os.Exit(0)
		}

	}
	fmt.Println("end mainLoop")
}

func cavDancing(ctx context.Context) {
	defer wg.Done()
	robotgo.KeyTap("x")

	select {
	case <-ctx.Done():
		return
	case <-time.After(296 * time.Millisecond):
	}
	robotgo.KeyTap("z")

	select {
	case <-ctx.Done():
		return
	case <-time.After(373 * time.Millisecond):
	}
	robotgo.KeyTap("x")

	select {
	case <-ctx.Done():
		return
	case <-time.After(178 * time.Millisecond):
	}
	robotgo.KeyTap("z")
	select {
	case <-ctx.Done():
		return
	case <-time.After(267 * time.Millisecond):
	}
	robotgo.KeyTap("x")

	select {
	case <-ctx.Done():
		return
	case <-time.After(154 * time.Millisecond):
	}
	robotgo.KeyTap("z")
}

func playSound(relativePath string) {
	cwd, err := os.Getwd()
	if err != nil {
		// fmt.Println("Error getting working directory:", err)
		return
	}
	mp3Path := filepath.Join(cwd, relativePath)
	// fmt.Println("Playing:", mp3Path)

	cmd := exec.Command("ffplay", "-nodisp", "-autoexit", "-loglevel", "quiet", mp3Path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error playing sound:", err)
	}
}

func pending(keyChan chan hook.Event, signal chan bool) {
	// keyChan2 := hook.Start()
	for key := range keyChan {
		if key.Kind == hook.KeyHold && key.Rawcode == 187 { // "if press capslock"
			go playSound(EndPendingPath)
			fmt.Printf("[pending1] %+v\n", key)
			// hook.End()
			signal <- true
			return
		} else if key.Kind == hook.KeyHold || key.Kind == hook.KeyDown {
			// fmt.Println("====== pending2 ======")
			fmt.Printf("[pending2] %+v\n", key)
		}
	}
}
