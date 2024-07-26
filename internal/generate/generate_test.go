package generate

import (
	"context"
	"os"
	"strings"
	"testing"
)

func TestGenerateReleaseNote(t *testing.T) {
	eexpect_1 := `
Release Note: v1.0.0

プロフィール

## 概要

名前: 山田太郎
年齢: 30歳
住所: 東京
趣味:
  - 読書
  - 旅行
  - 料理

## 変更内容

  - 名前: 山田太郎
  - 年齢: 30歳
  - 住所: 東京

## PR

  - PR: #1
  - PR: #2

## 破壊的変更

  - 名前: 山田太郎
  - 年齢: 30歳
  - 住所: 東京

## 修正内容

  - 読書
  - 旅行
  - 料理

## 既知の問題

  - 読書
  - 旅行
  - 料理

## KPT
### 🌻 Keep

  - 読書
  - 旅行
  - 料理

### 😨 Problem

  - 読書
  - 旅行
  - 料理

### 🌈 Try

  - 読書
  - 旅行
  - 料理

## その他

  - 読書
  - 旅行
  - 料理`

	ctx := context.Background()
	cases := []struct {
		name        string
		releaseNote ReleaseNote
		expected    string
		iserr       bool
	}{
		{
			name: "test case 1",
			releaseNote: ReleaseNote{
				Version:        "v1.0.0",
				Title:          "プロフィール",
				Overview:       "名前: 山田太郎\n年齢: 30歳\n住所: 東京\n趣味:\n  - 読書\n  - 旅行\n  - 料理",
				PR:             []string{"1", "2"},
				Change:         []string{"名前: 山田太郎", "年齢: 30歳", "住所: 東京"},
				Fixture:        []string{"読書", "旅行", "料理"},
				BreakingChange: []string{"名前: 山田太郎", "年齢: 30歳", "住所: 東京"},
				Issue:          []string{"読書", "旅行", "料理"},
				Keep:           []string{"読書", "旅行", "料理"},
				Problem:        []string{"読書", "旅行", "料理"},
				Try:            []string{"読書", "旅行", "料理"},
				Other:          []string{"読書", "旅行", "料理"},
			},
			expected: eexpect_1,
			iserr:    false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GenerateReleaseNote(ctx, tt.releaseNote)
			if (err != nil) != tt.iserr {
				t.Errorf("got: %v err is not nil but expected err is nil", err)
			}
			expectedTrimmed := strings.TrimSpace(tt.expected)
			resultTrimmed := strings.TrimSpace(result)
			if resultTrimmed != expectedTrimmed {
				t.Errorf("got %v, but expected %v", result, tt.expected)
			}
		})
	}
}

func TestCreateReleaseNoteFile(t *testing.T) {
	t.Setenv("ENV", os.Getenv("ENV"))
	ctx := context.Background()
	cases := []struct {
		name     string
		filename string
		data     string
		iserr    bool
	}{
		{
			name:     "test case 1",
			filename: "test.md",
			data:     "## Title \n## Overview\n## Summary\ntest data",
			iserr:    false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateReleaseNoteFile(ctx, tt.filename, tt.data)
			if (err != nil) != tt.iserr {
				t.Errorf("got: %v err is not nil but expected err is nil", err)
			}
		})
	}
}
