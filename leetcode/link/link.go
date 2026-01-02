package link

import "fmt"

// https://labuladong.github.io/algo/di-yi-zhan-da78c/shou-ba-sh-8f30d/shuang-zhi-0f7cc/
/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := new(ListNode)
	p := dummy
	p1 := list1
	p2 := list2
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}

	if p1 != nil {
		p.Next = p1
	}

	if p2 != nil {
		p.Next = p2
	}

	return dummy.Next
}

func Partition(head *ListNode, x int) *ListNode {
	left := new(ListNode)
	right := new(ListNode)
	pLeft := left
	pRight := right
	for ite := head; ite != nil; {
		if ite.Val < x {
			pLeft.Next = ite
			pLeft = pLeft.Next
		} else {
			pRight.Next = ite
			pRight = pRight.Next
		}

		tmp := ite.Next
		ite.Next = nil
		ite = tmp
	}

	pLeft.Next = right.Next

	return left.Next
}

func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	start := lists[0]
	lists = lists[1:]
	for _, list := range lists {
		start = mergeTwoLists(start, list)
	}
	return start
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	p := dummy
	fast := p
	slow := p
	dist := 0
	for fast != nil {
		if dist <= n {
			dist = dist + 1
		} else {
			slow = slow.Next
		}
		fast = fast.Next
	}

	if slow.Next != nil {
		tmp := slow.Next
		slow.Next = tmp.Next
		tmp.Next = nil
	}

	return dummy.Next
}

func MiddleNode(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// 反转链表
func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := new(ListNode)
	dummy.Next = head
	originFirst := head

	first := dummy
	second := head
	for second != nil {
		next := second.Next
		second.Next = first
		first = second
		second = next
	}

	originFirst.Next = nil
	return first
}

func ReverseListRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := ReverseListRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}

	var dummy = new(ListNode)
	dummy.Next = head
	tmp := dummy
	var front *ListNode = nil
	var end *ListNode = nil
	for tmp != nil {
		if left == 1 {
			front = tmp
		}
		if right == 0 {
			end = tmp
			break
		}
		tmp = tmp.Next
		right = right - 1
		left = left - 1
	}

	begin := front.Next
	endNext := end.Next
	end.Next = nil
	newBegin := ReverseListRecursion(begin)
	front.Next = newBegin
	begin.Next = endNext
	return dummy.Next
}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	tmpk := k
	if k == 1 || head == nil {
		return head
	}

	tmp := head
	for tmp != nil && k > 1 {
		tmp = tmp.Next
		k = k - 1
	}

	if tmp == nil {
		return head
	}

	nextBegin := tmp.Next
	tmp.Next = nil
	newBegin := ReverseListRecursion(head)
	head.Next = ReverseKGroup(nextBegin, tmpk)
	return newBegin
}

// 环形链表
func HasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil && slow != nil {
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return false
		}

		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func DetectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && slow != nil {
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return nil
		}
		slow = slow.Next
		if slow == fast {
			break
		}
	}

	newStart := head
	for newStart != nil && fast != nil {
		if newStart == fast {
			return newStart
		}
		newStart = newStart.Next
		fast = fast.Next
	}

	return nil
}

func FrontIter(head *ListNode, result []int) {
	if head == nil {
		return
	}

	result = append(result, head.Val)
	FrontIter(head, result)
	return
}

func EndIter(head *ListNode, result []int) {
	if head == nil {
		return
	}

	FrontIter(head, result)
	result = append(result, head.Val)
	return
}

// 回文链表
var left *ListNode

func reverse(right *ListNode) bool {
	if right == nil {
		return true
	}

	ret := reverse(right.Next)
	ret = ret && (right.Val == left.Val)
	left = left.Next
	return ret
}

func isPalindrome(head *ListNode) bool {
	left = head
	right := head
	return reverse(right)
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{
		Next: head,
		Val:  head.Val - 1,
	}
	fast := head
	slow := dummy
	enterEq := false
	cur := 0
	for fast != nil {
		if !enterEq {
			if fast.Next != nil {
				if fast.Val == fast.Next.Val {
					enterEq = true
					cur = fast.Val
				} else {
					slow = fast
				}
			}
			fast = fast.Next
		} else {
			if fast.Val != cur {
				slow.Next = fast
				enterEq = false
			} else {
				fast = fast.Next
			}
		}
	}

	if enterEq {
		slow.Next = fast
	}

	return dummy.Next
}

func swapNodes(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{
		Next: head,
	}

	fast := dummy
	slow := dummy
	var k1 *ListNode = nil
	num := 1
	for fast.Next != nil {
		fast = fast.Next
		if num == k {
			k1 = fast
		}
		if num > k {
			slow = slow.Next
		}
		num++
	}

	tmp := k1.Val
	k1.Val = slow.Val
	slow.Val = tmp
	return dummy.Next
}

func sortList(head *ListNode) *ListNode {
	return nil
}

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{
		Next: head,
	}
	fast := dummy
	slow := dummy

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return dummy.Next
}

func removeNodes(head *ListNode) *ListNode {
	reverseList := func(head *ListNode) *ListNode {
		if head == nil {
			return head
		}

		var last *ListNode = nil
		cur := head
		for cur != nil {
			tmp := cur.Next
			cur.Next = last
			last = cur
			cur = tmp
		}

		return last
	}

	newHead := reverseList(head)
	cur := newHead
	var last *ListNode = nil
	max := 0
	for cur != nil {
		if max > cur.Val {
			if last == nil {
				newHead = cur.Next
			}

			if last != nil {
				last.Next = cur.Next
			}

			tmp := cur
			cur = cur.Next
			tmp.Next = nil
		} else {
			max = cur.Val
			last = cur
			cur = cur.Next
		}
	}

	result := reverseList(newHead)
	return result
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{
		Next: head,
	}
	last := dummy
	cur := head
	for cur != nil {
		if cur.Val == val {
			last.Next = cur.Next
			cur = cur.Next
		} else {
			last = cur
			cur = cur.Next
		}
	}

	return dummy.Next
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	head := list2
	cur := list2
	for cur.Next != nil {
		cur = cur.Next
	}
	tail := cur

	delta := b - a
	fast := list1
	slow := list1
	num := 0
	for fast != nil {
		fast = fast.Next
		if num > delta {
			slow = slow.Next
		}
		num++
		if num >= b {
			break
		}
	}

	slow.Next = head
	tail.Next = fast.Next
	fast.Next = nil
	return list1
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	curA := headA
	curB := headB
	for curA != curB {
		if curA == nil {
			curA = headB
		}

		if curB == nil {
			curB = headA
		}

		curA = curA.Next
		curB = curB.Next
	}

	return curA
}

func nextLargerNodes(head *ListNode) []int {
	reverse := func(head *ListNode) *ListNode {
		if head == nil {
			return head
		}

		var pre *ListNode = nil
		cur := head
		for cur != nil {
			tmp := cur.Next
			cur.Next = pre
			pre = cur
			cur = tmp
		}
		return pre
	}

	newHead := reverse(head)

	cur := newHead
	var stack []int
	var result []int
	for cur != nil {
		for len(stack) != 0 && cur.Val >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			result = append(result, 0)
		} else {
			result = append(result, stack[len(stack)-1])
		}
		stack = append(stack, cur.Val)
		cur = cur.Next
	}

	fmt.Println(result)
	left := 0
	right := len(result) - 1
	for left < right {
		tmp := result[left]
		result[left] = result[right]
		result[right] = tmp
		left++
		right--
	}
	return result
}

func oddEvenList(head *ListNode) *ListNode {
	dummyFirst := &ListNode{}
	dummySecond := &ListNode{}

	cur := head
	curFirst := dummyFirst
	curSecond := dummySecond
	var first bool = true
	for cur != nil {
		if first {
			curFirst.Next = cur
			curFirst = curFirst.Next
			cur = cur.Next
			curFirst.Next = nil
			first = false
		} else {
			curSecond.Next = cur
			curSecond = curSecond.Next
			cur = cur.Next
			curSecond.Next = nil
			first = true
		}
	}

	curFirst.Next = dummySecond.Next
	return dummyFirst.Next
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	cur := head
	for cur != nil && cur.Next != nil {
		newNode := &ListNode{
			Val:  gcd(cur.Val, cur.Next.Val),
			Next: cur.Next,
		}
		tmp := cur.Next
		cur.Next = newNode
		cur = tmp
	}
	return head
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	if head == nil {
		return nil
	}

	var listLen int
	cur := head
	for cur != nil {
		listLen++
		cur = cur.Next
	}

	var subListLenMax int
	var subListMaxNum int
	var subListLenMin int
	var subListMinNum int
	if listLen <= k {
		subListLenMax = 1
		subListMaxNum = listLen
		subListLenMin = 0
		subListMinNum = k - listLen
	} else {
		subListLenMax = listLen / k
		subListLenMin = subListLenMax
		lastNode := listLen % k
		if lastNode > 0 {
			subListLenMax += 1
			subListMaxNum = lastNode
			subListMinNum = k - lastNode
		} else {
			subListMaxNum = k
			subListMinNum = 0
		}
	}

	getSubNode := func(head *ListNode, k int) (*ListNode, *ListNode) {
		if head == nil {
			return nil, head
		}

		dummy := &ListNode{
			Next: head,
		}
		cur := dummy
		for k > 0 && cur != nil {
			k--
			cur = cur.Next
		}

		if cur != nil {
			tmp := cur.Next
			cur.Next = nil
			return dummy.Next, tmp
		} else {
			return dummy.Next, nil
		}
	}

	cur = head
	var result []*ListNode
	for subListMaxNum > 0 {
		ret, newHead := getSubNode(cur, subListLenMax)
		cur = newHead
		result = append(result, ret)
		subListMaxNum--
	}

	for subListMinNum > 0 {
		ret, newHead := getSubNode(cur, subListLenMin)
		cur = newHead
		result = append(result, ret)
		subListMinNum--
	}

	return result
}
