package main

import "fmt"

func main() {
	fmt.Printf("Media: %.2f\n", media(2.0, 5.5, 7.2, 10.0))
	fmt.Printf("Media: %.2f\n", media())
}

// Função variática.
// Pode receber um número indefinido de paramentros.
func media(numeros ...float64) float64 {
	total := 0.0

	for _, valor := range numeros {
		total += valor
	}

	if total == 0.0 {
		return total
	}

	return total / float64(len(numeros))
}
