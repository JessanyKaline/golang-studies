package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	counter := 0
	totalGoRoutines := 100

	var wg sync.WaitGroup

	wg.Add(totalGoRoutines)
    // The function runtime.Gosched() yields the processor, allowing other Goroutines to run, only if there are other Goroutines to run.
	var mu sync.Mutex

	for i := 0; i < totalGoRoutines; i++ {
		go func() {
			//The function mu.Lock() locks the mutex mu. If the mutex is already locked, the calling Goroutine blocks until the mutex is unlocked.
			mu.Lock()
			v := counter
			//The function runtime.Gosched() yields the processor, allowing other Goroutines to run.
			runtime.Gosched()
			v++
			counter = v
			//The function mu.Unlock() unlocks the mutex mu. If the mutex is not locked, the function panics.
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Go Routines:", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Go Routines:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
	
}