package link

func lengthOfLongestSubstring(s string) int {
	var c2num map[byte]int = make(map[byte]int)
	var needShink bool
	var maxLen int = 0
	left := 0
	right := 0
	for right < len(s) {
		c := s[right]
		if n, ok := c2num[c]; ok {
			if n >= 1 {
				needShink = true
			} else {
				needShink = false
			}
			c2num[c]++
		} else {
			needShink = false
			c2num[c] = 1
		}
		right++
		if !needShink {
			if maxLen < right-left {
				maxLen = right - left
			}
		}
		for left < right && needShink {
			d := s[left]
			c2num[d]--
			num := c2num[d]
			if num == 1 {
				needShink = false
			}
			left++
		}
	}
	return maxLen
}

func longestValidParentheses(s string) int {
	return 0
}

func minmunTotalImpl(triangle [][]int, i, j int, table [][]int) int {
	if j > i || i < 0 || j < 0 {
		return 10001
	}

	if i == 0 && j == 0 {
		return triangle[0][0]
	}
	if table[i][j] != 10001 {
		return table[i][j]
	}

	var ret int
	left := minmunTotalImpl(triangle, i-1, j-1, table)
	up := minmunTotalImpl(triangle, i-1, j, table)
	if left < up {
		ret = left + triangle[i][j]
	} else {
		ret = up + triangle[i][j]
	}

	table[i][j] = ret
	return ret
}

func minimumTotal(triangle [][]int) int {
	tLen := len(triangle)
	var table [][]int = make([][]int, 0, tLen)
	for i := 0; i < tLen; i++ {
		var tmp []int = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			tmp[j] = 10001
		}
		table = append(table, tmp)
	}

	var min int = 10001
	for j := 0; j < tLen; j++ {
		get := minmunTotalImpl(triangle, tLen-1, j, table)
		if get < min {
			min = get
		}
	}

	return min
}

func maxProfit(prices []int) int {
	var hasStock bool
	//var maxProfit int
	var getStockPrice int
	var pricesList [][]int

	for i := 0; i < len(prices); i++ {
		if !hasStock {
			if i+1 >= len(prices) {
				break
			} else {
				if prices[i+1] <= prices[i] {
					continue
				} else {
					hasStock = true
					getStockPrice = prices[i]
				}
			}
		} else {
			if i+1 >= len(prices) {
				tmp := make([]int, 2)
				tmp[0] = getStockPrice
				tmp[1] = prices[i]
				pricesList = append(pricesList, tmp)
				hasStock = false
				getStockPrice = 0
			} else {
				if prices[i+1] > prices[i] {
					continue
				} else {
					tmp := make([]int, 2)
					tmp[0] = getStockPrice
					tmp[1] = prices[i]
					pricesList = append(pricesList, tmp)
					hasStock = false
					getStockPrice = 0
				}
			}
		}
	}

	return 0
}
