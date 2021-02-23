package main

import "fmt"

func main() {
	n := 0
	inc1(n)
	fmt.Println(n)
	inc2(&n)
	fmt.Println(n)
}

func inc1(n int) {
	n++
}

// Passando um ponteiro como parametro da função.
func inc2(n *int) {
	// revisão: * é usado para desrefenrinciar, ou seja,
	// ter acesso ao valor no qual o ponteiro aponta.
	*n++
}
