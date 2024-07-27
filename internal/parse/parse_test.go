package parse

import (
	"context"
	"fmt"
	"testing"

	"github.com/o-ga09/zibun-releaser/internal/generate"
	"github.com/stretchr/testify/assert"
)

func TestParsemarkdown(t *testing.T) {
	ctx := context.Background()
	path_template := "../../testdata/test_%d.md"
	cases := []struct {
		name     string
		expected generate.ReleaseNote
		iserr    bool
	}{
		{name: "1行のファイルの場合", expected: generate.ReleaseNote{
			Version:        "v1.0.0",
			Title:          "backlog for monthly",
			Date:           "2024-07-31",
			Overview:       "",
			PR:             nil,
			Change:         []string{"コーディングした", "サービス作った"},
			Fixture:        []string{"朝活が習慣化できた"},
			BreakingChange: []string{"転職した"},
			Issue:          []string{"副業やりたい"},
			Keep:           []string{"勉強会いっぱい行った", "サービス作った", "記事書いた"},
			Problem:        []string{"副業探しに苦戦中", "投資に失敗", "SNS苦戦中"},
			Try:            []string{"資格試験に申し込んだ", "サービス作る", "副業探す"},
			Other:          []string{"スター欲しいです"},
		}, iserr: false},
	}

	for i, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			path := fmt.Sprintf(path_template, i+4)
			result, err := ParseMarkdown(ctx, path)
			if (err != nil) != tt.iserr {
				t.Errorf("got: %v err is not nil but expected err is nil", err)
			}
			assert.Equal(t, tt.expected, result)
		})
	}
}
