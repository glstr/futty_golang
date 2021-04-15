package leetcode

import "testing"

func TestIsMatch(t *testing.T) {
	type unitcase struct {
		s      string
		p      string
		expect bool
	}

	cases := []unitcase{
		unitcase{
			s:      "aa",
			p:      "a",
			expect: false,
		},
		unitcase{
			s:      "aa",
			p:      "a*",
			expect: true,
		},
		unitcase{
			s:      "ab",
			p:      ".*",
			expect: true,
		},
		unitcase{
			s:      "aab",
			p:      "c*a*b",
			expect: true,
		},
		unitcase{
			s:      "mississippi",
			p:      "mis*is*p*.",
			expect: false,
		},
	}

	for _, c := range cases {
		ret := isMatch(c.s, c.p)
		if ret != c.expect {
			t.Errorf("case:%v, expect:%v, real:%v", c, c.expect, ret)
		}
	}
}

func TestMaxPathNum(t *testing.T) {
}

func TestTwoNum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	t.Logf("res:%v", res)
	if res[0] != 0 && res[1] != 1 {
		t.Errorf("res:%v, expect:0, 1", res)
	}
}

func TestWordBreak(t *testing.T) {
	type unicase struct {
		s      string
		wd     []string
		expect bool
	}

	cases := []unicase{
		unicase{
			s:      "leetcode",
			wd:     []string{"leet", "code"},
			expect: true,
		},
		unicase{
			s:      "applepenapple",
			wd:     []string{"apple", "pen"},
			expect: true,
		},
		unicase{
			s:      "catsandog",
			wd:     []string{"cats", "dog", "sand", "and", "cat"},
			expect: false,
		},
		unicase{
			s:      "applepie",
			wd:     []string{"pie", "pear", "apple", "peach"},
			expect: true,
		},
	}

	for _, c := range cases {
		ret := wordBreak(c.s, c.wd)
		if ret != c.expect {
			t.Errorf("c:%v, expect:%t, real:%t", c, c.expect, ret)
		}
	}
}
