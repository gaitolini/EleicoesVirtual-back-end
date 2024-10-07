// controllers/eleicoes_controller.go
package controllers

import (
	"EleicoesVirtual-back-end/models"
	"EleicoesVirtual-back-end/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CriarEleicao(w http.ResponseWriter, r *http.Request) {
	var novaEleicao models.Eleicao
	err := json.NewDecoder(r.Body).Decode(&novaEleicao)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	eleicao := services.CriarEleicao(novaEleicao)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(eleicao)
}

func ListarEleicoes(w http.ResponseWriter, r *http.Request) {
	lista := services.ListarEleicoes()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lista)
}

func ObterEleicao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	eleicao, existe := services.ObterEleicao(id)
	if !existe {
		http.Error(w, "Eleição não encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(eleicao)
}

func AtualizarEleicao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var eleicaoAtualizada models.Eleicao
	err := json.NewDecoder(r.Body).Decode(&eleicaoAtualizada)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	eleicao, existe := services.AtualizarEleicao(id, eleicaoAtualizada)
	if !existe {
		http.Error(w, "Eleição não encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(eleicao)
}

func DeletarEleicao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if !services.DeletarEleicao(id) {
		http.Error(w, "Eleição não encontrada", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
