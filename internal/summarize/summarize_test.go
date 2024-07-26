package summarize

import (
	"context"
	"os"
	"strings"
	"testing"
)

func TestSummarizeDocument(t *testing.T) {
	t.Setenv("APIKEY", os.Getenv("APIKEY"))
	ctx := context.Background()
	document := "これはテスト文章です。"
	cases := []struct {
		name     string
		expected string
		err      error
		iserr    bool
	}{
		{name: "文章を要約できる場合", expected: "[summarized]", err: nil, iserr: false},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SummarizeDocument(ctx, document)
			if (err != nil) != tt.iserr {
				t.Errorf("got: %v, but expected: %v", err, tt.err)
			}
			t.Log(result)
			if !strings.Contains(result, tt.expected) {
				t.Errorf("got: %v, but expected: %v is contained", result, tt.expected)
			}
		})
	}
}
