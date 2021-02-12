package main

import (
  "fmt"
  "io"
  "os"
)

func main() {
  lang := "Go"
  var name string
  var age int
  // Imprime na tela e adiciona uma quebra de linha no final.
  fmt.Println("Hello!")

  fmt.Printf("Hello %s\n\n", lang)
  fmt.Print("Entre com seu nome: ")
  fmt.Scan(&name)
  fmt.Print("Informe sua idade: ")
  fmt.Scanf("%d", &age)

  fmt.Printf("Olá %s sua idade é %d!\n\n", name, age)

  s := fmt.Sprint("Olá, ", lang)
  io.WriteString(os.Stdout, s)
}

