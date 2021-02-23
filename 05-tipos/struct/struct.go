package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}

	produto2 := produto{"Notebook", 1500.00, 0.10}

	fmt.Println(produto1)
	fmt.Println(produto1.precoComDesconto())
	fmt.Println(produto2)
	fmt.Println(produto2.precoComDesconto())
}
