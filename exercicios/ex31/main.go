// Crie e use um struct anônimo.
// Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.

package main

import (
	"fmt"
)

func main() {
	// inicializador da minha struct
	computador := struct {
		quantPecas  map[int]string
		perifericos []string
	}{

		// aqui eu preencho a minha struct
		// chave : valor
		quantPecas: map[int]string{
			1: "Placa-mãe",
			2: "ram",
			0: "placa de video",
		},
		// []slice {valores...}
		perifericos: []string{"teclado", "mouse"},
	}

	fmt.Println("periféricos de computador:")
	for _, v := range computador.perifericos {
		fmt.Printf("%s\n", v)
	}
	fmt.Println("=================================================================")
	for k, v := range computador.quantPecas {
		fmt.Printf("temos: %d unidade(s) disponível de %s\n", k, v)
	}

}
