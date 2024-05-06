package main

import "fmt"

type Node struct {
    value int
    next  *Node
}

type Fila struct {
    frente *Node
    traseira  *Node
}

func (f *Fila) inserir(valor int) {
    novoNode := &Node{valor, nil}

    if f.traseira == nil {
        f.frente = novoNode
        f.traseira = novoNode
    } else {
        f.traseira.next = novoNode
        f.traseira = novoNode
    }

    fmt.Println("Elemento inserido:", valor)
}

func (f *Fila) remover() int {
    if f.frente == nil {
        fmt.Println("A fila está vazia")
        return -1
    }

    valor := f.frente.value
    f.frente = f.frente.next

    if f.frente == nil {
        f.traseira = nil
    }

    return valor
}

func (f *Fila) percorrer() {
    if f.frente == nil {
        fmt.Println("A fila está vazia")
        return
    }

    atual := f.frente
    for atual != nil {
        fmt.Println(atual.value)
        atual = atual.next
    }
}

func main() {
    fila := Fila{}

    fila.inserir(1)
    fila.inserir(2)
    fila.inserir(3)

    fmt.Println("Elementos na fila:")
    fila.percorrer()

    fmt.Println("Elemento removido:", fila.remover())

    fmt.Println("Elementos na fila:")
    fila.percorrer()

    fila.inserir(4)

    fmt.Println("Elementos na fila:")
    fila.percorrer()

    fmt.Println("Elemento removido:", fila.remover())

    fmt.Println("Elementos na fila:")
    fila.percorrer()
}
