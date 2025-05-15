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

var cwd string
var (
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
)
var (
	CloseProgramRelativePath = "sound\\closeProgram.mp3"
	StartProgramRelativePath = "sound\\startProgram.mp3"
	PendingRelativePath      = "sound\\pending.mp3"
	EndPendingRelativePath   = "sound\\endPending.mp3"
)

func init() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}
}

func main() {
	fmt.Println("Current directory is: ", cwd)
	// Detect language input. if current input is Chinese, then reminds user.

	StartProgramPath := filepath.Join(cwd, StartProgramRelativePath)
	go playSound(StartProgramPath)
	fmt.Println("======= Listening for pressing by keyboard & mouse =======")
	keyChan := hook.Start()
	defer hook.End()
	ctx, cancel = context.WithCancel(context.Background())
	mainLoop(keyChan)
}

func generateVills() {
	time.Sleep(83 * time.Millisecond)
	robotgo.KeyTap("d")
	time.Sleep(75 * time.Millisecond)
	robotgo.KeyTap("d")
}

func mainLoop(keyChan chan hook.Event) {
	for key := range keyChan {
		if key.Kind == hook.KeyHold && key.Rawcode == 187 { // listen "=" key press (my new 'Capslock')
			// cause we might be dancing our cav (zxzx), so we need to stop function - cavDancin first
			cancel()
			wg.Wait()
			ctx, cancel = context.WithCancel(context.Background())
			// start to generate villagers from TC
			go generateVills()
		} else if key.Kind == hook.KeyDown && key.Rawcode == 67 { // listen for "c" or "C"
			wg.Add(1)
			go cavDancing(ctx)
		} else if key.Kind == hook.KeyDown && key.Rawcode == 13 && key.Mask == 2 { // pressing Enter + Ctrl
			CloseProgramPath := filepath.Join(cwd, CloseProgramRelativePath)
			go playSound(CloseProgramPath)
			fmt.Println("Enter + Ctrl pressed. Exiting program.")
			cancel()
			wg.Wait()
			time.Sleep(103 * time.Millisecond)
			os.Exit(0)
		} else if key.Kind == hook.KeyDown && key.Rawcode == 13 { // pressing Enter
			cancel()
			ctx, cancel = context.WithCancel(context.Background())
			fmt.Println("Enter pressed. Pending program.")
			PendingPath := filepath.Join(cwd, PendingRelativePath)
			go playSound(PendingPath)
			time.Sleep(53 * time.Millisecond)
			robotgo.KeyTap(";")
			signal := make(chan bool)
			go pending(keyChan, signal)
			// to pending the auto-typing function
			<-signal
			// fmt.Println("===== open the pending channel =======")
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

func playSound(absPath string) {
	cmd := exec.Command("ffplay", "-nodisp", "-autoexit", "-loglevel", "quiet", absPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error playing sound:", err)
	}
}

func pending(keyChan chan hook.Event, signal chan bool) {
	for key := range keyChan {
		if key.Kind == hook.KeyHold && key.Rawcode == 187 { // "if press capslock"
			StartProgramPath := filepath.Join(cwd, StartProgramRelativePath)
			go playSound(StartProgramPath)
			// fmt.Printf("[pending1] %+v\n", key)
			// to open the pending lock
			signal <- true
			return
		}
		// else if key.Kind == hook.KeyHold || key.Kind == hook.KeyDown {
		// 	// fmt.Println("====== pending2 ======")
		// 	fmt.Printf("[pending2] %+v\n", key)
		// }
	}
}

func playSound_relative(relativePath string) {
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
