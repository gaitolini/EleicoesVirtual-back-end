package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gaitolini/EleicoesVirtual-back-end/controllers"
	"github.com/gaitolini/EleicoesVirtual-back-end/services"
	"github.com/gorilla/mux"
)

func main() {

	_, err := os.Stat(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	if err != nil {
		log.Fatalf("Erro ao acessar o arquivo de credenciais: %v", err)
	} else {
		log.Println("Arquivo de credenciais encontrado com sucesso!")
	}

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

//Testando o CI-CD
