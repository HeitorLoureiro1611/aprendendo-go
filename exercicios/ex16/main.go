// Demonstre o resto da divisão por 4 de todos os números entre 10 e 100

package main

import (
	"fmt"
)

func main() {

	for i := 10; i <= 100; i++ {
		fmt.Printf("%d / 4 = %.2f\n", i, float64(i%4))
	}

}
