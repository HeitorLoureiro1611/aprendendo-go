package main

import (
	"fmt"
)

// recebe um int
func add(i int) {
	// modifica esse int q chegou na função
	i++
	fmt.Println("Na função está:", i)
} // aqui esse "i int" não existe mais

// Função que recebe um ponteiro de um int
func addPonteiro(i *int) {
	// aumenta 1 no valor da variável
	*i++
	fmt.Println("Na função está:", *i)
} // fazendo o mesmo que x + 1 na main

// ponteiro é o "endendereço de memória da variável em questão"
func main() {

	// x = 10 -> 0xc098490 (endereço da memoria que eu armazenei algo)
	// variavel -> está aqui
	x := 10  // variável em questão
	px := &x // px é onde eu vou guardar o endereço de x

	fmt.Println(&x)  // printando o endereço de x
	fmt.Println(px)  // printando a variável px (endereço de x armazenado em uma variável)
	fmt.Println(*px) // "dereference" printando o valor que está dentro do endereço de px

	fmt.Println("==================================")

	add(x) // aqui eu apenas atribuo a soma dentro dessa função, fora dela a variável não é alterada
	fmt.Println("Fora da função:", x)

	addPonteiro(&x) // aqui eu estou fazendo uma alteração diretamente da memória, mudando o valor dela
	fmt.Println("Fora da função:", x)
}

// um dos usos comums de ponteiros é quando existe um costume de uso de uma variável especifica
// caso não queira fazer multiplas variáveis copiando aquele valor, é possível fazer uma cópia com ponteiros
