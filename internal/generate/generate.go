package generate

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"text/template"
)

type ReleaseNote struct {
	Version        string
	Title          string
	Date           string
	Overview       string
	PR             []string
	Change         []string
	Fixture        []string
	BreakingChange []string
	Issue          []string
	Keep           []string
	Problem        []string
	Try            []string
	Other          []string
}

func GenerateReleaseNote(ctx context.Context, releaseNote ReleaseNote) (string, error) {
	const templateText = `
Release Note: {{ .Version }}

{{ .Title }}

## 概要

{{ .Overview }}

## 変更内容
{{ range .Change }}
  - {{ . }}
{{- end }}

## PR
{{ range .PR }}
  - PR: #{{ . }}
{{- end }}

## 破壊的変更
{{ range .BreakingChange }}
  - {{ . }}
{{- end }}

## 修正内容
{{ range .Fixture }}
  - {{ . }}
{{- end }}

## 既知の問題
{{ range .Issue }}
  - {{ . }}
{{- end }}

## KPT
### 🌻 Keep
{{ range .Keep }}
  - {{ . }}
{{- end }}

### 😨 Problem
{{ range .Problem }}
  - {{ . }}
{{- end }}

### 🌈 Try
{{ range .Try }}
  - {{ . }}
{{- end }}

## その他
{{ range .Other }}
  - {{ . }}
{{- end }}`

	// テンプレートを解析
	tmpl, err := template.New("releasenote").Parse(templateText)
	if err != nil {
		fmt.Println("テンプレートの解析エラー:", err)
		return "", err
	}

	var buf bytes.Buffer

	// テンプレートを実行し、結果をバッファに書き込む
	err = tmpl.Execute(&buf, releaseNote)
	if err != nil {
		return "", fmt.Errorf("テンプレートの実行エラー: %w", err)
	}

	// バッファの内容を文字列として返す
	return buf.String(), nil
}

func CreateReleaseNoteFile(ctx context.Context, path string, data string) error {
	if os.Getenv("ENV") == "CI" {
		path = fmt.Sprintf("/tmp/%s", path)
	} else {
		path = fmt.Sprintf("../../testdata/%s", path)
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}
