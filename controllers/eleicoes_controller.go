package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gaitolini/EleicoesVirtual-back-end/models"
	"github.com/gaitolini/EleicoesVirtual-back-end/services"
	"github.com/gorilla/mux"
)

// Cria uma eleição
func CriarEleicao(w http.ResponseWriter, r *http.Request) {
	var novaEleicao models.Eleicao
	err := json.NewDecoder(r.Body).Decode(&novaEleicao)
	if err != nil {
		http.Error(w, "Erro ao decodificar o corpo da requisição", http.StatusBadRequest)
		return
	}

	_, err = services.CriarEleicao(novaEleicao)
	if err != nil {
		http.Error(w, "Erro ao criar a eleição", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Eleição criada com sucesso"})
}

// Lista todas as eleições
func ListarEleicoes(w http.ResponseWriter, r *http.Request) {
	eleicoes, err := services.ListarEleicoes()
	if err != nil {
		http.Error(w, "Erro ao listar eleições", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(eleicoes)
}

// Obtém uma eleição por ID
func ObterEleicao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	eleicao, err := services.ObterEleicao(id)
	if err != nil {
		http.Error(w, "Erro ao obter a eleição", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(eleicao)
}

// Atualiza uma eleição por ID
func AtualizarEleicao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var eleicaoAtualizada models.Eleicao
	err := json.NewDecoder(r.Body).Decode(&eleicaoAtualizada)
	if err != nil {
		http.Error(w, "Erro ao decodificar o corpo da requisição", http.StatusBadRequest)
		return
	}

	err = services.AtualizarEleicao(id, eleicaoAtualizada)
	if err != nil {
		http.Error(w, "Erro ao atualizar a eleição", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Eleição atualizada com sucesso"})
}

// Deleta uma eleição por ID
func DeletarEleicao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	err := services.DeletarEleicao(id)
	if err != nil {
		http.Error(w, "Erro ao deletar a eleição", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Eleição deletada com sucesso"})
}
