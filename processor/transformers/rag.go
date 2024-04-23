package transformers

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/schema"
)

func RagResponse(query string, referData []schema.Document, llm *ollama.LLM) string {

	clone := chains.LoadStuffQA(llm)
	reply, err := chains.Call(context.TODO(), clone,
		map[string]any{
			"input_documents": referData,
			"question": fmt.Sprintf(`The provided document contains information about the user's usual patterns. 
	If content exists, refer to it to respond in a friendly and casual manner; 
	if not, answer as if you're chatting with a friend.\n
	User: %s`, query),
		})
	if err != nil {
		return ""
	}
	return reply["text"].(string)

}
