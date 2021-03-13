package main

import (
	"fmt"
	"time"
)

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}

	}
	return true
}

func primos(num int, ch chan int) {
	inicio := 2
	for i := 0; i < num; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				ch <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	close(ch)
}

func main() {
	ch := make(chan int, 30)

	go primos(cap(ch), ch)

	for primo := range ch {
		fmt.Printf("%d ", primo)
	}
	fmt.Println("Fim!!")
}
