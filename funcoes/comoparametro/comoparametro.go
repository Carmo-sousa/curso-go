package main

import "fmt"

// Passando funções como parametro de outras funções.
func main() {
	result := exec(mult, 3, 3)
	fmt.Println(result)
}

func mult(a, b int) int {
	return a * b
}

// Recebe uma função como arametro.
func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}
