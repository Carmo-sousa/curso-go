package main

import "fmt"

func main() {
  fmt.Println(notaParaConceito(10))
  fmt.Println(notaParaConceito(7))
  fmt.Println(notaParaConceito(4))
  fmt.Println(notaParaConceito(11))
}

func notaParaConceito(n float64) string {
  var nota = int(n)

  switch nota {
    case 10:
      // Por padrão o switch sai quando entra em um case.
      // Para mudar esse comportamento pode ser usado o
      // fallthrough
      fallthrough // Informa que o switch deve continuar.
    case 9:
      return "A"
    case 8, 7:
      return "B"
    case 6, 5:
      return "C"
    case 4, 3:
      return "D"
    case 2, 1, 0:
      return "E"
    default:
      return "Nota inválida!"
  }
}
