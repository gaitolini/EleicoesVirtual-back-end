package services

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var (
	FirebaseApp  *firebase.App     // Armazenar a instância do Firebase App
	FirebaseAuth *auth.Client      // Armazenar a instância do Firebase Auth
	Client       *firestore.Client // Armazenar a instância do Firestore
)

// InitializeFirebaseClient inicializa o Firebase App e Firestore com credenciais fornecidas
func InitializeFirebaseClient(credentialsJSON string) {
	ctx := context.Background()

	// Configuração das credenciais para Firebase
	opt := option.WithCredentialsJSON([]byte(credentialsJSON))

	// Inicializar o Firebase App
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Erro ao inicializar o Firebase App: %v", err)
	}
	FirebaseApp = app // Armazenar a instância do Firebase App

	// Inicializar o Firestore
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Erro ao inicializar Firestore: %v", err)
	}
	Client = client

	log.Println("Firestore inicializado com sucesso.")
}

// InitializeFirebaseAuth inicializa o Firebase Auth para verificação de tokens JWT
func InitializeFirebaseAuth(ctx context.Context) error {
	// Inicializar Firebase Auth a partir da instância do Firebase App
	authClient, err := FirebaseApp.Auth(ctx)
	if err != nil {
		log.Fatalf("Erro ao inicializar FirebaseAuth: %v", err)
		return err
	}
	FirebaseAuth = authClient
	log.Println("Firebase Auth inicializado com sucesso.")
	return nil
}

// VerifyIDToken verifica o token JWT com Firebase Auth
func VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	token, err := FirebaseAuth.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, err
	}
	return token, nil
}
