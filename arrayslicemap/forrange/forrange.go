package main

import "fmt"

func main() {
  numeros := [...]int{1, 2, 3, 4, 5} // compilador conta!

  // O range retorna dois elementos o primeiro é o indíce
  // O sengundo é o elemento
  for i, numero := range numeros {
    fmt.Printf("%d) %d\n", i, numero)
  }
}
