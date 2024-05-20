package main

import "fmt"

func main() {
	var k int
	posicoes := []int{1, 3, 5, 10, 25, 50, 100}

	fmt.Print("Digite um número: ")
	fmt.Scanln(&k)

	for _, posicao := range posicoes {
		if k <= posicao {
			fmt.Printf("Top %d\n", posicao)
			break
		}
	}
}
