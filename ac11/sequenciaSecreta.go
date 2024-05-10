package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	sequencia := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&sequencia[i])
	}

	maximoMarcado := maximoMarcadoSequencia(sequencia)

	fmt.Println(maximoMarcado)
}

func maximoMarcadoSequencia(sequencia []int) int {
	maximoMarcado := 1
	atualMarcado := 1

	for i := 1; i < len(sequencia); i++ {
		if sequencia[i] != sequencia[i-1] {
			atualMarcado++
		} else {
			atualMarcado = 1
		}

		if atualMarcado > maximoMarcado {
			maximoMarcado = atualMarcado
		}
	}

	return maximoMarcado
}
