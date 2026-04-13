// Crie uma slice contendo slices de strings ([][]string).
// Atribua valores a este slice multi-dimensional da seguinte maneira:
// "Nome", "Sobrenome", "Hobby favorito"
// Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.

package main

import (
	"fmt"
)

func main() {
	matriz := [][]string{
		//  []string{"Nome", "sobrenome", "Hobby"},
		[]string{"Heitor", "Loureiro", "Jogar video-game"},
		[]string{"Thainá", "Silva", "Ler"},
		[]string{"Gopher", "Golang", "Programar"},
	}

	for i := range matriz {
		fmt.Println(matriz[i])
	}
}
