//Crie constantes tipadas e não-tipadas.
//Demonstre seus valores.

package main

import (
	"fmt"
)

// não tipada
const x = 10

// tipada
const y int = 9

func main() {
	fmt.Printf("x = %v\ny = %v\n", x, y)
}
