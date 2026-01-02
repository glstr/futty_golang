package leetcode

import "testing"

func TestSearchInsert(t *testing.T) {
	type UnitCase struct {
		nums   []int
		target int
		expect int
	}

	cases := []UnitCase{
		{
			nums:   []int{1, 3, 5, 6},
			target: 5,
			expect: 2,
		},

		{
			nums:   []int{},
			target: 1,
			expect: 0,
		},

		{
			nums:   []int{1, 3, 5, 6},
			target: 7,
			expect: 4,
		},
	}

	for _, c := range cases {
		get := SearchInsert(c.nums, c.target)
		if get != c.expect {
			t.Errorf("c:%v, get:%d", c, get)
		}
	}
}

func TestSearchRange(t *testing.T) {
	type UnitCase struct {
		nums   []int
		target int
		expect []int
	}

	cases := []UnitCase{
		{
			nums:   []int{1, 1, 2},
			target: 1,
			expect: []int{0, 1},
		},

		{
			nums:   []int{1, 2, 2},
			target: 2,
			expect: []int{1, 2},
		},

		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			expect: []int{3, 4},
		},

		{
			nums:   []int{1, 3, 5, 6},
			target: 7,
			expect: []int{-1, -1},
		},

		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			expect: []int{3, 4},
		},
	}

	for _, c := range cases {
		get := SearchRange(c.nums, c.target)
		if get[0] != c.expect[0] || get[1] != c.expect[1] {
			t.Errorf("c:%v, get:%v", c, get)
		}
	}
}

func TestSearch(t *testing.T) {
	type UnitCase struct {
		nums   []int
		target int
		expect int
	}

	cases := []UnitCase{
		{
			nums:   []int{1, 1, 1, 1, 1, 2, 1, 1, 1},
			target: 2,
			expect: 5,
		},

		{
			nums:   []int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14},
			target: 5,
			expect: 8,
		},

		{
			nums:   []int{5, 5, 5, 1, 2, 3, 4, 5},
			target: 5,
			expect: 0,
		},
	}

	for _, c := range cases {
		get := Search(c.nums, c.target)
		if c.expect != get {
			t.Errorf("c:%v, get:%d", c, get)
		}
	}
}

func TestFindMin(t *testing.T) {
	type UnitCase struct {
		nums   []int
		expect int
	}

	cases := []UnitCase{
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			expect: 0,
		},
	}

	for _, c := range cases {
		get := FindMin(c.nums)
		if c.expect != get {
			t.Errorf("c:%v, get:%d", c, get)
		}
	}

}

func TestIsPalindrome(t *testing.T) {
	nodeA := ListNode{1, nil}
	nodeB := ListNode{2, nil}
	nodeC := ListNode{1, nil}
	//nodeD := ListNode{1, nil}
	nodeA.Next = &nodeB
	nodeB.Next = &nodeC
	//nodeC.Next = &nodeD

	get := IsPalindrome(&nodeA)
	if get != true {
		t.Errorf("get:%t, expect:%t", get, true)
	}
}

func TestRotate(t *testing.T) {
	type UnitCase struct {
		nums []int
		k    int
	}

	cases := []UnitCase{
		{
			nums: []int{1, 2, 3, 4, 5, 6},
			k:    4,
		},
	}

	for _, c := range cases {
		Rotate(c.nums, c.k)
		t.Logf("nums:%v", c.nums)
	}
}

func TestGenerate(t *testing.T) {
	result := generate(5)
	t.Errorf("%v", result)
}
