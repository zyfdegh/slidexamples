package basics

import (
	"fmt"
	"strings"
)

// ConcatKV - combine multiple values to a comma separated string, similar to URL query param.
//
// For exmaple, if key=a, values=[1,2,3], then this will return: a=1,2,3
func ConcatKV(key string, values []int64) string {
	if key == "" || len(values) == 0 {
		return ""
	}
	var str = fmt.Sprintf("%s=", key)
	for _, v := range values {
		str += fmt.Sprintf("%d,", v)
	}
	return strings.TrimRight(str, ",")
}
