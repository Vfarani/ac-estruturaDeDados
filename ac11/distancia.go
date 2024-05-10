package main

import (
    "fmt"
)

func main() {
    var distancia int
    fmt.Scanln(&distancia)

    tempo := distancia * 2
    fmt.Printf("%d minutos\n", tempo)
}
