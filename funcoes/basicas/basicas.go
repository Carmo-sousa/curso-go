package main

import "fmt"

func main() {
	f1()
	f2("Param 1", "Param 2")
	fmt.Println(f3())
	fmt.Println(f4("Param 1", "Param 2"))
	fmt.Println(f5())
}

// Não retorna valor.
func f1() {
	fmt.Println("F1")
}

// Recebe dois parametros, mas não retorna nada.
func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

// Não recebe parametros, mas retorna uma string
func f3() string {
	return "F3"
}

// Recebe dois valores e retorna uma uma string
func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

// Retorna múltiplos valores.
func f5() (string, string) {
	return "F5: Param 1", "F5: Param 2"
}
