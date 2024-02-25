package main

import (
    "fmt"
)

func main() {
	// Exemplo da função VerificaParidade
    fmt.Println("O número 10 é", VerificaParidade(10))
    fmt.Println("O número 7 é", VerificaParidade(7))

    // Exemplos da função MinhaPotencia
    fmt.Println("2 elevado a 3 é", MinhaPotencia(2, 3))
    fmt.Println("5 elevado a 0 é", MinhaPotencia(5, 0))

    // Exemplo da função ConverteCelsiusParaFahrenheit
    celsius := 25.0
    fahrenheit := ConverteCelsiusParaFahrenheit(celsius)
    fmt.Printf("%.2f graus Celsius equivalem a %.2f graus Fahrenheit\n", celsius, fahrenheit)

    // Exemplo da função CalculaMedia
    media := CalculaMedia(10.0, 15.0, 20.0)
    fmt.Println("Média:", media)


}
