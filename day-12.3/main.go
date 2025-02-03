package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Create two channels for communication between goroutines
	chan1 := make(chan int)
	chan2 := make(chan int)

	// Number of concurrent functions to run
	functions := 2

	// Start a goroutine to send values to chan1
	go send(10, chan1)

	// Start another goroutine to process values from chan1 and send results to chan2
	go otherFunction(functions, chan1, chan2)

	// Receive and print values from chan2
	for v := range chan2 {
		fmt.Println(v)
	}
}

//  sends a sequence of integers to the channel and then closes it
func send(value int, c chan int) {
	for i := 0; i < value; i++ {
		c <- i
	}
	close(c)
}

// otherFunction starts multiple goroutines to process values from c1 and send results to c2
func otherFunction(functions int, c1 chan int, c2 chan int) {
	var wg sync.WaitGroup

	// Start 'functions' number of goroutines
	for i := 0; i < functions; i++ {
		wg.Add(1)
		go func() {
			// Process each value from c1 and send the result to c2
			for n := range c1 {
				c2 <- work(n)
			}
			wg.Done()
		}()
	}
	// Wait for all goroutines to finish
	wg.Wait()
	// Close c2 after all processing is done
	close(c2)
}

// work simulates a time-consuming task by sleeping for 1 second and then returning the input value
func work(n int) int {
	time.Sleep(time.Second)
	return n
}

// convergência em Go (fan-in/fan-out):

// Padrão de Convergência (Fan-in/Fan-out)
// Fan-out: Distribuição do trabalho para múltiplas goroutines
// Fan-in: Convergência dos resultados em um único canal