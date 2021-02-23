package main

import "fmt"

func main() {
	aprovados := []string{"Rômulo", "Rodrigo", "Rogério"}

	// Desestruturando um slice com '...'
	imprimirAprovados(aprovados...)
}

func imprimirAprovados(aprovados ...string) {
	fmt.Println("===== Lista de aprovados =====")

	for i, aluno := range aprovados {
		fmt.Printf("%d - %s\n", i+1, aluno)
	}
}
