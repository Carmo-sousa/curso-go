package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Romulo":  10000.00,
		"Rodrigo": 10000.00,
	}

	funcsESalarios["Rogerio"] = 11000.00
	fmt.Println(funcsESalarios)

	for nome, salario := range funcsESalarios {
		fmt.Println("---------------------------------------")
		fmt.Printf("Nome: %s\nSalário: %.2f\n", nome, salario)
		fmt.Println("---------------------------------------")
	}
}
