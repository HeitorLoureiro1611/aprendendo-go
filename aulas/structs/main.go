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

// struct embutido sendo declarado
type profissional struct {
	pessoa
	titulo  string
	salario float64
}

func main() {

	// fortma explicita de preencher os campos
	cliente1 := cliente{
		nome:      "João",
		sobrenome: "Pereira",
		fumante:   false,
	}
	// forma implicita de preencher campos
	cliente2 := cliente{"joana", "pereira", true}

	pessoa1 := pessoa{
		nome:  "Jorge",
		idade: 30,
	}

	// structs embutido sendo preenchido
	pessoa2 := profissional{
		pessoa: pessoa{
			nome:  "Junior",
			idade: 22,
		},
		titulo:  "Professor",
		salario: 2300.0,
	}
	// acessando um item especifico do struct
	// ""objeto"".itemdesejado
	fmt.Println(pessoa1.nome)
	fmt.Println(pessoa2)
	fmt.Println(cliente1)
	fmt.Println(cliente2)
}
