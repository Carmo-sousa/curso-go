package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"gabriel": 1200.00,
			"Guga":    1000.00,
		},
		"J": {
			"José": 18000.00,
		},
		"P": {
			"Pedro": 100000.00,
		},
	}

	fmt.Println(funcsPorLetra)
	delete(funcsPorLetra, "P")
	fmt.Println(funcsPorLetra)

	for letra, funcs := range funcsPorLetra {
		fmt.Println("-------------------------")
		for nome, salario := range funcs {
			fmt.Printf("Letra: %s\n", letra)
			fmt.Printf("Nome: %s\nSalário: %.2f\n\n", nome, salario)
		}
	}
}
