package middleware

import (
	"log"
	"net/http"
)

// Lista de origens permitidas
var allowedOrigins = []string{
	"https://eleicoesvirtual.web.app",
	"https://outrodominio.com", // Adicione outras origens conforme necessário
}

// isAllowedOrigin verifica se a origem está na lista de permitidos
func isAllowedOrigin(origin string) bool {
	for _, o := range allowedOrigins {
		if o == origin {
			return true
		}
	}
	return false
}

// Middleware para adicionar cabeçalhos CORS manualmente
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		// Se o Origin estiver vazio (possível em algumas requisições de ferramentas como Postman)
		if origin == "" {
			origin = "https://eleicoesvirtual.web.app" // Origem padrão ou permitir sem origin
		}

		if isAllowedOrigin(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		} else {
			log.Printf("CORS bloqueado para a origem: %s", origin)
		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
