package main

import (
	"log"
	"net/http"

	"github.com/gaitolini/EleicoesVirtual-back-end/controllers"
	"github.com/gaitolini/EleicoesVirtual-back-end/services"
	"github.com/gorilla/mux"
)

func main() {
	// Inicializar o Firestore
	services.InitializeFirestoreClient()

	// Configurar roteador
	r := mux.NewRouter()

	// Rotas para eleições
	r.HandleFunc("/eleicoes", controllers.CriarEleicao).Methods("POST")
	r.HandleFunc("/eleicoes", controllers.ListarEleicoes).Methods("GET")
	r.HandleFunc("/eleicoes/{id}", controllers.ObterEleicao).Methods("GET")
	r.HandleFunc("/eleicoes/{id}", controllers.AtualizarEleicao).Methods("PUT")
	r.HandleFunc("/eleicoes/{id}", controllers.DeletarEleicao).Methods("DELETE")

	// Iniciar o servidor
	log.Println("Servidor rodando na porta 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
