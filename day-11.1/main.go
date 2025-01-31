package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var contador int
const qntGoroutines = 500

// Person struct represents a person with a name
type Person struct {
	name string
}

// Human interface defines a speak method
type Human interface {
	speak()
}

// speakSomething calls the speak method on a Human
func speakSomething(h Human){
	h.speak()
}

func main() {
	exercise() // Call the exercise function
	
	p1 := Person{"Jessany"}

	p1.speak() // Call the speak method on p1

	speakSomething((&p1)) // Call speakSomething with p1

	exerciseCreateGoroutines(qntGoroutines)
	wg.Wait()
	fmt.Println("Goroutines:\t", qntGoroutines, "\nContador:\t", contador)

}

// It is necessary to wait for each goroutine, so use defer and wg.Done()
func exercise() {
	wg.Add(3) // Add 3 to the WaitGroup counter
	go func() {
		defer wg.Done() // Decrement the counter when the goroutine completes
		go func() {
			defer wg.Done() // Decrement the counter when the goroutine completes
			fmt.Println("First Print!")
		}()
		go func() {
			defer wg.Done() // Decrement the counter when the goroutine completes
			fmt.Println("Second Print!")
		}()

	}()
	wg.Wait() // Wait for all goroutines to complete
}

// speak method for the Person type
func (p *Person) speak() {
	fmt.Println(p.name, "is speaking")
}

func exerciseCreateGoroutines(i int){
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			v := contador
			v++
			runtime.Gosched()
			contador = v
			mu.Unlock()
		}()
	}

}
