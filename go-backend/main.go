package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	// Load .env file if it exists (ignore errors if it doesn't)
	_ = godotenv.Load()

	// Initialize the service and handler
	emailService := NewEmailGeneratorService()

	// Warn if the API credentials aren't set
	if emailService.GroqAPIKey == "" {
		log.Println("WARNING: GROQ_API_KEY environment variable is not set! The API calls will fail.")
	}

	emailHandler := NewEmailHandler(emailService)

	// Create a new ServeMux
	mux := http.NewServeMux()
	mux.HandleFunc("/email/generate", emailHandler.HandleGenerateEmail)

	// Setup CORS matching the Java implementation
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		MaxAge:         3600,
	})

	// Wrap the mux with the CORS handler
	handler := c.Handler(mux)

	// Determine port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting Email Generator Go Backend on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}
