package main

import "fmt"

func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(4000))
}

func obterValorAprovado(numero int) int {
	// Adiciona o comando em uma pilha de execução LIFO
	// que vai ser chamada assim que a função retornar
	defer fmt.Println("Fim!")

	if numero >= 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}

	fmt.Println("Valor baixo...")
	return numero
}
