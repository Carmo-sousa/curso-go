package main

import (
	"fmt"
	"time"
)

// Channel - é a forma de comunicação entre goroutines
// channel também é um tipo!

func doisTresQuatroVezes(base int, ch chan int) {
	time.Sleep(time.Second)
	ch <- 2 * base // Escrevendo

	time.Sleep(time.Second)
	ch <- 3 * base

	time.Sleep(time.Second * 3)
	ch <- 4 * base
}

func main() {
	ch := make(chan int)
	go doisTresQuatroVezes(2, ch)

	a, b := <-ch, <-ch // Lendo os dados do canal

	fmt.Println(a, b)

	fmt.Println(<-ch)
}
