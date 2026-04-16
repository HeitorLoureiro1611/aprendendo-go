package main

import (
	"fmt"
)

// metodo é uma função anexada a um tipo
type pessoa struct {
	nome  string
	idade int
}

// metodo de uma função que recebe o tipo pessoa
func (p pessoa) oibomdia() {
	fmt.Println("bom dia, sou :", p.nome)
}

func main() {
	// objeto criado
	carlos := pessoa{"Carlos", 21}

	// usando um metodo com o objeto
	carlos.oibomdia()
}
