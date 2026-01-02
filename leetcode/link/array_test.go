package link

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	get := lengthOfLongestSubstring("abcabcbb")
	t.Errorf("%d", get)
}
