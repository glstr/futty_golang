package link

import (
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var sum int
	var cur int = 0
	for _, cg := range g {
		for cur < len(s) && cg > s[cur] {
			cur++
		}
		if cur == len(s) {
			return sum
		}
		cur++
		sum++
	}
	return sum
}

func candy(ratings []int) int {
	if len(ratings) < 2 {
		return len(ratings)
	}

	var results = make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		results[i] = 1
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			results[i] = results[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && results[i] <= results[i+1] {
			results[i] = results[i+1] + 1
		}
	}

	var sum int
	for _, r := range results {
		sum += r
	}
	return sum
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 0 {
		return false
	}

	if n == 0 {
		return true
	}

	if len(flowerbed) == 1 {
		if flowerbed[0] == 0 && n <= 1 {
			return true
		} else {
			return false
		}
	}

	var canSet int
	for i := 0; i < len(flowerbed); i++ {
		if i == 0 {
			if flowerbed[i] != 1 && flowerbed[i+1] != 1 {
				canSet++
				flowerbed[i] = 1
			}
		} else if i > 0 && i < len(flowerbed)-1 {
			if flowerbed[i] != 1 && flowerbed[i-1] != 1 && flowerbed[i+1] != 1 {
				canSet++
				flowerbed[i] = 1
			}
		} else {
			if flowerbed[i] != 1 && flowerbed[i-1] != 1 {
				canSet++
				flowerbed[i] = 1
			}
		}
	}

	return canSet >= n
}

func findMinArrowShots(points [][]int) int {
	if len(points) == 1 {
		return 1
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	getInsert := func(a []int, b []int) []int {
		if a[1] < b[0] {
			return nil
		} else {
			if a[1] < b[1] {
				return []int{b[0], a[1]}
			} else {
				return []int{b[0], b[1]}
			}
		}
	}

	cur := points[0]
	var result int
	for i := 1; i < len(points); i++ {
		get := getInsert(cur, points[i])
		if get == nil {
			result++
			cur = points[i]
		} else {
			cur = get
		}
	}
	result++
	return result
}
