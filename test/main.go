package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create a channel to receive keyboard input
	inputCh := make(chan string)

	// Start a goroutine to read keyboard input and send to the channel
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text() // Get the input line
			inputCh <- input        // Send input to channel
		}
		// If scanner encounters an error or EOF, close the channel
		close(inputCh)
	}()

	// Main goroutine listens for messages on the channel
	for input := range inputCh {
		fmt.Printf("Received input: %s\n", input)
		if input == "exit" {
			fmt.Println("Exiting program...")
			break
		}
	}
}
