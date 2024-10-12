package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// HandleError lida com erros e envia uma resposta HTTP apropriada
func HandleError(w http.ResponseWriter, err error, statusCode int) {
	log.Printf("Erro: %v", err)
	w.WriteHeader(statusCode)
	response := map[string]string{"error": err.Error()}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Erro ao codificar a resposta de erro: %v", err)
	}
}
