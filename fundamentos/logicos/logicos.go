package main

import "fmt"

func comprar(trab1 bool, trab2 bool) (bool, bool, bool) {
  comprarTV50 := trab1 && trab2 // E
  comprarTV32 := trab1 != trab2 // Ou exclusivo
  comprarSorvete := trab1 || trab2 // OU

  return comprarTV50, comprarTV32, comprarSorvete
}

func main() {
  tv50, tv32, sorvete := comprar(true, true)

  fmt.Printf("TV 50: %t\nTV 32: %t\nSorvete: %t\nSaudável: %t\n", tv50, tv32, sorvete, !sorvete)
}
