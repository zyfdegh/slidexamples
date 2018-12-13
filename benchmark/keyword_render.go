package benchmark

import (
	"strings"

	"meipian.cn/review/model/types"
)

// RenderKeywords sets color on keywords in content as HTML.
//
// content - the string to be rendered
// keywords - the matching substrings
// color - red, blue
func RenderKeywords(content string, keywords []string, color types.Color) string {
	var pairs []string
	for _, v := range keywords {
		pairs = append(pairs, v, renderColor(v, string(color)))
	}
	r := strings.NewReplacer(pairs...)
	return r.Replace(content)
}

// renderColor - add HTML color tags around str
func renderColor(str, color string) string {
	if len(str) == 0 {
		return ""
	}
	// return fmt.Sprintf("<font color=\"%s\">%s</font>", color, str)

	// faster
	return "<font color=\"" + color + "\">" + str + "</font>"
}
