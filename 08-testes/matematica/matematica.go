package matematica

import (
	"fmt"
	"strconv"
)

// Media retorna a média dos valores passados.
func Media(nuns ...float64) float64 {
	total := 0.0

	for _, num := range nuns {
		total += num
	}
	media := total / float64(len(nuns))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)

	return mediaArredondada
}
