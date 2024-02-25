package main

func MinhaPotencia(base, expoente int) int {
    if expoente == 0 {
        return 1
    }

    resultado := 1
    for i := 0; i < expoente; i++ {
        resultado *= base
    }
    return resultado
}
