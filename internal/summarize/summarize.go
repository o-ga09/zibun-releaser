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

	prompt := fmt.Sprintf(`以下の文章を要約してください:
				%s

				要約の条件:
				1. 日本語で書いてください。
				2. 100字程度で簡潔にまとめてください。
				3. 文章の先頭に[summarized]を付けてください。
				4. 重要なポイントのみを含めてください。`, document)

	result, err := model.Call(context.Background(), prompt, llms.WithMaxTokens(200))

	if err != nil {
		log.Fatalf("Failed to generate summary: %v", err)
	}

	if !strings.Contains(result, "[summarized]") {
		result = fmt.Sprintf("[summarized]\n\n%s", result)
	}
	return result, nil
}
