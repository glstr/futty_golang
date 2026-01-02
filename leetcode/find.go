package leetcode

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 请必须使用时间复杂度为 O(log n) 的算法。
// 提示:
// 1 <= nums.length <= 104
// -104 <= nums[i] <= 104
// nums 为 无重复元素 的 升序 排列数组
// -104 <= target <= 104
func SearchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	begin := 0
	end := len(nums) - 1
	for (end-begin)/2 > 0 {
		targetIndex := (end-begin)/2 + begin
		tarNum := nums[targetIndex]
		if target > tarNum {
			begin = targetIndex
		} else {
			end = targetIndex
		}
	}

	if begin == end {
		if target <= nums[begin] {
			return begin
		} else {
			return begin + 1
		}
	} else {
		if target <= nums[begin] {
			return begin
		} else if target > nums[begin] && target <= nums[end] {
			return begin + 1
		} else {
			return end + 1
		}
	}
}

// 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
// 如果数组中不存在目标值 target，返回 [-1, -1]。
// 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
// 0 <= nums.length <= 105
// -109 <= nums[i] <= 109
// nums 是一个非递减数组
// -109 <= target <= 109
func SearchRange(nums []int, target int) []int {
	defaultResult := []int{-1, -1}
	numsLen := len(nums)
	if numsLen <= 0 {
		return defaultResult
	}

	if target < nums[0] || target > nums[numsLen-1] {
		return defaultResult
	}

	begin := 0
	end := numsLen - 1

	for (end-begin)/2 > 0 {
		targetIndex := (end-begin)/2 + begin
		if nums[targetIndex] == target {
			begin := targetIndex
			for begin >= 0 {
				if nums[begin] != target {
					break
				}
				begin--
			}
			if begin < 0 || nums[begin] != target {
				begin++
			}

			end := targetIndex
			for end < numsLen {
				if nums[end] != target {
					break
				}
				end++
			}

			if end >= numsLen || nums[end] != target {
				end--
			}
			return []int{begin, end}

		} else if target < nums[targetIndex] {
			end = targetIndex
		} else {
			begin = targetIndex
		}
	}

	if begin == end {
		if target == nums[begin] {
			return []int{begin, end}
		} else {
			return defaultResult
		}
	} else {
		if target == nums[begin] && target == nums[end] {
			return []int{begin, end}
		} else if target == nums[begin] {
			return []int{begin, begin}
		} else if target == nums[end] {
			return []int{end, end}
		} else {
			return defaultResult
		}
	}
}

// 搜索旋转数组。给定一个排序后的数组，包含n个整数，但这个数组已被旋转过很多次了，次数不详。
// 请编写代码找出数组中的某个元素，假设数组元素原先是按升序排列的。若有多个相同元素，返回索引值最小的一个。
// 输入: arr = [15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14], target = 5
// arr = [8, 1, 2, 3, 4]
// 输出: 8（元素5在该数组中的索引）
// arr 长度范围在[1, 1000000]之间
func Search(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	if arr[0] == target {
		return 0
	}

	begin := 0
	end := len(arr) - 1

	for (end-begin)/2 > 0 {
		mid := (end + begin) / 2
		if arr[mid] == target {
			if arr[0] == target {
				return 0
			}

			for mid > 0 {
				mid--
				if arr[mid] != target {
					return mid + 1
				}
			}
			return mid + 1
		}

		for arr[mid] == arr[begin] && mid > begin {
			begin++
		}

		for arr[mid] == arr[end] && mid < end {
			end--
		}

		if arr[mid] <= arr[end] {
			if target >= arr[mid] && target <= arr[end] {
				begin = mid + 1
			} else {
				end = mid - 1
			}
		} else {
			if target <= arr[mid] && target >= arr[begin] {
				end = mid - 1
			} else {
				begin = mid + 1
			}
		}
	}

	if arr[begin] == target {
		return begin
	} else if arr[end] == target {
		return end
	} else {
		return -1
	}
}

// 已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,2,4,5,6,7] 在变化后可能得到：
// 若旋转 4 次，则可以得到 [4,5,6,7,0,1,2]
// 若旋转 7 次，则可以得到 [0,1,2,4,5,6,7]
// 注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。
//
// 给你一个元素值 互不相同 的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
// 你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
func FindMin(nums []int) int {
	if len(nums) <= 0 {
		return -1
	}

	begin := 0
	end := len(nums) - 1

	if nums[0] <= nums[end] {
		return nums[0]
	}

	for (end-begin)/2 > 0 {
		mid := (end + begin) / 2
		if nums[mid] > nums[0] {
			begin = mid
		} else {
			end = mid
		}
	}

	if begin == end {
		return nums[begin]
	} else {
		return nums[end]
	}
}

// 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
type ListNode struct {
	Val  int
	Next *ListNode
}

func IsPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	if head.Next == nil {
		return true
	}

	toEnd := head
	toMid := head
	isOdd := false
	for {
		toEnd = toEnd.Next
		if toEnd == nil {
			isOdd = true
			break
		}
		toEnd = toEnd.Next
		if toEnd == nil {
			isOdd = false
			break
		}
		toMid = toMid.Next
	}

	if toMid == nil {
		return false
	}

	nextSectionStart := toMid.Next
	if nextSectionStart == nil {
		return false
	}

	ite := head
	var front *ListNode = nil
	var next *ListNode = nil
	for ite != toMid {
		next = ite.Next
		ite.Next = front
		front = ite
		ite = next
	}
	toMid.Next = front

	if isOdd {
		toMid = toMid.Next
	}

	for {
		if toMid.Val != nextSectionStart.Val {
			return false
		}

		toMid = toMid.Next
		nextSectionStart = nextSectionStart.Next
		if toMid == nil {
			return true
		}
	}
}

// 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
func Rotate(nums []int, k int) {
	newK := k % len(nums)
	numsLen := len(nums)

	reverse := func(nums []int) {
		for i := 0; i < len(nums)/2; i++ {
			tmp := nums[i]
			nums[i] = nums[len(nums)-1-i]
			nums[len(nums)-1-i] = tmp
		}
	}

	reverse(nums[:numsLen-newK])
	reverse(nums[numsLen-newK:])
	reverse(nums)
}

// https://leetcode.cn/leetbook/read/top-interview-questions-easy/xn6gq1/
// 打乱数组

type Solution struct {
	origin []int
}

func Constructor(nums []int) Solution {
	return Solution{
		origin: nums,
	}
}

func (s *Solution) Reset() []int {
	return s.origin
}

func (s *Solution) Shuffle() []int {
	return nil
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

// https://leetcode.cn/leetbook/read/top-interview-questions-easy/xncfnv/
// 杨辉三角

func generate(numRows int) [][]int {
	var result [][]int
	for i := 1; i <= numRows; i++ {
		if i == 1 {
			result = append(result, []int{1})
		} else {
			var tmp []int
			for j := 1; j <= i; j++ {
				if j == 1 {
					tmp = append(tmp, 1)
				} else if j == i {
					tmp = append(tmp, 1)
				} else {
					tmp = append(tmp, result[i-2][j-2]+result[i-2][j-1])
				}
			}
			result = append(result, tmp)
		}
	}
	return result
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	left := inorderTraversal(root.Left)
	result = append(result, left...)

	result = append(result, root.Val)

	right := inorderTraversal(root.Right)
	result = append(result, right...)
	return result
}

func inorderTraversalEx(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var tmp []*TreeNode
	var result []int
	for root != nil || len(tmp) > 0 {
		for root != nil {
			tmp = append(tmp, root)
			root = root.Left
		}

		root = tmp[len(tmp)-1]
		tmp = tmp[:len(tmp)-1]
		result = append(result, root.Val)
		root = root.Right
	}
	return result
}
