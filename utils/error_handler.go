package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// HandleError lida com erros e envia uma resposta HTTP apropriada
func HandleError(w http.ResponseWriter, err error, statusCode int) {
	if err != nil {
		log.Printf("Erro: %v", err) // Registrar o erro para debugging
		w.WriteHeader(statusCode)

		// Criar uma resposta JSON com uma mensagem de erro
		response := map[string]string{"error": err.Error()}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("Erro ao codificar a resposta de erro: %v", err)
		}
	} else {
		// Caso não tenha uma mensagem de erro específica
		w.WriteHeader(statusCode)
		response := map[string]string{"error": "Ocorreu um erro, mas nenhuma mensagem foi fornecida."}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("Erro ao codificar a resposta de erro: %v", err)
		}
	}
}
