package main

import (
	"fmt"
	"time"

	"github.com/Carmo-sousa/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// Multiplexar - misturar (menssagens) em um canal.
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	ch := make(chan string)
	go encaminhar(entrada1, ch)
	go encaminhar(entrada2, ch)

	return ch
}

func main() {
	ch := juntar(
		html.Titulo("https://www.udemy.com/", "https://stackoverflow.com/"),
		html.Titulo("https://www.digitalocean.com/", "https://www.geeksforgeeks.org/"),
	)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	time.Sleep(time.Second)
}
