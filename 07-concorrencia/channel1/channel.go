package main

import "fmt"

func main() {
	// Criando um channel
	ch := make(chan int, 1)

	// Enviando dados para um canal (escrita)
	ch <- 1

	// Recebendo os dados do canal (leitura)
	<-ch

	ch <- 2

	fmt.Println(<-ch)

}
