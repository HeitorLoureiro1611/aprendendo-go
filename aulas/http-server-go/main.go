package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home func..."))
}

func TesteView(w http.ResponseWriter, r *http.Request) {

	// strconv.Atoi pega o valor do id que está na url do servidor
	// e transforma em um int, para fazer a condição if else
	// id é o resultado da conversão da string para int
	id, err := strconv.Atoi(r.PathValue("id"))
	// se o ID dado na url for menor que 1, pagina não vai carregar
	if err != nil || id < 1 {
		log.Print("Id Não é um numero valido")
		// metodo para retornar um 404 na pagina
		http.NotFound(w, r)
	} else {

		// formata o valor do id da sessão
		// Sprintf pq eu quero jogar a string em uma variavel
		msg := fmt.Sprintf("numero do id da sessão: %d...", id)

		w.Write([]byte("View Page\n"))

		w.Write([]byte(msg))
	}
}

func TesteCreate(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Form..."))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// objeto mux para mapear as paginas
	mux := http.NewServeMux()

	// paginas sendo inicializadas pelo mux no meu servidor
	// {$} implica que A pagina DEVE ser aquela, sem o {$} o mux entende que toda pagina é valida com a função home
	mux.HandleFunc("GET /{$}", Home)
	// TesteView se torna uma closure
	// Coringa {id}, o valor que será colocado na url será jogado para a função
	// e retorna a pagina referente ao id informado
	mux.HandleFunc("/view/{id}", TesteView)
	mux.HandleFunc("/create", TesteCreate)

	log.Print("Inicializando o servidor: http://localhost:8080/")
	err := http.ListenAndServe(":8080", mux)

	log.Fatal(err)

}
