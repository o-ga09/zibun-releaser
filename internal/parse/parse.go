package parse

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/o-ga09/zibun-releaser/internal/generate"
	"github.com/russross/blackfriday"
)

func ParseMarkdown(ctx context.Context, path string) (generate.ReleaseNote, error) {
	file, err := os.Open(path)
	if err != nil {
		return generate.ReleaseNote{}, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return generate.ReleaseNote{}, err
	}
	output := blackfriday.MarkdownCommon(data)
	log.Printf("output: %s", output)
	return generate.ReleaseNote{}, nil
}
