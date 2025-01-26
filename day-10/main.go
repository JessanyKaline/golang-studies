package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//The package sync provides basic synchronization primitives such as mutual exclusion locks.
var wg sync.WaitGroup

func main() {
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("Go Routines:", runtime.NumGoroutine())
	//The function aboutGoRuntime() is called to demonstrate the features of the Go runtime.
	aboutGoRuntime()

	fmt.Println("Go Routines:", runtime.NumGoroutine())

}

func aboutGoRuntime() {
	// Go runtime is responsible for managing the execution of Go programs.
	// It provides a lot of features that make Go programs efficient and fast.
	// Some of the features are:
	// 1. Goroutines: Goroutines are lightweight threads that are managed by the Go runtime.
	//    They are used to run concurrent tasks in a Go program.
	// 2. Garbage Collection: Go runtime has a garbage collector that automatically frees up memory that is no longer in use.
	//    This helps in managing memory efficiently and prevents memory leaks.
	// 3. Scheduler: Go runtime has a scheduler that is responsible for scheduling Goroutines on the available CPU cores.
	//    It ensures that Goroutines are executed efficiently and fairly.
    wg.Add(2)
	fmt.Println("Go Routines in aboutGoRunTime:", runtime.NumGoroutine())
	go func1()
	go func2()
	wg.Wait()
	

}

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Println("func1:", i)
		time.Sleep(10)
	}
	wg.Done()

}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("fun2:", i)
		time.Sleep(10)
	}
	wg.Done()
}
