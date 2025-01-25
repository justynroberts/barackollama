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

// Ollama API endpoint (update if necessary)
const ollamaAPIURL = "http://localhost:11434/api/generate"

// RequestPayload defines the structure of the request to the Ollama API
type RequestPayload struct {
        Model  string `json:"model"`
        Prompt string `json:"prompt"`
        Stream bool   `json:"stream"`
}

// ResponsePayload defines the structure of the response from the Ollama API
type ResponsePayload struct {
        Response string `json:"response"`
}

// sendPrompt sends a prompt to the Ollama API and returns the cleaned response
func sendPrompt(model, prompt string) (string, error) {
        // Create the request payload
        payload := RequestPayload{
                Model:  model,
                Prompt: prompt,
                Stream: false, // Set to true for streaming responses
        }

        // Convert the payload to JSON
        jsonData, err := json.Marshal(payload)
        if err != nil {
                return "", fmt.Errorf("error encoding JSON: %v", err)
        }

        // Send the request to the Ollama API
        resp, err := http.Post(ollamaAPIURL, "application/json", bytes.NewBuffer(jsonData))
        if err != nil {
                return "", fmt.Errorf("error sending request: %v", err)
        }
        defer resp.Body.Close()

        // Read the response body
        body, err := io.ReadAll(resp.Body)
        if err != nil {
                return "", fmt.Errorf("error reading response: %v", err)
        }

        // Parse the response
        var response ResponsePayload
        if err := json.Unmarshal(body, &response); err != nil {
                return "", fmt.Errorf("error decoding JSON: %v", err)
        }

        // Remove <think> and </think> tags from the response
        cleanedResponse := strings.ReplaceAll(response.Response, "<think>", "")
        cleanedResponse = strings.ReplaceAll(cleanedResponse, "</think>", "")

        return cleanedResponse, nil
}

// interactiveMode starts an interactive chat session with the model
func interactiveMode(model string) {
        fmt.Printf("Starting interactive session with model: %s\n", model)
        fmt.Println("Type 'exit' to quit.")

        reader := bufio.NewReader(os.Stdin)
        for {
                fmt.Print("You: ")
                prompt, _ := reader.ReadString('\n')
                prompt = strings.TrimSpace(prompt)

                if prompt == "exit" {
                        break
                }

                response, err := sendPrompt(model, prompt)
                if err != nil {
                        fmt.Printf("Error: %v\n", err)
                        continue
                }

                fmt.Printf("Model: %s\n", response)
        }
}

func main() {
        // Parse command-line arguments
        if len(os.Args) < 2 || os.Args[1] != "--model" {
                fmt.Println("Usage: ollama-cli --model <model-name> [--prompt <your-prompt> | --interactive]")
                os.Exit(1)
        }

        model := os.Args[2]

        if len(os.Args) > 3 && os.Args[3] == "--prompt" {
                // Single prompt mode
                if len(os.Args) < 5 {
                        fmt.Println("Error: --prompt requires a prompt argument")
                        os.Exit(1)
                }
                prompt := os.Args[4]

                response, err := sendPrompt(model, prompt)
                if err != nil {
                        fmt.Printf("Error: %v\n", err)
                        os.Exit(1)
                }
                fmt.Println(response)
        } else if len(os.Args) > 3 && os.Args[3] == "--interactive" {
                // Interactive mode
                interactiveMode(model)
        } else {
                fmt.Println("Error: Either --prompt or --interactive must be specified")
                os.Exit(1)
        }
}
