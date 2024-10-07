package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Bem-vindo ao EleicoesVirtual!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Servidor rodando na porta %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}

//New file
