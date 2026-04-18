// Crie um tipo struct "pessoa" que contenha os campos:
// nome
// sobrenome
// idade
// Crie um método para "pessoa" que demonstre o nome completo e a idade;
// Crie um valor de tipo "pessoa";
// Utilize o método criado para demonstrar esse valor.

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) nomecompleto() string {
	return p.nome + " " + p.sobrenome
}
func main() {

	carlos := pessoa{"Carlos", "Pereira", 39}

	fmt.Println("Meu nome é:", carlos.nomecompleto())

}
