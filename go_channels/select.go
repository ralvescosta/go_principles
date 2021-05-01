package main

import "fmt"

func Select_ReceveOnly() {

	chA := make(chan int)
	chB := make(chan int)

	x := 500

	go func(x int) {
		for i := 0; i < x; i++ {
			chA <- i
		}
	}(x / 2)

	go func(x int) {
		for i := 0; i < x; i++ {
			chB <- i
		}
	}(x / 2)

	for i := 0; i < x; i++ {
		select {
		case v := <-chA:
			fmt.Println("Channel A:", v)
		case v := <-chB:
			fmt.Println("Channel B:", v)
		}
	}
}

func Select_ReceveAndSend() {
	ch := make(chan int)
	quit := make(chan bool)

	go receiveQuit(ch, quit)

	sendToChannel(ch, quit)
}

func sendToChannel(ch chan<- int, quit <-chan bool) {
	ulalala := 1
	for {
		select {
		case ch <- ulalala:
			ulalala++
		case <-quit:
			return
		}
	}
}

func receiveQuit(ch <-chan int, quit chan<- bool) {
	for i := 0; i < 500; i++ {
		fmt.Println("Recebido: ", <-ch)
	}
	quit <- true
}

func Select_Odd_Even() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go sendNumbers(even, odd, quit)

	receive(even, odd, quit)
}

func sendNumbers(even, odd chan<- int, quit chan<- bool) {
	amount := 500

	for i := 0; i < amount; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
	quit <- true
}

func receive(even, odd <-chan int, quit chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("The number,", v, "Is Even")
		case v := <-odd:
			fmt.Println("The number,", v, "Is Odd")
		case <-quit:
			close(quit)
			return
		}
	}
}
