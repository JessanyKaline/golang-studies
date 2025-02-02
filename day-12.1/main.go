package main

import (
	"fmt"
)


func main() {


	numberEve := make(chan int)
	numberOdd := make(chan int)
	quit := make(chan bool)

	go sendNumbers(numberEve, numberOdd, quit)

	receiveNumbers(numberEve, numberOdd, quit)

}

func sendNumbers(numberEve, numberOdd chan int, quit chan bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			numberEve <- i
			fmt.Println("Send number even: ", i)
		} else {
			numberOdd <- i
			fmt.Println("Send number odd: ", i)
		}
	}
	close(numberEve)
	close(numberOdd)
	quit <- true
}

func receiveNumbers(numberEve, numberOdd chan int, quit chan bool) {
	for {
		select {
		case v, ok := <-numberEve:
			if !ok {
				fmt.Println("Channel numberEve closed")
			} else {
				fmt.Println("Receive number even: ", v)
			}
		case v, ok := <-numberOdd:
			if !ok {
				fmt.Println("Channel numberOdd closed")
			} else {
				fmt.Println("Receive number odd: ", v)
			}
		case <-quit:
			return 
			
		}
	}
}
