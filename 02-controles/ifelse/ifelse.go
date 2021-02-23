package main

import "fmt"

func main() {
  imprimirResultado(7.5)
  imprimirResultado(6.9)
}

func imprimirResultado(nota float64) {
  if nota >= 7.0 {
    fmt.Println("Aprovado com nota", nota)

  } else {
    fmt.Println("Reprovado com nota", nota)

  }
}

