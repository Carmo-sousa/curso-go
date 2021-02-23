package main

import "fmt"
import "time"

/*
  Em Go não tem o while
*/

func main() {
  cont := 0

  // For com a variavel de controle fora do escopo dele.
  for cont <= 10 {
    fmt.Println(cont)
    cont++
  }

  // For convencional.
  for i := 0; i <= 20; i += 2 {
    fmt.Println(i)
  }

  for i := 1; i <= 10; i++ {
    if i%2 == 0 {
      fmt.Print(i, " Par! ")
    } else {
      fmt.Print(i, " Impar! ")
    }
  }

  // For fazendo a função do while.
  for {
    fmt.Println("Sempre...")
    time.Sleep(time.Second)
  }

  // Veremos o foreach no capítulo de array!
}
