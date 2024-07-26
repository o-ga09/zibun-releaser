package main

import (
	"context"
	"log"
	"os"

	"github.com/o-ga09/zibun-releaser/internal/markdown"
	"github.com/o-ga09/zibun-releaser/internal/summarize"
)

func main() {
	ctx := context.Background()
	path := os.Getenv("FILEPATH")
	data, err := markdown.ReadMarkdown(ctx, path)
	if err != nil {
		log.Fatalf("failed to read markdown: %v", err)
	}
	summary, err := summarize.SummarizeDocument(ctx, data)
	if err != nil {
		log.Fatalf("failed to summarize document: %v", err)
	}
	log.Printf("summary: %v", summary)
}
