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
		{name: "1行のファイルの場合", expected: generate.ReleaseNote{}, iserr: false},
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
