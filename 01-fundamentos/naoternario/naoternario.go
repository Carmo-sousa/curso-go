package main

import "fmt"

// Não tem operadores ternarios em Go!

func main() {
  fmt.Println(obterResultado(6.2))
}

func obterResultado(nota float64) string {
  if nota >= 6.0 {
    return "Aprovado!"
  }

  return "Reprovado!"
}
