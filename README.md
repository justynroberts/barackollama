# Barack Ollama

Barack Ollama is a simple Go-based command-line interface (CLI) for interacting with the Ollama API. This tool allows you to send prompts to a specified model and receive responses, with support for both single-prompt and interactive modes.

## Features

Single Prompt Mode: Send a single prompt to the model and get a response.

Interactive Mode: Start an interactive chat session with the model.

Clean Output: Automatically removes <think> and </think> tags from the response.

Pre-built Executable: A pre-compiled executable is provided for easy use.

## Prerequisites

Ollama API: The Ollama API must be running and accessible (default: http://localhost:11434/api/generate).

Executable: If you're using the pre-built executable, no additional dependencies are required.

## Installation

Using the Pre-built Executable
Download the ollama-cli executable from the releases page.

Make the binary executable:
```bash
chmod +x ollama-cli
```

(Optional) Move the binary to a directory in your PATH for global access:
```bash
sudo mv ollama-cli /usr/local/bin/
```

Building from Source
Clone the repository:
```bash
git clone https://github.com/your-username/barack-ollama.git
cd barack-ollama
```

Build the executable:
```bash
go build -o ollama-cli
```

Make the binary executable:
```bash
chmod +x ollama-cli
```

## Usage

Single Prompt Mode
Send a single prompt to the model:
```bash
./ollama-cli --model <model-name> --prompt "<your-prompt>"
```

Example:
```bash
./ollama-cli --model llama2 --prompt "What is the capital of France?"
```

Interactive Mode
Start an interactive chat session with the model:
```bash
./ollama-cli --model <model-name> --interactive
```

Example:
```bash
./ollama-cli --model llama2 --interactive
```

In interactive mode, type exit to quit the session.

Example Output
Single Prompt Mode
```bash
$ ./ollama-cli --model llama2 --prompt "What is the capital of France?"
The capital of France is Paris.
```

Interactive Mode
```bash
$ ./ollama-cli --model llama2 --interactive
Starting interactive session with model: llama2
Type 'exit' to quit.
You: What is the capital of France?
Model: The capital of France is Paris.
You: Tell me a joke.
Model: Why don't scientists trust atoms? Because they make up everything!
You: exit
```

## Configuration

API Endpoint: By default, the CLI uses http://localhost:11434/api/generate. If your Ollama API is running on a different endpoint, update the ollamaAPIURL constant in main.go.

## Troubleshooting

Error: Connection Refused: Ensure the Ollama API is running and accessible at the specified endpoint.

Error: Invalid Model: Verify that the model name is correct and supported by the Ollama API.

Error: JSON Parsing Failed: Check the API response format and update the ResponsePayload struct if necessary.

## Contributing

Contributions are welcome! If you'd like to add features, fix bugs, or improve the code, follow these steps:

Fork the repository.

Create a new branch for your changes:
```bash
git checkout -b feature/your-feature-name
```

Commit your changes:
```bash
git commit -m "Add your commit message here"
```

Push your changes to your fork:
```bash
git push origin feature/your-feature-name
```

Open a pull request on the original repository.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgments

Built with Go for simplicity and performance.

Inspired by the Ollama API for interacting with language models.

How to Use
Copy the escaped Markdown content above.

Paste it into your README.md file.

Save the file.

When rendered as Markdown, it will display correctly. When viewed as plain text, it will remain uninterpreted.

Let me know if you need further assistance! 🚀
