package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type EmailHandler struct {
	Service *EmailGeneratorService
}

func NewEmailHandler(service *EmailGeneratorService) *EmailHandler {
	return &EmailHandler{
		Service: service,
	}
}

func (h *EmailHandler) HandleGenerateEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	log.Printf("Received request: %+v\n", req)

	// Validation
	if strings.TrimSpace(req.Subject) == "" && strings.TrimSpace(req.EmailContent) == "" {
		log.Println("Invalid request: Both subject and email content are empty.")
		http.Error(w, "Invalid email content: Both subject and email content are empty", http.StatusBadRequest)
		return
	}

	// Generate Reply
	response, err := h.Service.GenerateEmailReply(req)
	if err != nil {
		log.Printf("Error generating reply: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
