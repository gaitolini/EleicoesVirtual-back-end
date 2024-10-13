package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gaitolini/EleicoesVirtual-back-end/models"
	"github.com/gaitolini/EleicoesVirtual-back-end/services"
	"github.com/gaitolini/EleicoesVirtual-back-end/utils"
	"github.com/gorilla/mux"
)

// CriarEleicao cria uma nova eleição
func CriarEleicao(w http.ResponseWriter, r *http.Request) {
	var novaEleicao models.Eleicao
	if err := json.NewDecoder(r.Body).Decode(&novaEleicao); err != nil {
		utils.HandleError(w, err, http.StatusBadRequest)
		return
	}

	id, err := services.CriarEleicao(novaEleicao)
	if err != nil {
		utils.HandleError(w, err, http.StatusInternalServerError)
		return
	}

	novaEleicao.ID = id
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(novaEleicao); err != nil {
		utils.HandleError(w, err, http.StatusInternalServerError)
	}
}

// ListarEleicoes lista todas as eleições
func ListarEleicoes(w http.ResponseWriter, r *http.Request) {
	eleicoes, err := services.ListarEleicoes()
	if err != nil {
		utils.HandleError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(eleicoes); err != nil {
		utils.HandleError(w, err, http.StatusInternalServerError)
	}
}

// ObterEleicao obtém uma eleição específica
func ObterEleicao(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	eleicao, err := services.ObterEleicao(id)
	if err != nil {
		utils.HandleError(w, err, http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(eleicao); err != nil {
		utils.HandleError(w, err, http.StatusInternalServerError)
	}
}

// AtualizarEleicao atualiza uma eleição existente
func AtualizarEleicao(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var eleicaoAtualizada models.Eleicao
	if err := json.NewDecoder(r.Body).Decode(&eleicaoAtualizada); err != nil {
		utils.HandleError(w, err, http.StatusBadRequest)
		return
	}

	if err := services.AtualizarEleicao(id, eleicaoAtualizada); err != nil {
		utils.HandleError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeletarEleicao lida com a exclusão de uma eleição existente
func DeletarEleicao(w http.ResponseWriter, r *http.Request) {
	// Obter o ID da eleição a ser deletada
	id := r.URL.Path[len("/eleicoes/deletar/"):]
	if id == "" {
		utils.HandleError(w, fmt.Errorf("ID da eleição não fornecido"), http.StatusBadRequest)
		return
	}

	// Chamar o serviço para deletar a eleição
	err := services.DeletarEleicao(id)
	if err != nil {
		if err.Error() == fmt.Sprintf("Eleição com ID %s não encontrada", id) {
			utils.HandleError(w, err, http.StatusNotFound)
		} else {
			utils.HandleError(w, err, http.StatusInternalServerError)
		}
		return
	}

	// Retornar sucesso se a eleição for deletada
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Eleição deletada com sucesso"})
}
