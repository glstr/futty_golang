package leetcode

/*
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

说明:
s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
*/
func isMatch(s string, p string) bool {
	match := func(i, j int) bool {
		if i == 0 {
			return false
		}

		if p[j-1] == '.' {
			return true
		}

		return s[i-1] == p[j-1]
	}

	sl := len(s)
	pl := len(p)

	states := make([][]int, sl+1, sl+1)
	for i := 0; i <= sl; i++ {
		states[i] = make([]int, pl+1, pl+1)
	}

	states[0][0] = 1
	for i := 0; i <= sl; i++ {
		for j := 1; j <= pl; j++ {
			if p[j-1] == '*' {
				states[i][j] |= states[i][j-2]
				if match(i, j-1) {
					states[i][j] |= states[i-1][j]
				}
			} else {
				if match(i, j) {
					states[i][j] |= states[i-1][j-1]
				}
			}
		}
	}
	return states[sl][pl] == 1

}

/*
给定一个非空二叉树，返回其最大路径和。

本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。

示例 1:

输入: [1,2,3]

       1
      / \
     2   3

输出: 6
示例 2:

输入: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

输出: 42
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var sum int
var first bool

func nodeMaxNum(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftMax := max(nodeMaxNum(node.Left), 0)
	rightMax := max(nodeMaxNum(node.Right), 0)

	newSum := node.Val + leftMax + rightMax
	if first {
		sum = newSum
	} else {
		sum = max(newSum, sum)
	}
	return node.Val + max(leftMax, rightMax)
}

func maxPathNum(root *TreeNode) int {
	sum = 0
	first = true
	nodeMaxNum(root)
	return sum
}
