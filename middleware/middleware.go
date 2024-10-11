package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var authClient *auth.Client

func InitFirebaseAuth() {
	opt := option.WithCredentialsFile("path/to/your/firebase-service-account.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Erro ao inicializar o app Firebase: %v", err)
	}

	authClient, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Erro ao inicializar o Firebase Auth: %v", err)
	}
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Token não fornecido", http.StatusUnauthorized)
			return
		}

		// Extraindo o token do header "Authorization"
		idToken := strings.TrimSpace(strings.Replace(authHeader, "Bearer ", "", 1))

		// Verificar token
		token, err := authClient.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		// Usuário autenticado com sucesso
		log.Printf("Usuário autenticado: %v", token.UID)
		next.ServeHTTP(w, r)
	})
}
