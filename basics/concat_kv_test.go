package basics

import (
	"testing"
)

func TestConcatKV(t *testing.T) {
	var testcases = []struct {
		key    string
		values []int64

		expect string
	}{
		{"", []int64{}, ""},
		{"array", []int64{int64(1)}, "array=1"},
		{"array", []int64{int64(1), int64(2)}, "array=1,2"},
		{"ids", []int64{int64(1234567890), int64(4567890123), int64(5678901234)}, "ids=1234567890,4567890123,5678901234"},
	}

	for _, v := range testcases {
		got := ConcatKV(v.key, v.values)

		if got != v.expect {
			t.Fail()
		}
	}
}
