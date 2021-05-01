package main

import (
	"fmt"
)

func RangeChannel() {
	ch := make(chan int)

	go increment(10, ch)

	for c := range ch {
		fmt.Println(c)
	}
}

func increment(t int, ch chan<- int) {
	for i := 0; i < t; i++ {
		ch <- i * 2
	}
	close(ch)
}
