package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// o conteúdo que será passado para o arquivo deve ser um []byte
	conteudo := []byte("O arquivo foi criado!")

	// escrever em um arquivo (nomedoarquivo, o []byte que vai ser inserido, a permissão de usuario)
	err := os.WriteFile("teste.txt", conteudo, 0666)
	// tratamento de erros, se existir um erro, é mostrado pra mim
	if err != nil {
		log.Fatal(err)
	}

	// ler o conteúdo do arquivo, retorna um []byte
	arquivo, errR := os.ReadFile("teste.txt")
	if errR != nil {
		log.Fatal(errR)
	}

	fmt.Println(string(arquivo))
}
