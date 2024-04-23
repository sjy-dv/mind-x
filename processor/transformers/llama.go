package transformers

import "github.com/tmc/langchaingo/llms/ollama"

func LoadLLAMA3() (*ollama.LLM, error) {
	llm, err := ollama.New(ollama.WithModel("llama3:latest"))
	if err != nil {
		return nil, err
	}
	return llm, nil
}
