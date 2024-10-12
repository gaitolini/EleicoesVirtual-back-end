package services

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var (
	FirebaseApp *firebase.App
	Client      *firestore.Client
)

// InitializeFirestoreClient inicializa o cliente Firestore com credenciais fornecidas
func InitializeFirestoreClient(credentialsJSON string) {
	ctx := context.Background()

	// Remover a quebra de linha adicional das chaves privadas se necessário
	opt := option.WithCredentialsJSON([]byte(credentialsJSON))

	// Inicializar a aplicação Firebase com as credenciais
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Erro ao inicializar o Firebase App: %v", err)
	}

	// Armazenar a referência da aplicação
	FirebaseApp = app

	// Inicializar o cliente Firestore
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Erro ao inicializar o Firestore: %v", err)
	}

	Client = client
	log.Println("Firestore inicializado com sucesso.")
}
