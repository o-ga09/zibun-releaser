package markdown

import (
	"context"
	"io"
	"os"
)

func ReadMarkdown(cyx context.Context, path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return string(data), nil
}
