package benchmark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderColor(t *testing.T) {
	var testcases = []struct {
		str   string
		color string

		expect string
	}{
		{"", "", ""},
		{"str1", "red", "<font color=\"red\">str1</font>"},
		{"str2", "blue", "<font color=\"blue\">str2</font>"},
	}

	for _, c := range testcases {
		got := renderColor(c.str, c.color)
		assert.Equal(t, c.expect, got)
	}
}

func BenchmarkRenderKeywords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RenderKeywords("习近平主席王岐山副主席", []string{"习近平", "王岐山"}, "red")
	}
}
