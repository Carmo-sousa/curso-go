package main

import (
	"fmt"
	"math"
	"reflect"
)

/*
Tipos das variaveis em Go!

Inteiros
  Sem sinal (só positivo)... uint8 uint16 uint32 uint64
  Com sinal... int8 int16 int32 int64

  boolean


  Por padrão o tipo literal do ponto fluante vai ser sempre float64.
*/

func main() {
	var b byte = 255
	var i rune = 'a' // Representa um mapeamento da tabela unicode. rune == float32
	maxInt := math.MaxInt64
	start := true
	text := `
Uma string de muitas linhas!
Mais uma
Mais uma
	`

	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))
	fmt.Println("O byte é", reflect.TypeOf(b))
	fmt.Println("Tamanho maximo do int64", maxInt)
	fmt.Println("O rune é", reflect.TypeOf(i))
	fmt.Println("Valor do i", i)
	fmt.Println("Valor booleano", reflect.TypeOf(start))
	fmt.Println(text)

}
