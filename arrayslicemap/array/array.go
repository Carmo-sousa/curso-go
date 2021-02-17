package main

import "fmt"

func main() {
  /* Arrays
    1. São homogêneos (mesmo tipo)
    2. São estaticos (tamanho fixo)
  */

  var notas [3]float64
  fmt.Println(notas)

  notas[0] = 7.8
  notas[1] = 6.0
  notas[2] = 10.0

  fmt.Println(notas)

  total := 0.0

  for i := 0; i < len(notas); i++ {
    total += notas[i]
  }

  media := total / float64(len(notas))

  fmt.Printf("Media: %.2f\n", media)
}
