package services

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var authClient *auth.Client

// Inicializa o cliente Firebase Auth
func InitializeAuthClient() {
	ctx := context.Background()

	// Obter o caminho do arquivo JSON de credenciais do Firebase a partir de uma variável de ambiente
	credentialsFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if credentialsFile == "" {
		log.Fatalf("Variável de ambiente GOOGLE_APPLICATION_CREDENTIALS não definida")
	}

	opt := option.WithCredentialsFile(credentialsFile)

	// Inicializa o app do Firebase
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Erro ao inicializar o Firebase App: %v", err)
	}

	// Inicializa o cliente Auth
	authClient, err = app.Auth(ctx)
	if err != nil {
		log.Fatalf("Erro ao criar cliente Auth: %v", err)
	}

	log.Println("Firebase Auth inicializado com sucesso!")
}

// Função de cadastro de usuário (signup) com email e senha
func SignupWithEmail(email, password string) (string, error) {
	ctx := context.Background()

	// Parâmetros para criar o usuário
	params := (&auth.UserToCreate{}).Email(email).Password(password)
	userRecord, err := authClient.CreateUser(ctx, params)
	if err != nil {
		return "", err
	}

	log.Printf("Usuário criado com sucesso: %v", userRecord.UID)
	return userRecord.UID, nil
}

// Função de login com email e senha
func LoginWithEmail(email, password string) (string, error) {
	ctx := context.Background()

	// Autenticar o usuário pelo email
	userRecord, err := authClient.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	// Aqui normalmente você verificaria a senha (via frontend ou com Firebase Authentication)
	return userRecord.UID, nil
}
