package link

import (
	"sort"
)

func letterCombinations(digits string) []string {
	dlen := len(digits)
	var path []byte = make([]byte, dlen)
	var result []string
	if dlen == 0 {
		return nil
	}

	var d2s map[byte]string = map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxzy",
	}

	var dfs func(i int)
	dfs = func(i int) {
		if i == dlen {
			result = append(result, string(path))
			return
		}
		str := d2s[digits[i]]
		for j := 0; j < len(str); j++ {
			path[i] = str[j]
			dfs(i + 1)
		}
	}
	dfs(0)
	return result
}

func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	nlen := 2 * n
	var path []byte = make([]byte, nlen)
	var result []string
	var leftNum int
	var rightNum int
	var dfs func(i int)
	dfs = func(i int) {
		if i == nlen {
			result = append(result, string(path))
			return
		}

		for _, v := range []byte{'(', ')'} {
			if v == '(' {
				leftNum++
			} else {
				rightNum++
			}
			path[i] = v
			if leftNum >= rightNum && leftNum <= n {
				dfs(i + 1)
			}
			if v == '(' {
				leftNum--
			} else {
				rightNum--
			}
		}
	}
	dfs(0)
	return result
}

func combinationSum(candidates []int, target int) [][]int {
	var path []int
	var sum int
	var result [][]int
	var dfs func(start int)
	dfs = func(start int) {
		if sum == target {
			var newPath []int = make([]int, len(path))
			copy(newPath, path)
			result = append(result, newPath)
			return
		}

		if sum > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			v := candidates[i]
			sum += v
			path = append(path, v)
			dfs(i)
			sum -= v
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return result
}

func combinationSum2(candidates []int, target int) [][]int {
	var path []int
	var sum int
	var result [][]int
	var dfs func(start int)
	var hasFound map[string]bool = make(map[string]bool)
	sort.Ints(candidates)
	dfs = func(start int) {
		if sum == target {
			var newPath []int = make([]int, len(path))
			copy(newPath, path)
			var key []rune = make([]rune, len(newPath))
			for i, v := range newPath {
				key[i] = rune(v)
			}
			if _, ok := hasFound[string(key)]; ok {
				return
			}
			hasFound[string(key)] = true
			result = append(result, newPath)
			return
		}

		if sum > target {
			return
		}

		var lastSum int = 0
		for i := start; i < len(candidates); i++ {
			v := candidates[i]
			path = append(path, v)
			sum += v
			if sum != lastSum {
				lastSum = sum
				dfs(i + 1)
			}
			sum -= v
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return result
}

func permute(nums []int) [][]int {
	nLen := len(nums)
	if nLen == 0 {
		return nil
	}

	var path []int = make([]int, nLen)
	var result [][]int
	var dataMap map[int]bool = make(map[int]bool)
	for _, num := range nums {
		dataMap[num] = true
	}

	var dfs func(i int)
	dfs = func(i int) {
		if i == nLen {
			newPath := make([]int, nLen)
			copy(newPath, path)
			result = append(result, newPath)
			return
		}

		for k, v := range dataMap {
			if v {
				path[i] = k
				dataMap[k] = false
				dfs(i + 1)
				dataMap[k] = true
			}
		}

	}
	dfs(0)
	return result
}
