package main

import (
	"fmt"
	"math"
)

func main() {
	var n, h, l, c float64

	for {
		fmt.Print("Digite o número de triângulos: ")
		if _, err := fmt.Scan(&n); err != nil {
			break
		}

		fmt.Print("Digite a altura (h), a base (c) e o comprimento (l): ")
		fmt.Scan(&h, &c, &l)

		hipotenusa := math.Hypot(h, c) 
		area := (n * hipotenusa * l) / 10000.0

		fmt.Printf("Área total: %.4f m²\n", area)
	}
}
