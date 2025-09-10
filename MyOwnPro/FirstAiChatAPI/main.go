// error fix working
// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// )

// // Request body format
// type ChatRequest struct {
// 	Model     string        `json:"model"`
// 	Messages  []ChatMessage `json:"messages"`
// 	MaxTokens int           `json:"max_tokens"`
// }

// type ChatMessage struct {
// 	Role    string `json:"role"`
// 	Content string `json:"content"`
// }

// // Response format (simplified)
// type ChatResponse struct {
// 	Choices []struct {
// 		Message ChatMessage `json:"message"`
// 	} `json:"choices"`
// }

// func main() {
// 	apiKey := os.Getenv("OPENROUTER_API_KEY")
// 	if apiKey == "" {
// 		fmt.Println("❌ Please set OPENROUTER_API_KEY in your environment variables.")
// 		return
// 	}

// 	url := "https://openrouter.ai/api/v1/chat/completions"

// 	// Create request body with max_tokens
// 	reqBody := ChatRequest{
// 		Model: "openai/gpt-4.1-mini",
// 		Messages: []ChatMessage{
// 			{Role: "user", Content: "Hello AI bot, can you hear me?"},
// 		},
// 		MaxTokens: 500, // 👈 keep it small to fit free credits
// 	}

// 	// Encode to JSON
// 	jsonData, err := json.Marshal(reqBody)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Make POST request
// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		panic(err)
// 	}
// 	req.Header.Set("Authorization", "Bearer "+apiKey)
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()

// 	// Read response
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if resp.StatusCode != 200 {
// 		fmt.Println("❌ Error:", resp.Status)
// 		fmt.Println(string(body))
// 		return
// 	}

// 	// Parse JSON
// 	var chatResp ChatResponse
// 	err = json.Unmarshal(body, &chatResp)
// 	if err != nil {
// 		panic(err)
// 	}

//		if len(chatResp.Choices) > 0 {
//			fmt.Println("✅ Bot reply:", chatResp.Choices[0].Message.Content)
//		} else {
//			fmt.Println("⚠️ No response from model")
//		}
//	}
//
// user input
// package main

// import (
// 	"bufio"
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"strings"
// )

// // Request body format
// type ChatRequest struct {
// 	Model     string        `json:"model"`
// 	Messages  []ChatMessage `json:"messages"`
// 	MaxTokens int           `json:"max_tokens"`
// }

// type ChatMessage struct {
// 	Role    string `json:"role"`
// 	Content string `json:"content"`
// }

// // Response format (simplified)
// type ChatResponse struct {
// 	Choices []struct {
// 		Message ChatMessage `json:"message"`
// 	} `json:"choices"`
// }

// func main() {
// 	apiKey := os.Getenv("OPENROUTER_API_KEY")
// 	if apiKey == "" {
// 		fmt.Println("❌ Please set OPENROUTER_API_KEY in your environment variables.")
// 		return
// 	}

// 	url := "https://openrouter.ai/api/v1/chat/completions"
// 	reader := bufio.NewReader(os.Stdin)

// 	for {
// 		// Ask user for input
// 		fmt.Print("\nYou: ")
// 		userInput, _ := reader.ReadString('\n')
// 		userInput = strings.TrimSpace(userInput)

// 		// Exit condition
// 		if strings.ToLower(userInput) == "exit" {
// 			fmt.Println("👋 Goodbye!")
// 			break
// 		}

// 		// Create request body
// 		reqBody := ChatRequest{
// 			Model: "openai/gpt-4.1-mini",
// 			Messages: []ChatMessage{
// 				{Role: "user", Content: userInput},
// 			},
// 			MaxTokens: 500,
// 		}

// 		// Encode to JSON
// 		jsonData, err := json.Marshal(reqBody)
// 		if err != nil {
// 			panic(err)
// 		}

// 		// Make POST request
// 		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
// 		if err != nil {
// 			panic(err)
// 		}
// 		req.Header.Set("Authorization", "Bearer "+apiKey)
// 		req.Header.Set("Content-Type", "application/json")

// 		client := &http.Client{}
// 		resp, err := client.Do(req)
// 		if err != nil {
// 			panic(err)
// 		}
// 		defer resp.Body.Close()

// 		// Read response
// 		body, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			panic(err)
// 		}

// 		if resp.StatusCode != 200 {
// 			fmt.Println("❌ Error:", resp.Status)
// 			fmt.Println(string(body))
// 			continue
// 		}

// 		// Parse JSON
// 		var chatResp ChatResponse
// 		err = json.Unmarshal(body, &chatResp)
// 		if err != nil {
// 			panic(err)
// 		}

// 		if len(chatResp.Choices) > 0 {
// 			fmt.Println("Bot:", chatResp.Choices[0].Message.Content)
// 		} else {
// 			fmt.Println("⚠️ No response from model")
// 		}
// 	}
// }

//.api file

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type apiConfigData struct {
	OpenRouterApiKey string `json:"OpenRouterApiKey"`
}

// Request body format
type ChatRequest struct {
	Model     string        `json:"model"`
	Messages  []ChatMessage `json:"messages"`
	MaxTokens int           `json:"max_tokens"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Response format (simplified)
type ChatResponse struct {
	Choices []struct {
		Message ChatMessage `json:"message"`
	} `json:"choices"`
}

func loadApiConfig(filename string) (apiConfigData, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return apiConfigData{}, err
	}

	var c apiConfigData
	err = json.Unmarshal(data, &c)
	if err != nil {
		return apiConfigData{}, err
	}

	return c, nil
}

func main() {
	// Load API key from .apiConfig file
	config, err := loadApiConfig(".apiConfig")
	if err != nil {
		fmt.Println("❌ Failed to load .apiConfig:", err)
		return
	}

	apiKey := config.OpenRouterApiKey
	if apiKey == "" {
		fmt.Println("❌ API key missing in .apiConfig file")
		return
	}

	url := "https://openrouter.ai/api/v1/chat/completions"
	reader := bufio.NewReader(os.Stdin)

	for {
		// Ask user for input
		fmt.Print("\nYou: ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		// Exit condition
		if strings.ToLower(userInput) == "exit" {
			fmt.Println("👋 Goodbye!")
			break
		}

		// Create request body
		reqBody := ChatRequest{
			Model: "openai/gpt-4.1-mini",
			Messages: []ChatMessage{
				{Role: "user", Content: userInput},
			},
			MaxTokens: 500,
		}

		// Encode to JSON
		jsonData, err := json.Marshal(reqBody)
		if err != nil {
			panic(err)
		}

		// Make POST request
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			panic(err)
		}
		req.Header.Set("Authorization", "Bearer "+apiKey)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// Read response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		if resp.StatusCode != 200 {
			fmt.Println("❌ Error:", resp.Status)
			fmt.Println(string(body))
			continue
		}

		// Parse JSON
		var chatResp ChatResponse
		err = json.Unmarshal(body, &chatResp)
		if err != nil {
			panic(err)
		}

		if len(chatResp.Choices) > 0 {
			fmt.Println("Bot:", chatResp.Choices[0].Message.Content)
		} else {
			fmt.Println("⚠️ No response from model")
		}
	}
}
