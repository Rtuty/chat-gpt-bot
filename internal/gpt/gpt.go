package gpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	ChatGPTAPI    = "https://api.openai.com/v1/chat/completions"
	ChatGPTApiKey = os.Getenv("GPT_API_KEY")
)

type ChatGPTRequest struct {
	Model     string `json:"model"`
	Prompt    string `json:"prompt"`
	MaxTokens int    `json:"max_tokens"`
}

type ChatGPTResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func GetChatGPTResponse(prompt string) (string, error) {
	requestBody, err := json.Marshal(ChatGPTRequest{
		Model:     "gpt-3.5-turbo",
		Prompt:    prompt,
		MaxTokens: 150,
	})
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", ChatGPTAPI, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", ChatGPTApiKey))

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatalf("GetChatGPTResponse defer body close error: %v", err)
		}
	}()

	var chatGPTResponse ChatGPTResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatGPTResponse); err != nil {
		return "", err
	}

	return chatGPTResponse.Choices[0].Text, nil
}
