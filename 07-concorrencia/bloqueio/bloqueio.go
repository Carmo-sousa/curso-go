package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	time.Sleep(time.Second)
	ch <- 1 // Operação bloqueante
	fmt.Println("Só depois que foi lido!")
}

func main() {
	c := make(chan int) // Canal sem buffer

	go rotina(c)

	fmt.Println(<-c)
	fmt.Println("Foi lido!")
	fmt.Println(<-c)
	fmt.Println("Fim!")
}
