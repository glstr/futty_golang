package link

import "testing"

func arryToList(input []int) *ListNode {
	dummy := &ListNode{
		Next: nil,
	}

	cur := dummy
	for _, v := range input {
		node := &ListNode{
			Val: v,
		}
		cur.Next = node
		cur = cur.Next
	}

	return dummy.Next
}

func TestNextLargerNodes(t *testing.T) {

	node3 := ListNode{
		Val: 5,
	}

	node2 := ListNode{
		Val:  1,
		Next: &node3,
	}

	node1 := ListNode{
		Val:  2,
		Next: &node2,
	}

	get := nextLargerNodes(&node1)
	t.Errorf("get:%v", get)
}

func TestSplitListToParts(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	k := 3
	head := arryToList(array)
	get := splitListToParts(head, k)
	for _, v := range get {
		if v != nil {
			t.Errorf("get:%v", v.Val)
		}
	}
}
