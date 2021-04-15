package leetcode

//longest-substring-without-repeating-characters
func LengthOfLongestSubstring(s string) int {
	temp := make(map[string]int)
	len := 0
	len_temp := 0
labe:
	for index, c := range s {
		if v, ok := temp[string(c)]; !ok {
			len_temp += 1
			temp[string(c)] = index
		} else {
			temp = make(map[string]int)
			if len_temp > len {
				len = len_temp
			}
			s = s[v+1:]
			len_temp = 0
			goto labe
		}
	}

	if len < len_temp {
		len = len_temp
	}
	return len
}
