package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func WithAtomic() {
	var counter int64 = 0
	amountOfGoRoutine := 10

	var waitGroup sync.WaitGroup

	waitGroup.Add(amountOfGoRoutine)

	for i := 0; i < amountOfGoRoutine; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()
	fmt.Println(counter)
}
