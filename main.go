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
)

func main() {

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}
	fmt.Println("Current directory is: ", cwd)
	// Detect language input. if current input is Chinese, then reminds user.

	go playSound("sound\\startProgram.mp3")
	fmt.Println("======= Listening for pressing by keyboard & mouse =======")
	// keyChan := hook.Start()
	// defer hook.End()
	ctx, cancel = context.WithCancel(context.Background())
	for {
		keyChan := hook.Start()
		mainLoop(keyChan)
		hook.End()
	}
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
		} else if key.Kind == hook.KeyDown && key.Rawcode == 69 { // listen for "e" or "E"
			// fmt.Printf("pressing 'e' %+v\n", key)
			cancel()
			wg.Wait()
			ctx, cancel = context.WithCancel(context.Background())
		} else if key.Kind == hook.KeyDown && key.Rawcode == 13 && key.Mask == 2 {
			fmt.Println("Enter + Ctrl pressed. Exiting program.")
			time.Sleep(103 * time.Millisecond)
			robotgo.KeyTap("*")
			// fmt.Println("Ctrl + Enter pressed. Exiting program.")
			cancel()
			wg.Wait()
			os.Exit(0)
		} else if key.Kind == hook.KeyDown && key.Rawcode == 13 { // pressing Enter
			fmt.Println("Enter pressed. Exiting program.")
			go playSound(CloseProgramPath)
			time.Sleep(103 * time.Millisecond)
			robotgo.KeyTap(";")
			// fmt.Println("Enter pressed. Exiting program END.")

			cancel()

			keyChan2 := hook.Start()
			signal := make(chan bool)
			go pending(keyChan2, signal)
			<-signal
			fmt.Println("===== open the pending channel =======")
			ctx, cancel = context.WithCancel(context.Background())

			// keyChan = hook.Start()
			// hook.End()
			break
			// pendingCh := make(chan string)
			// hook.End()
			// time.Sleep(100 * time.Millisecond) // let goroutines respond to cancel
			// os.Exit(0)
		}
		// }else {
		// 	fmt.Printf("%+v\n", key)
		// }

		// else if key.Kind == hook.KeyHold || key.Kind == hook.KeyDown || key.Kind == hook.KeyUp {
		// 	fmt.Printf("%+v\n", key)
		// }
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

func pending(keyChan2 chan hook.Event, signal chan bool) {
	// keyChan2 := hook.Start()
	for key2 := range keyChan2 {
		if key2.Kind == hook.KeyHold && key2.Rawcode == 187 { // "if press capslock"
			// fmt.Println("[pending1]")
			fmt.Printf("[pending1] %+v\n", key2)
			// hook.End()
			signal <- true
		} else if key2.Kind == hook.KeyHold || key2.Kind == hook.KeyDown {
			// fmt.Println("====== pending2 ======")
			fmt.Printf("[pending2] %+v\n", key2)
		}
	}
}
