package main

import (
	"encoding/json"
	"fmt"
)

// Mapeamento pro json
type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct to json
	p1 := produto{
		ID:    1,
		Nome:  "Notebook",
		Preco: 7000.00,
		Tags:  []string{"Promoção", "Eletrônico"},
	}
	p1Json, _ := json.Marshal(p1)
	p2, _ := json.MarshalIndent(p1, "", "   ")
	fmt.Println(string(p1Json))
	fmt.Println(string(p2))

	var p3 produto
	jsonString := `{"id":2, "nome":"Caneta", "preco": 1.50,"tags":["promoção", "escrita"]}`
	json.Unmarshal([]byte(jsonString), &p3)
	fmt.Println(p3.Tags[1])
}
