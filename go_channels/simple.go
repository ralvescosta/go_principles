package main

import "fmt"

func FirstChannel() {
	ch := make(chan int)

	go func() {
		ch <- 32
	}()

	fmt.Println(<-ch)
}
