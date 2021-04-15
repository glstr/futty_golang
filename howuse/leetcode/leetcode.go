package leetcode

import "strings"

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

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/
func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	numMap := make(map[int]int, len(nums))
	for i, val := range nums {
		numMap[val] = i
	}

	for i, val := range nums {
		if j, ok := numMap[target-val]; ok {
			if j != i {
				res[0] = i
				res[1] = j
				return res
			}
		}
	}

	return res
}

/*
给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

说明：

拆分时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。

输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false
*/

func wordBreak(s string, wordDict []string) bool {
	var temp = make(map[string]bool)
	return isWordBreak(s, wordDict, temp)
}

func isWordBreak(s string, wordDict []string, hasResult map[string]bool) bool {
	if s == "" {
		return true
	}

	for _, word := range wordDict {
		if strings.HasPrefix(s, word) {
			subStr := strings.TrimPrefix(s, word)
			if v, ok := hasResult[subStr]; ok {
				if v {
					return true
				}
			} else {
				hasResult[subStr] = isWordBreak(subStr, wordDict, hasResult)
				if hasResult[subStr] {
					return true
				}
			}
		}
	}

	return false
}

/*给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{}
	cur := pre
	carry := 0
	for l1 != nil || l2 != nil {
		var l1Var, l2Var int
		if l1 != nil {
			l1Var = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			l2Var = l2.Val
			l2 = l2.Next
		}

		sum := carry + l1Var + l2Var
		newVal := sum % 10
		carry = sum / 10
		nextNode := &ListNode{
			Val: newVal,
		}

		cur.Next = nextNode
		cur = cur.Next
	}

	if carry > 0 {
		nextNode := &ListNode{
			Val: 1,
		}
		cur.Next = nextNode
	}
	return pre.Next
}

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/
func lengthOfLongestSubstring(s string) int {

}
