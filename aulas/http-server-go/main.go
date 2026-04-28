package main

import (
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home func..."))
}

func TesteView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Teste View Page..."))

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
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/view", TesteView)
	mux.HandleFunc("/create", TesteCreate)

	log.Print("Inicializando o servidor: http://localhost:8080/")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
