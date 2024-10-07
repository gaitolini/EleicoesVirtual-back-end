package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type Eleicao struct {
	ID         string `json:"id"`
	Nome       string `json:"nome"`
	Descricao  string `json:"descricao"`
	DataInicio string `json:"data_inicio"`
	DataFim    string `json:"data_fim"`
}

var eleicoes = make(map[string]Eleicao)
var mu sync.Mutex

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/eleicoes", criarEleicao).Methods("POST")
	r.HandleFunc("/eleicoes", listarEleicoes).Methods("GET")
	r.HandleFunc("/eleicoes/{id}", obterEleicao).Methods("GET")
	r.HandleFunc("/eleicoes/{id}", atualizarEleicao).Methods("PUT")
	r.HandleFunc("/eleicoes/{id}", deletarEleicao).Methods("DELETE")

	fmt.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func criarEleicao(w http.ResponseWriter, r *http.Request) {
	var novaEleicao Eleicao
	err := json.NewDecoder(r.Body).Decode(&novaEleicao)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	eleicoes[novaEleicao.ID] = novaEleicao
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(novaEleicao)
}

func listarEleicoes(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var lista []Eleicao
	for _, eleicao := range eleicoes {
		lista = append(lista, eleicao)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lista)
}

func obterEleicao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	mu.Lock()
	eleicao, existe := eleicoes[id]
	mu.Unlock()

	if !existe {
		http.Error(w, "Eleição não encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(eleicao)
}

func atualizarEleicao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	mu.Lock()
	_, existe := eleicoes[id]
	mu.Unlock()

	if !existe {
		http.Error(w, "Eleição não encontrada", http.StatusNotFound)
		return
	}

	var eleicaoAtualizada Eleicao
	err := json.NewDecoder(r.Body).Decode(&eleicaoAtualizada)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	eleicaoAtualizada.ID = id

	mu.Lock()
	eleicoes[id] = eleicaoAtualizada
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(eleicaoAtualizada)
}

func deletarEleicao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	mu.Lock()
	_, existe := eleicoes[id]
	if existe {
		delete(eleicoes, id)
	}
	mu.Unlock()

	if !existe {
		http.Error(w, "Eleição não encontrada", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
