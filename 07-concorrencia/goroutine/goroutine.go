package main

import (
	"fmt"
	"time"
)

// Goroutines
func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "Pq não fala comigo?", 3)
	// fale("João", "Só posso falar depois de você!", 1)

	// go fale("Maria", "Ei...", 500)
	// go fale("João", "Oi...", 500)

	go fale("Rômulo", "Oi", 10)
	fale("Maria", "Pq não fala comigo?", 5)

	fmt.Println("Fim!")
	// time.Sleep(time.Second * 5)
}
