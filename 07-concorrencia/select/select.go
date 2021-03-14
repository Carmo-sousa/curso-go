package main

import (
	"fmt"
	"time"

	"github.com/Carmo-sousa/html"
)

func oMaisRapido(url1, url2, url3 string) string {
	ch1 := html.Titulo(url1)
	ch2 := html.Titulo(url2)
	ch3 := html.Titulo(url3)

	// Estrutura de controle específica para concorrência.
	select {
	case t1 := <-ch1:
		return t1
	case t2 := <-ch2:
		return t2
	case t3 := <-ch3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"
		// default:
		// 	return "Sem resposta ainda!"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.google.com.br/",
		"https://www.geeksforgeeks.org/",
		"https://stackoverflow.com/",
	)

	fmt.Println(campeao)
}
