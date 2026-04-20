// Crie um struct "pessoa"
// Crie uma função chamada mudeMe que tenha *pessoa como parâmetro.
// Essa função deve mudar um valor armazenado no endereço *pessoa.
// Dica: a maneira "correta" para fazer dereference de um valor em um struct seria (*valor).campo
// Mas consta uma exceção na documentação. Link: https://golang.org/ref/spec#Selectors
// "As an exception, if the type of x is a named pointer type and (*x).f is a valid selector
// expression denoting a field (but not a method),
// → x.f is shorthand for (*x).f." ←
// Ou seja, podemos usar tanto o atalho p1.nome quanto o tradicional (*p1).nome
// Crie um valor do tipo pessoa;
// Use a função mudeMe e demonstre o resultado.

package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func mudeMe(p *pessoa) {
	// é possível usar uma forma melhor apenas com "p.idade += 1" que faz o mesmo trabalho
	// mas (*p).idade += 1 demonstra que estou fazendo uma dereference com ponteiros
	// basicamente boas praticas, não tem importância para o funcionamento do codigo
	(*p).idade += 1
}

func main() {
	pessoa1 := pessoa{nome: "Marcelo", idade: 30}
	fmt.Println(pessoa1)
	mudeMe(&pessoa1) // vai no endereço original e troca um valor
	fmt.Println(pessoa1)
}
