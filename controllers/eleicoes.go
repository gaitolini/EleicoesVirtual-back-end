package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gaitolini/EleicoesVirtual-back-end/models"
	"github.com/gaitolini/EleicoesVirtual-back-end/services"
	"github.com/gaitolini/EleicoesVirtual-back-end/utils"
	"github.com/gorilla/mux"
)

// CriarEleicao cria uma nova eleição
func CriarEleicao(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

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
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

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
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

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

// AtualizarEleicao lida com a atualização de uma eleição existente
func AtualizarEleicao(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	log.Println("Entrando no controlador de AtualizarEleicao")

	var eleicaoAtualizada models.Eleicao
	if err := json.NewDecoder(r.Body).Decode(&eleicaoAtualizada); err != nil {
		utils.HandleError(w, err, http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		utils.HandleError(w, nil, http.StatusBadRequest)
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
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, "ID não fornecido", http.StatusBadRequest)
		return
	}

	err := services.DeletarEleicao(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao deletar a eleição: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Eleição com ID %s deletada com sucesso.", id)))
}
