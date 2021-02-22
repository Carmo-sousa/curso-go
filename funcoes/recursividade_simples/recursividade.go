package main

import "fmt"

func main() {
	result := fatorial(6)

	fmt.Println(result)

}

func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1

	default:
		fatorialAnterior := fatorial(n - 1)
		return n * fatorialAnterior
	}
}
