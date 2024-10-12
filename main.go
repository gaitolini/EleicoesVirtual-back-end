package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gaitolini/EleicoesVirtual-back-end/controllers"
	"github.com/gaitolini/EleicoesVirtual-back-end/middleware"
	"github.com/gaitolini/EleicoesVirtual-back-end/services"
)

func main() {
	// Ler o arquivo de configuração JSON do Firebase
	file, err := ioutil.ReadFile("firebaseConfig.json")
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo firebaseConfig.json: %v", err)
	}

	// Inicializar o cliente Firestore
	services.InitializeFirestoreClient(string(file))

	os.Setenv("ENVIRONMENT", "development") // apenas para testes locais, remova em produção

	// Configurar as rotas
	http.HandleFunc("/eleicoes", middleware.Auth(controllers.HandleEleicoes))
	http.HandleFunc("/eleicoes/criar", middleware.Auth(controllers.CriarEleicao))
	http.HandleFunc("/eleicoes/listar", middleware.Auth(controllers.ListarEleicoes))
	http.HandleFunc("/eleicoes/atualizar", middleware.Auth(controllers.AtualizarEleicao))
	http.HandleFunc("/eleicoes/deletar", middleware.Auth(controllers.DeletarEleicao))

	// Iniciar o servidor HTTP
	port := ":8081"
	log.Printf("Iniciando servidor na porta %s...", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
