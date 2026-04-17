package main

import (
	"fmt"
)

// closure é uma função que lembra e mantém acesso
// a variáveis do escopo onde foi criada.
// cada interação dessa função se da a qual variável ela está atrelada

// função que recebe uma função e retorna um inteiro
func contador() func() int {
	// parte do codigo que compartilha para todo o funcionamento da função
	// persistencia independente da variável
	x := 0

	return func() int {
		// a partir daqui, varia o resultado pra cada variável
		x++
		return x
	}
}

func main() {

	// essa variável se tornou uma função
	contadora := contador()

	fmt.Println(contadora())
	fmt.Println(contadora())
	fmt.Println(contadora())
	fmt.Println(contadora())

	fmt.Println("Fim do primeiro contador")
	fmt.Println("==========================")
	fmt.Println("inicio segundo contador")
	// se tornou uma função e vai funcionar de forma independente do contadora
	contadorb := contador()

	fmt.Println(contadorb())
	fmt.Println(contadorb())

}
