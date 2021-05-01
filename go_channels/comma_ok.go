package main

import "fmt"

func CommaOK() {
	ch := make(chan int)

	go func() {
		ch <- 42
		close(ch)
	}()

	value, ok := <-ch
	fmt.Println(value, ok)

	value, ok = <-ch
	fmt.Println(value, ok)
}
