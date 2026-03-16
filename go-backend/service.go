package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const groqApiURL = "https://api.groq.com/openai/v1/chat/completions"
const defaultModel = "llama-3.3-70b-versatile" // Updated to Llama 3.3 70B

type EmailGeneratorService struct {
	Client     *http.Client
	GroqAPIKey string
}

func NewEmailGeneratorService() *EmailGeneratorService {
	// Attempt to load from environment variables
	apiKey := os.Getenv("GROQ_API_KEY")

	return &EmailGeneratorService{
		Client:     &http.Client{},
		GroqAPIKey: apiKey,
	}
}

func (s *EmailGeneratorService) GenerateEmailReply(req EmailRequest) (string, error) {
	if s.GroqAPIKey == "" {
		return "", errors.New("GROQ_API_KEY is not configured")
	}

	prompt := s.buildPrompt(req)

	// Construct the Groq API request body
	groqReq := GroqRequest{
		Model: defaultModel,
		Messages: []GroqMessage{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	reqBody, err := json.Marshal(groqReq)
	if err != nil {
		return "", fmt.Errorf("failed to marshal groq request: %w", err)
	}

	// Create the HTTP Request
	httpReq, err := http.NewRequest(http.MethodPost, groqApiURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", fmt.Errorf("failed to create http request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+s.GroqAPIKey)

	// Execute the request
	resp, err := s.Client.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("failed to execute http request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	// Handle specific HTTP error codes
	if resp.StatusCode == 429 {
		return "Sorry, the Groq API rate limit has been exceeded. Please try again after some time.", nil
	} else if resp.StatusCode != http.StatusOK {
		return fmt.Sprintf("Error from Groq API: %d - %s", resp.StatusCode, string(bodyBytes)), nil
	}

	// Extract text from the successful response
	return s.extractResponse(bodyBytes)
}

func (s *EmailGeneratorService) buildPrompt(req EmailRequest) string {
	tonePart := ""
	if strings.TrimSpace(req.Tone) != "" {
		tonePart = "Use a " + req.Tone + " tone. "
	}

	// Check if this is a request to generate a new email from a subject
	if strings.TrimSpace(req.Subject) != "" {
		return "Generate a professional email body based on the following subject line. " + tonePart +
			"Do not include a subject line in the response. Subject: " + req.Subject
	}

	// Default to generating a reply to existing content
	return "Generate a professional email reply. " + tonePart +
		"Do not include a subject line. Original email: " + req.EmailContent
}

func (s *EmailGeneratorService) extractResponse(body []byte) (string, error) {
	var groqResp GroqResponse
	err := json.Unmarshal(body, &groqResp)
	if err != nil {
		return "", fmt.Errorf("failed to parse groq response: %w", err)
	}

	if len(groqResp.Choices) > 0 && groqResp.Choices[0].Message.Content != "" {
		return groqResp.Choices[0].Message.Content, nil
	}

	return "No valid response from API", nil
}
