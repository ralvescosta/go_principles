package main

import (
	"fmt"
	"runtime"
	"sync"
)

func WithMutex() {
	counter := 0
	amountOfGoRoutine := 10

	var waitGroup sync.WaitGroup
	var mutex sync.Mutex

	waitGroup.Add(amountOfGoRoutine)

	for i := 0; i < amountOfGoRoutine; i++ {
		go func() {
			mutex.Lock()
			v := counter

			runtime.Gosched()

			v++
			counter = v

			mutex.Unlock()
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()
	fmt.Println(counter)
}
