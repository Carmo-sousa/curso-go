package main

import "fmt"

func main() {
	result1, _ := fatorial(6)
	_, err := fatorial(-1)

	fmt.Println(result1)

	if err != nil {
		fmt.Println(err)
	}
}

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("número inválido: %d", n)

	case n == 0:
		return 1, nil

	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}
