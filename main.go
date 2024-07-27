package main

import (
	"context"
	"log"
	"os"

	"github.com/o-ga09/zibun-releaser/internal/generate"
	"github.com/o-ga09/zibun-releaser/internal/markdown"
	"github.com/o-ga09/zibun-releaser/internal/parse"
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

	parseMd, err := parse.ParseMarkdown(ctx, path)
	if err != nil {
		log.Fatalf("failed to parse markdown: %v", err)
	}

	parseMd.Overview = summary
	releaseNoteText, err := generate.GenerateReleaseNote(ctx, parseMd)
	if err != nil {
		log.Fatalf("failed to generate release note: %v", err)
	}

	err = generate.CreateReleaseNoteFile(ctx, "releasenote.md", releaseNoteText)
	if err != nil {
		log.Fatalf("failed to create release note file: %v", err)
	}
}
