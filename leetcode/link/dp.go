package link

func canJump(nums []int) bool {
	var cache map[int]bool = make(map[int]bool)
	var dfs func(nums []int, begin int, cache map[int]bool) bool
	maxIndex := len(nums) - 1
	dfs = func(nums []int, begin int, cache map[int]bool) bool {
		if begin > maxIndex {
			cache[begin] = false
			return false
		} else if begin == maxIndex {
			cache[begin] = true
			return true
		}
		if nums[begin] == 0 {
			cache[begin] = false
			return false
		}
		if nums[begin]+begin >= maxIndex {
			cache[begin] = true
			return true
		}

		if isCan, ok := cache[begin]; ok {
			return isCan
		}

		for step := 1; step <= nums[begin]; step++ {
			if begin+step > len(nums)-1 {
				break
			}
			isCan := dfs(nums, begin+step, cache)
			if isCan {
				cache[begin] = true
				return true
			}
		}
		cache[begin] = false
		return false
	}
	return dfs(nums, 0, cache)
}

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	var maxIndex = len(nums) - 1
	var cacheTable map[int]int = make(map[int]int)
	var dfs func(nums []int, begin int, cacheTable map[int]int) int
	dfs = func(nums []int, begin int, cacheTable map[int]int) int {
		if begin > maxIndex {
			cacheTable[begin] = -1
			return -1
		}
		if nums[begin] == 0 {
			cacheTable[begin] = -1
			return -1
		}

		if isCan, ok := cacheTable[begin]; ok {
			return isCan
		}

		var minNext int = -1
		for step := 1; step <= nums[begin]; step++ {
			if nums[begin]+begin >= maxIndex {
				cacheTable[begin] = 1
				return 1
			}
			next := dfs(nums, begin+step, cacheTable)
			if next != -1 {
				if minNext == -1 {
					minNext = next
				} else {
					if minNext > next {
						minNext = next
					}
				}
			}
		}

		if minNext != -1 {
			minNext += 1
		}

		cacheTable[begin] = minNext
		return minNext
	}

	return dfs(nums, 0, cacheTable)
}

func uniquePaths(m int, n int) int {
	var dfs func(i, j int) int
	var cache [][]int = make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
	}
	dfs = func(i, j int) int {
		if i == m-1 && j == n-1 {
			return 1
		} else if i == m-1 {
			return 1
		} else if j == n-1 {
			return 1
		}

		if cache[i][j] != 0 {
			return cache[i][j]
		}

		ret := dfs(i+1, j) + dfs(i, j+1)
		cache[i][j] = ret
		return ret
	}

	return dfs(0, 0)
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	var cache [][]int = make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if obstacleGrid[i][j] == 1 {
			return 0
		}

		if i == m-1 && j == n-1 && obstacleGrid[i][j] == 0 {
			return 1
		} else if i == m-1 && j == n-1 && obstacleGrid[i][j] == 1 {
			return 0
		}

		if cache[i][j] != 0 {
			return cache[i][j]
		}

		var ret int
		if i < m-1 && obstacleGrid[i+1][j] == 0 {
			ret += dfs(i+1, j)
		}

		if j < n-1 && obstacleGrid[i][j+1] == 0 {
			ret += dfs(i, j+1)
		}
		cache[i][j] = ret
		return ret
	}
	return dfs(0, 0)
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	var cache [][]int = make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if cache[i][j] != 0 {
			return cache[i][j]
		}
		var ret int
		if i < m-1 && j < n-1 {
			down := grid[i][j] + dfs(i+1, j)
			right := grid[i][j] + dfs(i, j+1)
			if down > right {
				ret = right
			} else {
				ret = down
			}
		} else if i == m-1 && j < n-1 {
			ret = grid[i][j] + dfs(i, j+1)
		} else if i < m-1 && j == n-1 {
			ret = grid[i][j] + dfs(i+1, j)
		} else if i == m-1 && j == n-1 {
			ret = grid[i][j]
		} else {
			return 0
		}
		cache[i][j] = ret
		return ret
	}
	return dfs(0, 0)
}

func lengthOfLIS(nums []int) int {
	var nlen int = len(nums)
	var dp []int = make([]int, nlen)
	var ret int = 0
	for i := 0; i < nlen; i++ {
		if i == 0 {
			dp[i] = 1
			ret = 1
		} else {
			var maxLen int
			for j := i - 1; j >= 0; j-- {
				if nums[i] > nums[j] {
					if maxLen == 0 {
						maxLen = dp[j]
					} else {
						if maxLen < dp[j] {
							maxLen = dp[j]
						}
					}
				}
			}
			if maxLen == 0 {
				dp[i] = 1
			} else {
				dp[i] = maxLen + 1
			}

			if ret < dp[i] {
				ret = dp[i]
			}
		}
	}
	return ret
}

func maxSubArray(nums []int) int {
	var dp []int = make([]int, len(nums))
	var maxRet int
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[i] = nums[i]
			maxRet = dp[i]
		} else {
			if dp[i-1] < 0 {
				dp[i] = nums[i]
			} else {
				dp[i] = dp[i-1] + nums[i]
			}
			if maxRet < dp[i] {
				maxRet = dp[i]
			}
		}
	}
	return maxRet
}

func maxSubArrayEx(nums []int) int {
	var dfs func(i int) int
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var maxRet int
	var dp []int = make([]int, len(nums))
	dfs = func(i int) int {
		if dp[i] != 0 {
			return dp[i]
		}

		var ret int
		if i == 0 {
			ret = nums[i]
		} else {
			ret = max(nums[i], dfs(i-1)+nums[i])
		}
		dp[i] = ret
		return ret
	}

	dfs(len(nums) - 1)
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			maxRet = dp[i]
		} else {
			if maxRet < dp[i] {
				maxRet = dp[i]
			}
		}
	}

	return maxRet
}

func longestCommonSubsequence(text1 string, text2 string) int {
	return 0
}
