package main

import "fmt"

// Pseudo Herança.

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       // Campos anonimos, ferrari herda os atributos do carro.
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 120
	f.turboLigado = true

	fmt.Printf("O turbo do %s está ligado? %v\n", f.nome, f.turboLigado)
}
