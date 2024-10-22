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

		// Se a origem estiver na lista de permitidas, adicionar os cabeçalhos CORS
		if origin != "" && isAllowedOrigin(origin) {
			log.Printf("CORS permitido para a origem: %s", origin)
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		} else {
			log.Printf("CORS bloqueado para a origem: %s", origin)
			// Se a origem não for permitida, não adicionamos os cabeçalhos CORS
		}

		// Verificar se é uma requisição preflight (OPTIONS)
		if r.Method == http.MethodOptions {
			log.Printf("Recebendo requisição OPTIONS para %s", r.URL.Path)
			w.WriteHeader(http.StatusOK)
			return
		}

		// Passar para o próximo middleware
		next.ServeHTTP(w, r)
	})
}
