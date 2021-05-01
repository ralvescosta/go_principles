package main

import "fmt"

func DirectionalChannels() {
	ch := make(chan int)

	go sendOnly(ch)

	receiveOnly(ch)
}

func sendOnly(ch chan<- int) {
	ch <- 45
}

func receiveOnly(ch <-chan int) {
	fmt.Println(<-ch)
}
