package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gaitolini/EleicoesVirtual-back-end/controllers"
	"github.com/gaitolini/EleicoesVirtual-back-end/middleware"
	"github.com/gaitolini/EleicoesVirtual-back-end/services"
	"github.com/gorilla/mux"
)

func main() {
	// Ler o arquivo de configuração JSON do Firebase
	file, err := ioutil.ReadFile("firebaseConfig.json")
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo firebaseConfig.json: %v", err)
	}

	// Inicializar o cliente Firestore
	services.InitializeFirestoreClient(string(file))

	// Definir o ambiente como "development" para testes locais
	os.Setenv("ENVIRONMENT", "development") // apenas para testes locais, remova em produção

	// Configurar as rotas usando mux
	r := mux.NewRouter()
	r.HandleFunc("/eleicoes/criar", middleware.Auth(controllers.CriarEleicao)).Methods("POST")
	r.HandleFunc("/eleicoes/listar", middleware.Auth(controllers.ListarEleicoes)).Methods("GET")
	r.HandleFunc("/eleicoes/atualizar", middleware.Auth(controllers.AtualizarEleicao)).Methods("PUT")
	r.HandleFunc("/eleicoes/deletar/{id}", middleware.Auth(controllers.DeletarEleicao)).Methods("DELETE")

	// Iniciar o servidor HTTP
	port := ":8081"
	log.Printf("Iniciando servidor na porta %s...", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
