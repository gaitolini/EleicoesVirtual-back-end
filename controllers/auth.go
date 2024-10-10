package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gaitolini/EleicoesVirtual-back-end/services"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Rota para inscrição de novos usuários
func Signup(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Erro ao processar a requisição", http.StatusBadRequest)
		return
	}

	uid, err := services.SignupWithEmail(req.Email, req.Password)
	if err != nil {
		http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuário criado com sucesso", "uid": uid})
}

// Rota para login de usuários
func Login(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Erro ao processar a requisição", http.StatusBadRequest)
		return
	}

	uid, err := services.LoginWithEmail(req.Email, req.Password)
	if err != nil {
		http.Error(w, "Erro ao autenticar", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Login realizado com sucesso", "uid": uid})
}
