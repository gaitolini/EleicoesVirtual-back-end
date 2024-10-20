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
		if origin == "" {
			log.Printf("Origin não encontrado. Usando valor padrão.")
			origin = "https://eleicoesvirtual.web.app" // Origem padrão
		}

		log.Printf("Origin: %s", origin)

		// Configurar CORS
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Responder imediatamente se for uma requisição preflight (OPTIONS)
		if r.Method == http.MethodOptions {
			log.Printf("Recebendo requisição OPTIONS para %s", r.URL.Path)
			w.WriteHeader(http.StatusOK)
			return
		}

		// Passar para o próximo middleware
		next.ServeHTTP(w, r)
	})
}
