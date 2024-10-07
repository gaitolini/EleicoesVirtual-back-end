package services

import (
	"EleicoesVirtual-back-end/models"
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
)

func CriarEleicao(novaEleicao models.Eleicao) (*firestore.DocumentRef, error) {
	ctx := context.Background()

	_, err := client.Collection("eleicoes").Doc(novaEleicao.ID).Set(ctx, novaEleicao)
	if err != nil {
		log.Printf("Erro ao criar a eleição: %v", err)
		return nil, err
	}

	fmt.Printf("Eleição %s criada com sucesso!\n", novaEleicao.Nome)
	return client.Collection("eleicoes").Doc(novaEleicao.ID), nil
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

func DeletarEleicao(id string) error {
	ctx := context.Background()

	_, err := client.Collection("eleicoes").Doc(id).Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}
