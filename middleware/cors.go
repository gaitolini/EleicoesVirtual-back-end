package middleware

import (
	"log"
	"net/http"
)

// Lista de origens permitidas
var allowedOrigins = []string{
	"https://eleicoesvirtual.web.app", // Adicione outras origens permitidas conforme necessário
	"http://localhost:3000",
	"http://localhost:8082",
}

// Verifica se a origem está na lista permitida
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
		log.Printf("Origin recebido: %s", origin)

		// Se a origem estiver permitida, adicionar os cabeçalhos CORS
		if origin != "" && isAllowedOrigin(origin) {
			log.Printf("CORS permitido para a origem: %s", origin)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			// w.Header().Set("Access-Control-Allow-Credentials", "false")
		} else {
			log.Printf("CORS bloqueado para a origem: %s", origin)
		}

		// Se for uma requisição preflight (OPTIONS), responder imediatamente
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Passar para o próximo middleware
		next.ServeHTTP(w, r)
	})
}
