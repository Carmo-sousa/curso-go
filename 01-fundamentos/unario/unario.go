package main

import "fmt"

func main() {
  /*
    Em Go não é permitido o incremento ou decremento de uma variável
    dentro de uma comparação

    Ex:
      x++ == x
  */
  x := 1
  y := 2

  // Operador unario de incremento...
  x++ // x += 1 ou x = x + 1
  fmt.Println(x)

  y-- // y -= 1 ou y = y - 1
  fmt.Println(y)
}
