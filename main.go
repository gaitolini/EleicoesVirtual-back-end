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

	// Configurar as rotas usando o mux
	r := mux.NewRouter()

	// Middleware para adicionar cabeçalhos CORS manualmente
	r.Use(middleware.CorsMiddleware)

	// Registrar rotas CRUD para eleições
	r.HandleFunc("/eleicoes", middleware.Auth(controllers.CriarEleicao)).Methods(http.MethodPost)
	r.HandleFunc("/eleicoes", middleware.Auth(controllers.ListarEleicoes)).Methods(http.MethodGet)
	r.HandleFunc("/eleicoes/{id}", middleware.Auth(controllers.AtualizarEleicao)).Methods(http.MethodPut)
	r.HandleFunc("/eleicoes/{id}", middleware.Auth(controllers.DeletarEleicao)).Methods(http.MethodDelete)
	r.HandleFunc("/eleicoes/{id}", middleware.Auth(controllers.ObterEleicao)).Methods(http.MethodGet)

	// Registrar log para todas as requisições
	r.Use(loggingMiddleware)

	// Iniciar o servidor HTTP
	port := ":8081"
	log.Printf("Iniciando servidor na porta %s...", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}

// Middleware para logar todas as requisições HTTP
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Recebendo solicitação: %s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
