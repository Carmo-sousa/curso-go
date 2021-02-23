package main

import "fmt"

func main() {
  i := 1

  // Go não tem operações aritiméticas de ponteiros
  var p *int = nil

  p = &i // Pegando o endereço da variável i.
  *p++

  fmt.Println(i, *p)
}
