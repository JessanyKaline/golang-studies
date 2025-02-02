package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a new channel that merges the output of two work functions
	c := converges(work("John"), work("Jessany"))
	
	// Print the first 16 messages received from the merged channel
	for i := 0; i < 16; i++ {
		fmt.Println(<-c)
	}
}

// work function starts a goroutine that sends messages to a channel
func work(s string) chan string {
	channel := make(chan string)

	// Start a goroutine that sends messages to the channel
	go func(s string, c chan string) {
		for i := 1; ; i++ {
			// Send a formatted string to the channel
			c <- fmt.Sprintf("Function %v says: %v", s, i)
			// Sleep for 500 milliseconds
			time.Sleep(time.Millisecond * 500)
		}
	}(s, channel)

	return channel
}

// converges function merges two input channels into a single output channel
func converges(c1, c2 chan string) chan string {
	newChan := make(chan string)

	// Start a goroutine that forwards messages from c1 to newChan
	go func() {
		for {
			newChan <- <-c1
		}
	}()

	// Start a goroutine that forwards messages from c2 to newChan
	go func() {
		for {
			newChan <- <-c2
		}
	}()

	return newChan
}
