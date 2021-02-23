package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	var total float64
	for _, itens := range p.itens {
		total += itens.preco * float64(itens.qtde)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{1, 2, 12.10},
			item{2, 3, 2.00},
			item{3, 5, 120.00},
		},
	}

	fmt.Printf("Valor total do pedido R$ %.2f\n", pedido.valorTotal())
}
