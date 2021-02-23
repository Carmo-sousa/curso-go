package main

import (
	"fmt"
	"strconv"
)

func main() {
	// float64
	x := 2.4
	// int
	y := 2
	result := x / float64(y)

	fmt.Println("Conversão de tipos em Go!")
	fmt.Println(result)

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuídado...
	fmt.Println("Tete " + string(97))

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 23)

	b, _ := strconv.ParseBool("true")

	if b {
		fmt.Println("Verdadeiro...")
	}
}
