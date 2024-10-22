package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gaitolini/EleicoesVirtual-back-end/services"
)

// AuthMiddleware verifica se o token JWT é válido
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extrair o token JWT do cabeçalho Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// O token geralmente vem no formato "Bearer {token}"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}
		idToken := parts[1]

		// Verificar o token usando Firebase Auth
		token, err := services.VerifyIDToken(r.Context(), idToken)
		if err != nil {
			log.Printf("Erro ao verificar o token: %v", err)
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// O token é válido; permitir o acesso e passar para o próximo middleware
		log.Printf("Token válido para o usuário: %v", token.UID)
		next.ServeHTTP(w, r)
	})
}
