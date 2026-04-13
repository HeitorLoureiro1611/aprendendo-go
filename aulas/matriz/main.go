package main

import (
	"fmt"
)

func main() {
	// criação da matriz | slices of slices
	matriz := [][]int{
		//    0  1  2
		[]int{0, 1, 2},   // 0
		[]int{3, 4, 5},   // 1
		[]int{6, 7, 8},   // 2
		[]int{9, 10, 11}, // 3
	}
	// formatando a matriz no terminal
	for i := range matriz {
		fmt.Println(matriz[i])
	}
	//			  [linha] [coluna]
	fmt.Println(matriz[3][2])
}
