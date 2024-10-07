// main.go
package main

import (
	"EleicoesVirtual-back-end/controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/eleicoes", controllers.CriarEleicao).Methods("POST")
	r.HandleFunc("/eleicoes", controllers.ListarEleicoes).Methods("GET")
	r.HandleFunc("/eleicoes/{id}", controllers.ObterEleicao).Methods("GET")
	r.HandleFunc("/eleicoes/{id}", controllers.AtualizarEleicao).Methods("PUT")
	r.HandleFunc("/eleicoes/{id}", controllers.DeletarEleicao).Methods("DELETE")

	fmt.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
