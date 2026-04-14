package main

import (
	"fmt"
)

// declaração do tipo struct
type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario float64
}

func main() {

	// uso do tipo struct
	cliente1 := cliente{
		nome:      "João",
		sobrenome: "Pereira",
		fumante:   false,
	}
	// outra forma de usar um struct
	cliente2 := cliente{"joana", "pereira", true}

	pessoa1 := pessoa{
		nome:  "Jorge",
		idade: 30,
	}

	pessoa2 := profissional{
		pessoa: pessoa{
			nome:  "Junior",
			idade: 22,
		},
		titulo:  "Professor",
		salario: 2300.0,
	}
	fmt.Println(pessoa1.nome)
	fmt.Println(pessoa2)
	fmt.Println(cliente1)
	fmt.Println(cliente2)
}
