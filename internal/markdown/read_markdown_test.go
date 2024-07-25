package markdown

import (
	"context"
	"fmt"
	"testing"
)

func TestReadMarkdown(t *testing.T) {
	ctx := context.Background()
	path_template := "../../testdata/test_%d.md"
	cases := []struct {
		name     string
		expected string
		iserr    bool
	}{
		{name: "1行のファイルの場合", expected: "# test", iserr: false},
		{name: "複数行のファイルの場合", expected: "# test\n\n## title\n\n### summary\n\n", iserr: false},
		{name: "ファイルが存在しないの場合", expected: "", iserr: true},
	}

	for i, tt := range cases {
		path := fmt.Sprintf(path_template, i+1)
		result, err := ReadMarkdown(ctx, path)
		if (err != nil) != tt.iserr {
			t.Errorf("got: %v err is not nil but expected err is nil", err)
		}
		if result != tt.expected {
			t.Errorf("got %v, but expected %v", result, tt.expected)
		}
	}
}
