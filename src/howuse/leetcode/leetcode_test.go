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
