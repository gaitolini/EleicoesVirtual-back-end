package middleware

import (
	"log"
	"net/http"
)

// Middleware para adicionar cabeçalhos CORS manualmente
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("CORS Middleware ativado para %s %s", r.Method, r.URL.Path)

		// Permitir qualquer origem ou uma lista de origens permitidas
		w.Header().Set("Access-Control-Allow-Origin", "https://eleicoesvirtual.web.app")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Verifica se é uma requisição preflight (OPTIONS)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Passar para o próximo middleware
		next.ServeHTTP(w, r)
	})
}
