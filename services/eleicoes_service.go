// services/eleicoes_service.go
package services

import (
	"EleicoesVirtual-back-end/models"
	"sync"
)

var (
	eleicoes = make(map[string]models.Eleicao)
	mu       sync.Mutex
)

func CriarEleicao(novaEleicao models.Eleicao) models.Eleicao {
	mu.Lock()
	eleicoes[novaEleicao.ID] = novaEleicao
	mu.Unlock()

	return novaEleicao
}

func ListarEleicoes() []models.Eleicao {
	mu.Lock()
	defer mu.Unlock()

	var lista []models.Eleicao
	for _, eleicao := range eleicoes {
		lista = append(lista, eleicao)
	}
	return lista
}

func ObterEleicao(id string) (models.Eleicao, bool) {
	mu.Lock()
	defer mu.Unlock()

	eleicao, existe := eleicoes[id]
	return eleicao, existe
}

func AtualizarEleicao(id string, eleicaoAtualizada models.Eleicao) (models.Eleicao, bool) {
	mu.Lock()
	defer mu.Unlock()

	_, existe := eleicoes[id]
	if !existe {
		return models.Eleicao{}, false
	}

	eleicaoAtualizada.ID = id
	eleicoes[id] = eleicaoAtualizada
	return eleicaoAtualizada, true
}

func DeletarEleicao(id string) bool {
	mu.Lock()
	defer mu.Unlock()

	_, existe := eleicoes[id]
	if existe {
		delete(eleicoes, id)
		return true
	}
	return false
}
