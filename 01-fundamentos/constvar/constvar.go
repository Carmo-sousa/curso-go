package main

import (
  "fmt"
  "math"
)

func main() {

  // Constate em Go
  const PI float64 = 3.1415

  // Uma das formas de declarar uma varíavel.
  var raio = 3.2

  // Forma curta de declarar uma varíavel.
  area := PI * math.Pow(raio, 2)

  const (
    a = 1
    b = 2
    c = 3
  )

  var d, e bool = true, false

  fmt.Printf("PI: %f\nRaio: %f\nArea: %f\n\n", PI, raio, area)
  fmt.Printf("%d %d %d\n\n", a, b, c)
  fmt.Printf("%t %t\n", d, e)
}
