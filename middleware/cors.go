package middleware

import (
	"log"
	"net/http"
)

// Middleware para adicionar cabeçalhos CORS manualmente
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("CORS Middleware ativado para %s %s", r.Method, r.URL.Path)

		origin := r.Header.Get("Origin")
		log.Printf("Origin: %s", origin)
		// Definir os cabeçalhos CORS dependendo da origem da requisição
		if origin == "http://localhost:3000" || origin == "https://eleicoesvirtual.web.app" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "https://eleicoesvirtual.web.app") // Origem padrão (pode ser alterado)
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Responder imediatamente se for uma requisição preflight OPTIONS
		if r.Method == http.MethodOptions {
			log.Printf("Recebendo requisição OPTIONS para %s", r.URL.Path)
			w.WriteHeader(http.StatusOK)
			return
		}

		// Passar para o próximo middleware se não for OPTIONS
		next.ServeHTTP(w, r)
	})
}
