// Crie um programa que utilize a declaração switch,
// onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".

package main

import (
	"fmt"
)

func main() {

	esporteFavorito := "Basquete"

	switch esporteFavorito {
	case "Futebol":
		fmt.Println("Você gosta de jogar bola")
	case "Basquete":
		fmt.Println("Você quer se tornar o LeBron James")
	case "Natação":
		fmt.Println("Nada igual o Michael Phelps")
	}
}
