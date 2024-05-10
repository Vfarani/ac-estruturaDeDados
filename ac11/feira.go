package main

import (
	"fmt"
)

type Produto struct {
	Nome  string
	Preco float64
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		produtos := make(map[string]float64)
		for j := 0; j < m; j++ {
			var nome string
			var preco float64
			fmt.Scan(&nome, &preco)
			produtos[nome] = preco
		}

		var p int
		fmt.Scan(&p)

		var total float64
		for j := 0; j < p; j++ {
			var nome string
			var quantidade int
			fmt.Scan(&nome, &quantidade)
			if preco, ok := produtos[nome]; ok {
				total += float64(quantidade) * preco
			}
		}

		fmt.Printf("R$ %.2f\n", total)
	}
}
