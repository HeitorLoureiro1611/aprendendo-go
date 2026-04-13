package main

import (
	"fmt"
)

func main() {

	// declaração de uma slice usando make
	// make(tipo, tamanho, capacidade)
	// caso saia da capacidade inicial
	// é criado uma nova lista dinâmica com o dobro da capacidade
	// e então copiado os valores da lista antiga para a nova e deletada a antiga
	slice := make([]int, 5, 10)

	// inserindo numeros no tamanho do slice
	for i := range slice {
		slice[i] = i
	}

	fmt.Println("Primeira Lista:", slice, "|", len(slice), "|", cap(slice))

	// colocando numeros na capacidade total do slice
	for x := len(slice); x < cap(slice); x++ {
		// para colocar novos items além do tamanho, é preciso usar append
		slice = append(slice, x)
	}

	fmt.Println("Primeira Lista: ", slice, "|", len(slice), "|", cap(slice))

	// Aqui acontece a criação da novaa lista com a capacidade dobrada e a copia
	// dos valores antigos, saiu da capacidade total da slice
	slice = append(slice, 10)

	fmt.Println("Segunda Lista: ", slice, "|", len(slice), "|", cap(slice))
}
