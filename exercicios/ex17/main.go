// Crie um programa que demonstre o funcionamento da declaração if.

package main

import (
	"fmt"
)

func main() {
	op := false
	if !op {
		fmt.Println("True!")
	} else {
		fmt.Println("False :(")
	}
}
