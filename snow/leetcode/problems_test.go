package leetcode

import (
	"testing"
)

func TestLengthOfLongestSubString(t *testing.T) {
	text := "abcedbcde"
	result := LengthOfLongestSubstring(text)
	t.Logf("len_res:%d", result)
}
