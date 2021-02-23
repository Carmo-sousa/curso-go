package main

import "fmt"
import "reflect"

func main() {
  a1 := [3]int{1, 2, 3} // Array (Tamanho fixo)
  s1 := []int{1, 2, 3} // Slice (Tamanho variável)

  fmt.Println(a1, reflect.TypeOf(a1), s1, reflect.TypeOf(s1))

  a2 := [5]int{1, 2, 3, 4, 5}

  for i, numero := range a2 {
    fmt.Printf("%d) %d\n", i, numero)
  }

  s2 := a2[1:3]
  fmt.Println(s2)

  s3 := a2[2: ]
  fmt.Println(s3)
}
