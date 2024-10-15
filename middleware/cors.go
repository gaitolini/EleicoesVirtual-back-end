package middleware

import (
	"log"
	"net/http"
)

// Middleware para adicionar cabeçalhos CORS manualmente
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("CORS Middleware ativado para %s %s", r.Method, r.URL.Path)

		// Definir os cabeçalhos de CORS
		w.Header().Set("Access-Control-Allow-Origin", "https://eleicoesvirtual.web.app")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Se a requisição for OPTIONS, respondemos imediatamente
		if r.Method == http.MethodOptions {
			log.Printf("Requisição preflight OPTIONS recebida e tratada")
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
