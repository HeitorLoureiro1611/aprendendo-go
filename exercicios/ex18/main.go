// Crie um programa que utilize a declaração switch, sem switch statement especificado.

package main

import (
	"fmt"
)

func main() {
	x := 2
	switch {
	case x == 1:
		fmt.Println("Um")
	case x == 2:
		fmt.Println("Dois")
	case x == 3:
		fmt.Println("Três")
	}
}
