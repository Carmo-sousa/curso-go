package main

import "math"

// iniciando uma variavel com letra
// maiúscula ela é PÚBLICA (visivel dentro e fora do pacote)!

// Iniciado com letra minúscula é PRIVADO (visivel apenas dentro do pacote)!

// EX:
//	Ponto - gerará algo público
//	ponto ou _Ponto gerará algo privado

// O Ponto representa uma coordenada no plano cartesiano.
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia é responsável por calcular a distância linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}

