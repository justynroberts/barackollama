# BarackOllama
A simple Go-based command-line interface (CLI) for interacting with the Ollama API. This tool allows you to send prompts to a specified model and receive responses, with support for both single-prompt and interactive modes.

## Features

- **Single Prompt Mode**: Send a single prompt to the model and get a response.
- **Interactive Mode**: Start an interactive chat session with the model.
- **Clean Output**: Automatically removes `<think>` and `</think>` tags from the response.

## Prerequisites

- **Go**: Ensure Go is installed on your system. You can download it from [golang.org](https://golang.org/dl/).
- **Ollama API**: The Ollama API must be running and accessible (default: `http://localhost:11434/api/generate`).

## Installation

1. Clone the repository:
   \`\`\`bash
   git clone https://github.com/your-username/ollama-cli.git
   cd ollama-cli
   \`\`\`

2. Build the CLI:
   \`\`\`bash
   go build -o ollama-cli
   \`\`\`

3. Make the binary executable:
   \`\`\`bash
   chmod +x ollama-cli
   \`\`\`

4. (Optional) Move the binary to a directory in your `PATH` for global access:
   \`\`\`bash
   sudo mv ollama-cli /usr/local/bin/
   \`\`\`

## Usage

### Single Prompt Mode

Send a single prompt to the model:
\`\`\`bash
./ollama-cli --model <model-name> --prompt "<your-prompt>"
\`\`\`

Example:
\`\`\`bash
./ollama-cli --model llama2 --prompt "What is the capital of France?"
\`\`\`

### Interactive Mode

Start an interactive chat session with the model:
\`\`\`bash
./ollama-cli --model <model-name> --interactive
\`\`\`

Example:
\`\`\`bash
./ollama-cli --model llama2 --interactive
\`\`\`

In interactive mode, type `exit` to quit the session.

### Example Output

#### Single Prompt Mode
\`\`\`bash
$ ./ollama-cli --model llama2 --prompt "What is the capital of France?"
The capital of France is Paris.
\`\`\`

#### Interactive Mode
\`\`\`bash
$ ./ollama-cli --model llama2 --interactive
Starting interactive session with model: llama2
Type 'exit' to quit.
You: What is the capital of France?
Model: The capital of France is Paris.
You: Tell me a joke.
Model: Why don't scientists trust atoms? Because they make up everything!
You: exit
\`\`\`

## Configuration

- **API Endpoint**: By default, the CLI uses `http://localhost:11434/api/generate`. If your Ollama API is running on a different endpoint, update the `ollamaAPIURL` constant in `main.go`.

## Contributing

Contributions are welcome! If you'd like to add features, fix bugs, or improve the code, follow these steps:

1. Fork the repository.
2. Create a new branch for your changes:
   \`\`\`bash
   git checkout -b feature/your-feature-name
   \`\`\`
3. Commit your changes:
   \`\`\`bash
   git commit -m "Add your commit message here"
   \`\`\`
4. Push your changes to your fork:
   \`\`\`bash
   git push origin feature/your-feature-name
   \`\`\`
5. Open a pull request on the original repository.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Troubleshooting

- **Error: Connection Refused**: Ensure the Ollama API is running and accessible at the specified endpoint.
- **Error: Invalid Model**: Verify that the model name is correct and supported by the Ollama API.
- **Error: JSON Parsing Failed**: Check the API response format and update the `ResponsePayload` struct if necessary.

---

## Acknowledgments

- Built with Go for simplicity and performance.
- Inspired by the Ollama API for interacting with language models.
