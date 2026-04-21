package main

import (
	// Importação da biblioteca json
	"encoding/json"
	"fmt"
)

// Go diferencia Publico de Privado com a primeira letra em maiúsculo
// Pessoa > publico
// pessoa > privado
type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int64
}

func main() {
	Cliente := Pessoa{"Delo", "Phantae", 26}

	var cliente2 Pessoa

	// b -> []bytes, err -> erro
	// marshal é a ordenação em json
	b, err1 := json.Marshal(Cliente)
	// caso tenha algum erro, mostre o erro
	if err1 != nil {
		fmt.Println(err1)
	}
	// ver o objeto em uma formatação json
	// é necessario parsear como string pois é um []bytes
	fmt.Println(string(b))

	// ##################################################################
	// Unmarshal e decodificando JSON

	// json é tratado como []bytes
	un := []byte(`{"Nome":"Delo","Sobrenome":"Phantae","Idade":26}`)

	// declaração do unmarshal
	// erro vai ser gerado
	// Unmarshal([]bytes, &var)
	err := json.Unmarshal(un, &cliente2)
	if err != nil {
		fmt.Println(err)
	}
	// não é lidado como uma variável decodificada de json
	fmt.Println(cliente2)

}
