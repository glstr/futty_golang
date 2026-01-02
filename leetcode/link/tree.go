package link

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var cur []*TreeNode
	var next []*TreeNode
	if root == nil {
		return nil
	}
	cur = append(cur, root)
	var result [][]int
	var fromLeft bool = true
	for len(cur) != 0 {
		if fromLeft {
			var tmp []int
			for i := len(cur) - 1; i >= 0; i-- {
				if cur[i].Left != nil {
					next = append(next, cur[i].Left)
				}
				if cur[i].Right != nil {
					next = append(next, cur[i].Right)
				}
				tmp = append(tmp, cur[i].Val)
			}
			result = append(result, tmp)
			cur = next
			next = nil
			fromLeft = false
		} else {
			var tmp []int
			for i := len(cur) - 1; i >= 0; i-- {
				if cur[i].Right != nil {
					next = append(next, cur[i].Right)
				}
				if cur[i].Left != nil {
					next = append(next, cur[i].Left)
				}
				tmp = append(tmp, cur[i].Val)
			}
			result = append(result, tmp)
			cur = next
			next = nil
			fromLeft = true
		}
	}
	return result
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil && q != nil {
		return false
	}

	if p != nil && q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var cur []*TreeNode
	var next []*TreeNode

	cur = append(cur, root)
	var result [][]int
	for len(cur) != 0 {
		var tmp []int
		for i := 0; i < len(cur); i++ {
			if cur[i].Left != nil {
				next = append(next, cur[i].Left)
			}
			if cur[i].Right != nil {
				next = append(next, cur[i].Right)
			}
			tmp = append(tmp, cur[i].Val)
		}
		cur = next
		next = nil
		result = append(result, tmp)
	}

	left := 0
	right := len(result) - 1
	for left < right {
		tmp := result[right]
		result[right] = result[left]
		result[left] = tmp
		left++
		right--
	}
	return result
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	type cacheItem struct {
		get    bool
		result int
	}

	var dfs func(root *TreeNode, get bool) int
	var cache map[*TreeNode]cacheItem = make(map[*TreeNode]cacheItem)
	dfs = func(root *TreeNode, get bool) int {
		if root == nil {
			return 0
		}

		if item, ok := cache[root]; ok {
			if item.get == get {
				return item.result
			}
		}

		var result int
		if get {
			result = root.Val + dfs(root.Left, false) + dfs(root.Right, false)
		} else {
			result = max(dfs(root.Left, true), dfs(root.Left, false)) +
				max(dfs(root.Right, true), dfs(root.Right, false))
		}
		cache[root] = cacheItem{
			get:    get,
			result: result,
		}
		return result
	}

	return max(dfs(root, false), dfs(root, true))
}

func widthOfBinaryTree(root *TreeNode) int {
	type TreeNodeItem struct {
		node  *TreeNode
		index int
	}

	var cur []*TreeNodeItem
	var next []*TreeNodeItem
	var widthMax int
	rootItem := &TreeNodeItem{
		node:  root,
		index: 1,
	}
	cur = append(cur, rootItem)
	for len(cur) != 0 {
		width := cur[len(cur)-1].index - cur[0].index + 1
		if width > widthMax {
			widthMax = width
		}
		for _, item := range cur {
			if item.node.Left != nil {
				treeNodeItem := &TreeNodeItem{
					node:  item.node.Left,
					index: item.index*2 - 1,
				}
				next = append(next, treeNodeItem)
			}

			if item.node.Right != nil {
				treeNodeItem := &TreeNodeItem{
					node:  item.node.Right,
					index: item.index * 2,
				}
				next = append(next, treeNodeItem)
			}
		}
		cur = next
		next = nil
	}
	return widthMax
}

func rightSideView(root *TreeNode) []int {
	var result []int
	var depth int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		depth++
		if len(result) < depth {
			result = append(result, root.Val)
		}
		dfs(root.Right)
		dfs(root.Left)
		depth--
	}
	dfs(root)
	return result
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	var dfs func(root *TreeNode)
	var result int
	var sum int
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		sum += root.Val
		if sum == targetSum {
			result++
		}
		dfs(root.Left)
		dfs(root.Right)
		sum -= root.Val
	}
	dfs(root)
	return result + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var path []int
	var depth int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		depth++
		if len(path) < depth {
			path = append(path, root.Val)
		}
		dfs(root.Left)
		dfs(root.Right)
		depth--
	}
	dfs(root)
	return path[len(path)-1]
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	var maxNum int
	var cur []*TreeNode
	var next []*TreeNode
	cur = append(cur, root)
	for len(cur) != 0 {
		for i := 0; i < len(cur); i++ {
			if cur[i].Left != nil {
				next = append(next, cur[i].Left)
			}
			if cur[i].Right != nil {
				next = append(next, cur[i].Right)
			}
			if i == 0 {
				maxNum = cur[i].Val
			} else {
				if cur[i].Val > maxNum {
					maxNum = cur[i].Val
				}
			}
		}
		result = append(result, maxNum)
		cur = next
		next = nil
	}
	return result
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root, p, q)
	right := lowestCommonAncestor(root, p, q)

	if left != nil && right == nil {
		return left
	} else if left == nil && right != nil {
		return right
	} else if left != nil && right != nil {
		return root
	} else {
		return nil
	}
}

func getNumber(root *TreeNode, ops [][]int) int {
	if root == nil {
		return 0
	}
	var newOps [][]int = make([][]int, 0, len(ops))
	var cur []int = ops[0]
	for i := 1; i < len(ops); i++ {
		target := ops[i]
		if cur[0] == target[0] && cur[1] <= target[1] && cur[2] >= target[2] {
		} else if cur[0] == target[0] && target[1] <= cur[1] && target[2] >= cur[2] {
			cur = target
		} else {
			newOps = append(newOps, cur)
			cur = target
		}
	}
	newOps = append(newOps, cur)

	var result int
	var hasSet map[int]bool = make(map[int]bool)
	const setBlue int = 0
	const setRed int = 1
	for _, op := range newOps {
		opType := op[0]
		x := op[1]
		y := op[2]

		var dfs func(root *TreeNode, opType, x, y int)
		dfs = func(root *TreeNode, opType, x, y int) {
			if root == nil {
				return
			}

			if opType == setRed {
				if root.Val >= x && root.Val <= y {
					if isSet, ok := hasSet[root.Val]; ok {
						if !isSet {
							hasSet[root.Val] = true
							result++
						}
					} else {
						hasSet[root.Val] = true
						result++
					}
					dfs(root.Left, opType, x, root.Val)
					dfs(root.Right, opType, root.Val, y)
				} else if root.Val < x {
					dfs(root.Right, opType, x, y)
				} else {
					dfs(root.Left, opType, x, y)
				}
			} else if opType == setBlue {
				if root.Val >= x && root.Val <= y {
					if isSet, ok := hasSet[root.Val]; ok {
						if isSet {
							hasSet[root.Val] = false
							result--
						}
					}
					dfs(root.Left, opType, x, root.Val)
					dfs(root.Right, opType, root.Val, y)
				} else if root.Val < x {
					dfs(root.Right, opType, x, y)
				} else {
					dfs(root.Left, opType, x, y)
				}
			}
		}

		if result == 0 && opType == setBlue {
			continue
		} else {
			dfs(root, opType, x, y)
		}
	}
	return result
}

func findFrequentTreeSum(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	var dataMap map[int]int = make(map[int]int)

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		ret := root.Val + dfs(root.Left) + dfs(root.Right)
		dataMap[ret]++
		return ret
	}

	dfs(root)
	var maxNum int
	for _, v := range dataMap {
		if v > maxNum {
			maxNum = v
		}
	}

	for k, v := range dataMap {
		if v == maxNum {
			result = append(result, k)
		}
	}
	return result
}
