package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gaitolini/EleicoesVirtual-back-end/models"
	"github.com/google/uuid"
	"google.golang.org/api/option"
)

// Variável global para o Firestore
var client *firestore.Client

// Inicializa o Firestore com a chave privada do Firebase
func InitializeFirestoreClient() {
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

	// Cria o cliente Firestore
	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Erro ao criar cliente Firestore: %v", err)
	}

	client = firestoreClient
	fmt.Println("Firestore inicializado com sucesso!")
}

func CriarEleicao(novaEleicao models.Eleicao) (*firestore.DocumentRef, string, error) {
	ctx := context.Background()

	// Gerar um ID aleatório, se estiver vazio
	if novaEleicao.ID == "" {
		novaEleicao.ID = uuid.New().String() // Gera um UUID
	}

	// Verificar se as datas estão vazias e preencher com uma data padrão, se necessário
	if novaEleicao.DataInicio.IsZero() {
		novaEleicao.DataInicio = time.Now()
	}
	if novaEleicao.DataFim.IsZero() {
		novaEleicao.DataFim = time.Now().AddDate(0, 3, 0)
	}

	// Criar o documento no Firestore
	_, err := client.Collection("eleicoes").Doc(novaEleicao.ID).Set(ctx, novaEleicao)
	if err != nil {
		log.Printf("Erro ao criar a eleição: %v", err)
		return nil, "", err
	}

	log.Printf("Eleição %s criada com sucesso!", novaEleicao.Nome)
	return client.Collection("eleicoes").Doc(novaEleicao.ID), novaEleicao.ID, nil
}

func ListarEleicoes() ([]models.Eleicao, error) {
	ctx := context.Background()
	iter := client.Collection("eleicoes").Documents(ctx)

	var eleicoes []models.Eleicao
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}
		var eleicao models.Eleicao
		doc.DataTo(&eleicao)
		eleicoes = append(eleicoes, eleicao)
	}

	return eleicoes, nil
}

func ObterEleicao(id string) (models.Eleicao, error) {
	ctx := context.Background()
	doc, err := client.Collection("eleicoes").Doc(id).Get(ctx)
	if err != nil {
		return models.Eleicao{}, err
	}

	var eleicao models.Eleicao
	doc.DataTo(&eleicao)

	return eleicao, nil
}

func AtualizarEleicao(id string, eleicaoAtualizada models.Eleicao) error {
	ctx := context.Background()

	_, err := client.Collection("eleicoes").Doc(id).Set(ctx, eleicaoAtualizada)
	if err != nil {
		return err
	}

	return nil
}

// Deletar uma eleição
func DeletarEleicao(id string) error {
	ctx := context.Background()

	log.Printf("Tentando deletar a eleição com ID: %s", id)

	// Realizar a deleção no Firestore
	_, err := client.Collection("eleicoes").Doc(id).Delete(ctx)
	if err != nil {
		log.Printf("Erro ao deletar a eleição com ID %s: %v", id, err)
		return err
	}

	log.Printf("Eleição com ID %s deletada com sucesso", id)
	return nil
}
