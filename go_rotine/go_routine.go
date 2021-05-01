package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func GoRoutine() {
	fmt.Println("Amount of CPU: ", runtime.NumCPU())
	fmt.Println("Amount of GoRoutine", runtime.NumGoroutine())

	wg.Add(2)

	go func1()
	go func2()

	fmt.Println("Amount of GoRoutine", runtime.NumGoroutine())

	wg.Wait()
}

func func1() {
	for i := 0; i < 100; i++ {
		fmt.Println("Func 1: ", i)
		time.Sleep(50)
	}
	wg.Done()
}

func func2() {
	for i := 0; i < 100; i++ {
		fmt.Println("Func 2: ", i)
		time.Sleep(30)
	}
	wg.Done()
}
