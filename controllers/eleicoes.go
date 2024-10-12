package controllers

import (
	"encoding/json"
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

// DeletarEleicao deleta uma eleição específica
func DeletarEleicao(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := services.DeletarEleicao(id); err != nil {
		utils.HandleError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
