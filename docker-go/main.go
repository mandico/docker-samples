package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

type Status struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// Definindo o cabeçalho da resposta como application/json
		w.Header().Set("Content-Type", "application/json")

		// Criando e preenchendo a estrutura de status
		status := Status{
			Status:    "ok",
			Timestamp: time.Now(),
		}

		// Codificando e enviando a resposta
		json.NewEncoder(w).Encode(status)

		// Log da requisição usando slog
		slog.Info("Received request", "method", r.Method, "path", r.URL.Path, "status", status.Status, "timestamp", status.Timestamp)
	})

	// Iniciando o servidor na porta 8080
	slog.Info("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
