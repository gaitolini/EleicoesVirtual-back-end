package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

// Auth é um middleware que verifica o token de autenticação na requisição
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Verificar se estamos em ambiente de desenvolvimento
		if os.Getenv("ENVIRONMENT") == "development" {
			log.Println("Ambiente de desenvolvimento detectado. Ignorando autenticação.")
			next.ServeHTTP(w, r)
			return
		}

		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Token de autenticação não fornecido", http.StatusUnauthorized)
			return
		}

		// Verificar o formato do token
		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Formato do token inválido", http.StatusUnauthorized)
			return
		}

		tokenString = parts[1]

		// Verificar se o token é válido usando JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verificar se o método de assinatura é o esperado
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				http.Error(w, "Método de assinatura inválido", http.StatusUnauthorized)
				return nil, nil
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		// Caso o token seja válido, prosseguir com a requisição
		next.ServeHTTP(w, r)
	}
}

// isValidToken verifica se o token é válido usando JWT (agora substituído pela lógica de validação no middleware)
func isValidToken(token string) bool {
	// Esta função agora é substituída pela validação JWT dentro do middleware
	log.Println("Função isValidToken obsoleta - use a validação JWT no middleware Auth")
	return false
}
