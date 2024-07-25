package summarize

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
)

func SummarizeDocument(ctx context.Context, document string) (string, error) {
	apikey := os.Getenv("APIKEY")
	opts := googleai.WithAPIKey(apikey)
	model, err := googleai.New(ctx, opts)
	if err != nil {
		return "", err
	}

	text := `Artificial intelligence (AI) is intelligence—perceiving, synthesizing, and inferring information—demonstrated by machines, as opposed to intelligence displayed by animals and humans. Example tasks in which this is done include speech recognition, computer vision, translation between (natural) languages, as well as other mappings of inputs.

			AI applications include advanced web search engines (e.g., Google Search), recommendation systems (used by YouTube, Amazon, and Netflix), understanding human speech (such as Siri and Alexa), self-driving cars (e.g., Waymo), generative or creative tools (ChatGPT and AI art), automated decision-making, and competing at the highest level in strategic game systems (such as chess and Go).`

	prompt := fmt.Sprintf(`以下の文章を要約してください:
				%s

				要約の条件:
				1. 日本語で書いてください。
				2. 100字程度で簡潔にまとめてください。
				3. 文章の先頭に[summarized]を付けてください。
				4. 重要なポイントのみを含めてください。`, text)

	result, err := model.Call(context.Background(), prompt, llms.WithMaxTokens(200))

	if err != nil {
		log.Fatalf("Failed to generate summary: %v", err)
	}

	if !strings.Contains(result, "[summarized]") {
		result = fmt.Sprintf("[summarized]\n\n%s", result)
	}
	return result, nil
}
