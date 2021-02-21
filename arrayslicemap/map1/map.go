package main

import "fmt"

func main() {
	// Map - dict[python]
	// Maps devem ser inicializados!
	// var aprovados map[int]string
	aprovados := make(map[int]string)

	aprovados[123456789] = "Mária"
	aprovados[524524253] = "Pedro"
	aprovados[781781349] = "Ana"

	// fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("CPF: %d\nNome: %s\n", cpf, nome)
	}

	// Acessando um valor
	fmt.Println(aprovados[123456789])

	// Deletando um valor
	delete(aprovados, 781781349)
	fmt.Println(aprovados)
}
