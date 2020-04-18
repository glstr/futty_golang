package model

import "testing"

func TestQueue(t *testing.T) {
	exa := []int{123, 233, 345}
	q := NewQueue()
	for _, v := range exa {
		q.Push(v)
	}

	var res []int
	for q.Len() > 0 {
		t := q.Pop()
		if v, ok := t.(int); ok {
			res = append(res, v)
		}
	}

	if res[0] != exa[0] ||
		res[1] != exa[1] ||
		res[2] != exa[2] {
		t.Errorf("fail, expect:%v, real:%v", exa, res)
	}
	t.Logf("res:%v", res)
}

func TestStack(t *testing.T) {
	exa := []int{123, 132, 321}
	s := NewStack()
	for _, v := range exa {
		s.Push(v)
	}

	var res []int
	for s.Len() > 0 {
		t := s.Pop()
		if v, ok := t.(int); ok {
			res = append(res, v)
		}
	}

	if res[0] != exa[2] ||
		res[1] != exa[1] ||
		res[2] != exa[0] {
		t.Errorf("fail, expect:%v, real:%v", exa, res)
	}
	t.Logf("res:%v", res)
}
