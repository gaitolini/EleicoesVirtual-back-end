package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gaitolini/EleicoesVirtual-back-end/models"
	"github.com/gaitolini/EleicoesVirtual-back-end/services"
	"github.com/gaitolini/EleicoesVirtual-back-end/utils"
)

// HandleEleicoes é responsável por lidar com todas as operações CRUD relacionadas a eleições
func HandleEleicoes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ListarEleicoes(w, r)
	case http.MethodPost:
		CriarEleicao(w, r)
	case http.MethodPut:
		AtualizarEleicao(w, r)
	case http.MethodDelete:
		DeletarEleicao(w, r)
	default:
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
	}
}

// CriarEleicao lida com a criação de uma nova eleição
func CriarEleicao(w http.ResponseWriter, r *http.Request) {
	var novaEleicao models.Eleicao
	if err := json.NewDecoder(r.Body).Decode(&novaEleicao); err != nil {
		utils.HandleError(w, err, http.StatusBadRequest)
		return
	}

	if err := services.CriarEleicao(novaEleicao); err != nil {
		utils.HandleError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(novaEleicao); err != nil {
		utils.HandleError(w, err, http.StatusInternalServerError)
	}
}

// ListarEleicoes lida com a listagem de eleições
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

// AtualizarEleicao lida com a atualização de uma eleição existente
func AtualizarEleicao(w http.ResponseWriter, r *http.Request) {
	var eleicaoAtualizada models.Eleicao
	if err := json.NewDecoder(r.Body).Decode(&eleicaoAtualizada); err != nil {
		utils.HandleError(w, err, http.StatusBadRequest)
		return
	}

	// Extrair o ID da eleição a ser atualizada
	id := r.URL.Query().Get("id")
	if strings.TrimSpace(id) == "" {
		http.Error(w, "ID da eleição é obrigatório", http.StatusBadRequest)
		return
	}

	if err := services.AtualizarEleicao(id, eleicaoAtualizada); err != nil {
		utils.HandleError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(eleicaoAtualizada); err != nil {
		utils.HandleError(w, err, http.StatusInternalServerError)
	}
}

// DeletarEleicao lida com a exclusão de uma eleição existente
func DeletarEleicao(w http.ResponseWriter, r *http.Request) {
	// Extrair o ID da eleição a ser deletada
	id := r.URL.Query().Get("id")
	if strings.TrimSpace(id) == "" {
		http.Error(w, "ID da eleição é obrigatório", http.StatusBadRequest)
		return
	}

	if err := services.DeletarEleicao(id); err != nil {
		utils.HandleError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Eleição deletada com sucesso"))
}
