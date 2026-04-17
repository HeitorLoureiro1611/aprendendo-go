package main

import (
	"fmt"
)

// Callback é "esperar uma coisa terminar pra realizar outra"
// recebe uma função com esses argumentos, e uma slice de numeros
func somaSomentePares(f func(x []int) int, y []int) int {
	var slice []int
	for _, v := range y {
		if v%2 == 0 {
			// adiciona todos os numeros pares da slice numa outra slice
			slice = append(slice, v)
		}
	}
	// depois de montar a slice com apenas pares, chama a tal função recebendo o slice montado dentro da função
	// por isso que na main não coloco os parâmetros que ela vai receber
	total := f(slice)
	return total
}

// Função de somar numeros de uma slice
func soma(z []int) int {
	somatorio := 0
	for _, v := range z {
		somatorio += v
	}
	// retorna a soma disso
	return somatorio
}

func main() {
	// t vai ser o resultado que a função somaSomentePares entregar
	// parametros: a função soma, slice
	// preciso passar a soma como argumento pois vou utilizar ela dentro da minha função
	//   func pai         func filha
	t := somaSomentePares(soma, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60})
	fmt.Println(t)
}
