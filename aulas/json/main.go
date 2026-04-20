package main

import (
	// Importação da biblioteca json
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int64
}

func main() {
	Cliente := Pessoa{"Delo", "Phantae", 26}

	// b -> []bytes, err -> erro
	// marshal é a ordenação em json
	b, err := json.Marshal(Cliente)
	// caso tenha algum erro, mostre o erro
	if err != nil {
		fmt.Println(err)
	}
	// ver o objeto em uma formatação json
	// é necessario parsear como string pois é um []bytes
	fmt.Println(string(b))
}
