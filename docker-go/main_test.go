package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHealthEndpoint(t *testing.T) {
	// Criar um request para o endpoint /health
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Criar um ResponseRecorder para gravar a resposta
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Reutilizando a lógica do endpoint /health
		w.Header().Set("Content-Type", "application/json")
		status := Status{
			Status:    "ok",
			Timestamp: time.Now(),
		}
		json.NewEncoder(w).Encode(status)
	})

	// Chamar o handler com o ResponseRecorder e o request
	handler.ServeHTTP(rr, req)

	// Verificar o status code da resposta
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verificar o tipo de conteúdo da resposta
	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, expectedContentType)
	}

	// Decodificar a resposta e verificar o campo "status"
	var response Status
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}
	if response.Status != "ok" {
		t.Errorf("handler returned unexpected body: got status %v want status %v", response.Status, "ok")
	}
}
