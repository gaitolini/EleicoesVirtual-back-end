package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gaitolini/EleicoesVirtual-back-end/models"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CriarEleicao cria uma nova eleição no Firestore
func CriarEleicao(eleicao models.Eleicao) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	docRef, _, err := Client.Collection("eleicoes").Add(ctx, eleicao)
	if err != nil {
		log.Printf("Erro ao criar eleição: %v", err)
		return "", err
	}

	return docRef.ID, nil
}

// ListarEleicoes lista todas as eleições do Firestore
func ListarEleicoes() ([]models.Eleicao, error) {
	ctx := context.Background()
	iter := Client.Collection("eleicoes").Documents(ctx)

	var eleicoes []models.Eleicao
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Erro ao listar eleições: %v", err)
			return nil, err
		}

		var eleicao models.Eleicao
		if err := doc.DataTo(&eleicao); err != nil {
			log.Printf("Erro ao decodificar dados da eleição: %v", err)
			return nil, err
		}
		eleicao.ID = doc.Ref.ID // Adicionar o ID do documento à estrutura
		eleicoes = append(eleicoes, eleicao)
	}

	return eleicoes, nil
}

// ObterEleicao obtém uma eleição específica pelo ID do Firestore
func ObterEleicao(id string) (models.Eleicao, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	doc, err := Client.Collection("eleicoes").Doc(id).Get(ctx)
	if err != nil {
		log.Printf("Erro ao obter eleição: %v", err)
		return models.Eleicao{}, err
	}

	var eleicao models.Eleicao
	if err := doc.DataTo(&eleicao); err != nil {
		log.Printf("Erro ao decodificar dados da eleição: %v", err)
		return models.Eleicao{}, err
	}
	eleicao.ID = doc.Ref.ID // Adicionar o ID do documento à estrutura

	return eleicao, nil
}

// AtualizarEleicao atualiza uma eleição existente no Firestore
func AtualizarEleicao(id string, eleicao models.Eleicao) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Referência ao documento
	docRef := Client.Collection("eleicoes").Doc(id)

	// Verificar se o documento existe
	doc, err := docRef.Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			log.Printf("Erro: Eleição com ID %s não encontrada.", id)
			return fmt.Errorf("Eleição com ID %s não encontrada", id)
		}
		log.Printf("Erro ao buscar eleição: %v", err)
		return err
	}

	log.Printf("Documento encontrado: %v", doc.Ref.ID)

	// Atualizar o documento
	_, err = docRef.Set(ctx, eleicao)
	if err != nil {
		log.Printf("Erro ao atualizar eleição: %v", err)
		return err
	}

	log.Printf("Eleição com ID %s atualizada com sucesso.", id)
	return nil
}

// DeletarEleicao remove uma eleição do Firestore
// DeletarEleicao remove uma eleição do Firestore
func DeletarEleicao(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Referência ao documento a ser deletado
	docRef := Client.Collection("eleicoes").Doc(id)

	// Verificar se o documento existe antes de tentar deletar
	_, err := docRef.Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			log.Printf("Erro: Eleição com ID %s não encontrada.", id)
			return fmt.Errorf("Eleição com ID %s não encontrada", id)
		}
		log.Printf("Erro ao buscar eleição: %v", err)
		return err
	}

	// Deletar o documento
	_, err = docRef.Delete(ctx)
	if err != nil {
		log.Printf("Erro ao deletar eleição: %v", err)
		return err
	}

	log.Printf("Eleição com ID %s deletada com sucesso.", id)
	return nil
}
