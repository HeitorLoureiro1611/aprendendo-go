package main

import (
	"fmt"
)

// maps == tables {chave -> valor}
// "heitor" -> 22; "luis" -> 20...
func main() {

	// inicialização vazia
	//           chave : valor
	amigos := map[string]int{}

	// adicionando chave e valor
	amigos["Heitor"] = 1
	amigos["Gabriel"] = 2
	amigos["Giovanni"] = 3
	amigos["josé"] = 4

	// printando toda a lista
	fmt.Println(amigos)

	// printando o valor que tem na chave especifica
	fmt.Println(amigos["Heitor"]) // saida: 1
	fmt.Println("=======================")

	// comma ok idiom:
	// forma de distinção de valor vazio pra valor inicializado
	// _ desconsidera o uso do valor
	// ok: valor booleano, tem ou não algum valor dentro dessa chave?
	_, ok := amigos["estranho"]
	fmt.Println(ok) // saida: false
	fmt.Println("=======================")

	// checa se tenho um item especifico no map
	// caso tiver, retorna o valor
	// caso não tiver, retorna uma menságem
	// valor, ok    checa se o boolean é true (existe no map)
	if item, ok := amigos["josé"]; ok {
		fmt.Println(item)
	} else {
		fmt.Println("Valor inexistente!")
	}

	// Chave "Heitor" deletada do map
	delete(amigos, "Heitor")
	fmt.Println("=======================")

	// printando com range
	// key == chave de reconhecimento do map
	// value == valor dentro daquela chave
	for key, value := range amigos {
		fmt.Println(key, ":", value)
	}

}
