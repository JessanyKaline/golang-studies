package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)

	go rangeChannelLoop(2, c)
	printsChannel(c)

	selectChannel()

}

func aboutChannel() {

	//Create channel in go
	//Send and Receive
	//When channel is send, is not possible convert receive
	c := make(chan int)

	//When send something for channel use chan <- x
	//The channel only send
	sendChannel(c)

	// When receive something for channel use <-chan
	//The channel only receive
	receiveChannel(c)
	wg.Wait()
}

func sendChannel(c chan int) {
	c <- 0 //Sending value 0 for channel
}

func receiveChannel(c chan int) {
	value := <-c //Receiving value for channel
	fmt.Println(value)
}

func rangeChannelLoop(i int, s chan int) {
	for j := 0; j < i; j++ {
		s <- j
	}

	//Here close the channel
	close(s)
}

func printsChannel(r <-chan int) {
	for v := range r {
		fmt.Println("Receive Channel: ", v)
	}

}

func selectChannel() {
	//Select is used to wait multiple channel operations
	//It blocks until one of its cases can run, then it executes that case.
	//It chooses one at random if multiple are ready.
	//If none are ready, it blocks until at least one is ready.
	c1 := make(chan int)
	c2 := make(chan int)

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			c1 <- 1
		}
		close(c1)
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			c2 <- 2
		}
		close(c2)
	}()

	//In use of select, is necessary to use for loop
	//
	for {
		select {
		case v, ok := <-c1:
			if !ok {
				c1 = nil
			} else {
				fmt.Println("Receive from c1: ", v)
			}
		case v, ok := <-c2:
			if !ok {
				c2 = nil
			} else {
				fmt.Println("Receive from c2: ", v)
			}
		}
		if c1 == nil && c2 == nil {
			break
		}
	}

	wg.Wait()

}

