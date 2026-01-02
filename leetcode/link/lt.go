package link

import (
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/group-anagrams/?envType=study-plan-v2&envId=top-100-liked
func groupAnagrams(strs []string) [][]string {
	var tmp map[string][]string = make(map[string][]string)

	for _, s := range strs {
		d := make([]int, 26)
		for _, c := range s {
			index := c - 'a'
			d[index] = d[index] + 1
		}
		var index string
		for k, v := range d {
			index = index + "_" + strconv.FormatInt(int64(k), 10) + ":" + strconv.FormatInt(int64(v), 10)
		}
		tmp[index] = append(tmp[index], s)
	}

	var result [][]string
	for _, v := range tmp {
		result = append(result, v)
	}

	return result
}

// https://leetcode.cn/problems/longest-consecutive-sequence/description/?envType=study-plan-v2&envId=top-100-liked
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var tmp = make(map[int]bool, 10000+1)
	var hasFound = make(map[int]bool, 10000+1)
	for _, n := range nums {
		tmp[n] = true
	}

	var result = 1
	for _, n := range nums {
		var tmpResult = 1
		hasFound[n] = true
		for m := n + 1; tmp[m] && !hasFound[m]; m++ {
			tmpResult++
			hasFound[m] = true
		}
		for j := n - 1; tmp[j] && !hasFound[j]; j-- {
			tmpResult++
			hasFound[j] = true
		}
		if result < tmpResult {
			result = tmpResult
		}
	}
	return result
}

// https://leetcode.cn/problems/3sum/description/?envType=study-plan-v2&envId=top-100-liked
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)
	var numMap = make(map[int]int, 3000)

	var result [][]int
	var hasFound = make(map[string]bool)
	for k, v := range nums {
		numMap[v] = k
	}
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		for j := i + 1; j < len(nums)-1; j++ {
			target := 0 - (nums[i] + nums[j])
			if target < nums[j] {
				break
			}
			k, ok := numMap[target]
			if !ok {
				continue
			} else {
				if i == k || j == k {
					continue
				}
			}
			var tmpResult []int
			tmpResult = append(tmpResult, nums[i])
			tmpResult = append(tmpResult, nums[j])
			tmpResult = append(tmpResult, target)
			var index string
			for _, v := range tmpResult {
				index = index + ":" + strconv.FormatInt(int64(v), 10)
			}

			_, ok = hasFound[index]
			if !ok {
				hasFound[index] = true
				result = append(result, tmpResult)
			}
		}
	}

	return result
}
