package main

import (
  "fmt"
)

func main() {
  fmt.Println(getType(10))
  fmt.Println(getType("Oi"))
  fmt.Println(getType(12.5))
}

func getType(i interface{}) string {
  switch i.(type) {
    case int:
      return "Inteiro"
    case float32, float64:
      return "Real"
    case string:
      return "String"
    case func():
      return "Função"
    default:
      return "Não sei..."
  }
}
