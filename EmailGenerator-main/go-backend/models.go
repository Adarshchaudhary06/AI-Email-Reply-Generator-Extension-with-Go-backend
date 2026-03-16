package main

// EmailRequest represents the incoming JSON payload from the Chrome extension
type EmailRequest struct {
	Subject      string `json:"subject"`
	EmailContent string `json:"emailContent"`
	Tone         string `json:"tone"`
}

// GroqMessage represents a message in the chat
type GroqMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// GroqRequest represents the JSON structure to send to the Groq API (OpenAI compatible)
type GroqRequest struct {
	Model    string        `json:"model"`
	Messages []GroqMessage `json:"messages"`
}

// GroqResponse represents the JSON structure returned by the Groq API
type GroqResponse struct {
	Choices []GroqChoice `json:"choices"`
}

type GroqChoice struct {
	Message GroqMessage `json:"message"`
}
