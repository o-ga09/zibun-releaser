package parse

import (
	"context"
	"io"
	"os"
	"strings"

	"github.com/o-ga09/zibun-releaser/internal/generate"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
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

	// goldmarkのパーサーを作成
	md := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
		),
	)

	// コンテキストを作成
	context := parser.NewContext()
	// Markdownをパース
	_ = md.Parser().Parse(text.NewReader(data), parser.WithContext(context))
	// フロントマターを取得
	metaData := meta.Get(context)

	parser := goldmark.DefaultParser()
	root := parser.Parse(text.NewReader(data))

	mdMap := make(map[string][]string)
	nowKey := ""

	err = ast.Walk(root, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		switch n.Kind() {
		case ast.KindHeading:
			text := string(n.Text(data))
			heading := n.(*ast.Heading)
			if heading.Level == 2 || heading.Level == 3 {
				nowKey = strings.TrimSpace(text)
			}
		case ast.KindListItem:
			if nowKey != "" {
				text := string(n.Text(data))
				mdMap[nowKey] = append(mdMap[nowKey], text)
			}
		}

		return ast.WalkContinue, nil
	})
	if err != nil {
		return generate.ReleaseNote{}, err
	}

	result := generate.ReleaseNote{
		Title:          metaData["Release Title"].(string),
		Version:        metaData["Release Version"].(string),
		Date:           metaData["Release Date"].(string),
		Overview:       "",
		Change:         mdMap["変更内容"],
		PR:             mdMap["PR"],
		BreakingChange: mdMap["破壊的変更"],
		Fixture:        mdMap["修正内容"],
		Issue:          mdMap["既知の問題"],
		Keep:           mdMap["🌻 Keep"],
		Problem:        mdMap["😨 Problem"],
		Try:            mdMap["🌈 Try"],
		Other:          mdMap["その他"],
	}
	return result, nil
}
