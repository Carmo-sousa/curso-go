package main

import (
	"fmt"
	"time"
)

func falar(pessoa string) <-chan string {
	ch := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			ch <- fmt.Sprintf("%s falando: %d", pessoa, i)
		}
	}()
	return ch
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	ch := make(chan string)

	go func() {
		// ! Tem a mesma função que um while em python.
		for {
			select {
			case s := <-entrada1:
				ch <- s
			case s := <-entrada2:
				ch <- s
			}
		}
	}()

	return ch
}

func main() {
	ch := juntar(
		falar("João"),
		falar("Maria"),
	)

	fmt.Println(<-ch, <-ch)
	fmt.Println(<-ch, <-ch)
	fmt.Println(<-ch, <-ch)
}
